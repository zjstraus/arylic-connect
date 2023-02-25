/*
arylic-connect, an API broker for Arylic Audio devices
Copyright (C) 2023  Zach Strauss

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package websocketControl

import (
	"context"
	"encoding/json"
	"golang.org/x/net/html"
)

type incomingStatusChangeMessage struct {
	Input  string `json:"input"`
	Volume int    `json:"vol"`
	Track  struct {
		Source   string `json:"source"`
		State    string `json:"state"`
		Index    int    `json:"index"`
		Mode     string `json:"mode"`
		Elapsed  int    `json:"elapsed"`
		Duration int    `json:"duration"`
		Meta     struct {
			Title  string `json:"title"`
			Artist string `json:"artist"`
			Album  string `json:"album"`
			Image  string `json:"image"`
		} `json:"meta"`
	} `json:"track"`
}

func (msg incomingStatusChangeMessage) Normalize() StatusChangeMessage {
	return StatusChangeMessage{
		Input:    msg.Input,
		Source:   msg.Track.Source,
		State:    msg.Track.State,
		Index:    msg.Track.Index,
		Mode:     msg.Track.Mode,
		Elapsed:  msg.Track.Elapsed,
		Duration: msg.Track.Duration,
		Title:    html.UnescapeString(msg.Track.Meta.Title),
		Artist:   html.UnescapeString(msg.Track.Meta.Artist),
		Album:    html.UnescapeString(msg.Track.Meta.Album),
		Image:    msg.Track.Meta.Image,
		Volume:   msg.Volume,
	}
}

type StatusChangeMessage struct {
	Input  string
	Source string

	State string
	Index int
	Mode  string

	Elapsed  int
	Duration int

	Title  string
	Artist string
	Album  string
	Image  string

	Volume int
}

func (rpc *RPC) StatusChangeChannel(ctx context.Context) <-chan StatusChangeMessage {
	outputChan := make(chan StatusChangeMessage)
	inputChan := make(chan []byte)
	rpc.transport.RegisterPersistentReader("STATUS", inputChan)

	go func() {
		defer func() {
			rpc.transport.UnregisterPersistentReader("STATUS", inputChan)
			close(outputChan)
			close(inputChan)
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case input := <-inputChan:
				outputMessage := incomingStatusChangeMessage{}
				parseErr := json.Unmarshal(input, &outputMessage)
				if parseErr == nil {
					select {
					case outputChan <- outputMessage.Normalize():
					// Cool, send worked
					default:
						// just pass on send fails
					}
				}
			}
		}
	}()

	return outputChan
}
