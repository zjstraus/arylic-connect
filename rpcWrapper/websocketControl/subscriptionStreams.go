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
)

type StatusChangeMessage struct {
	Input  string          `json:"input"`
	Volume int             `json:"vol"`
	Track  json.RawMessage `json:"track"`
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
				outputMessage := StatusChangeMessage{}
				parseErr := json.Unmarshal(input, &outputMessage)
				if parseErr == nil {
					select {
					case outputChan <- outputMessage:
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
