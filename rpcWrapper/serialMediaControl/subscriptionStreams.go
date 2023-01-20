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

package serialMediaControl

import (
	"arylic-connect/transport"
	"context"
	"encoding/hex"
	"encoding/json"
	"log"
	"regexp"
	"strconv"
)

type MetadataChangeMessage struct {
	Title     string `json:"title"`
	Artist    string `json:"artist"`
	Album     string `json:"album"`
	Vendor    string `json:"vendor"`
	Skiplimit int    `json:"skiplimit"`
}

func (rpc *RPC) MetadataChangeChannel(ctx context.Context) <-chan MetadataChangeMessage {
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		replyPrefix = "AXX+MEA+DAT"
	}

	outputChan := make(chan MetadataChangeMessage)
	inputChan := make(chan []byte)
	rpc.transport.RegisterPersistentReader(replyPrefix, inputChan)

	go func() {
		defer func() {
			rpc.transport.UnregisterPersistentReader(replyPrefix, inputChan)
			close(outputChan)
			close(inputChan)
		}()

		parser := regexp.MustCompile(`DAT([{\,:}\w\s"]*)&`)

		for {
			select {
			case <-ctx.Done():
				return
			case input := <-inputChan:
				matches := parser.FindSubmatch(input)
				if matches == nil {
					log.Printf("Got a malformed response in streaming metadata")
					continue
				}

				outputMessage := MetadataChangeMessage{}
				parseErr := json.Unmarshal(matches[1], &outputMessage)
				if parseErr == nil {
					album, _ := hex.DecodeString(outputMessage.Album)
					artist, _ := hex.DecodeString(outputMessage.Artist)
					title, _ := hex.DecodeString(outputMessage.Title)
					vendor, _ := hex.DecodeString(outputMessage.Vendor)
					outputMessage.Album = string(album)
					outputMessage.Artist = string(artist)
					outputMessage.Title = string(title)
					outputMessage.Vendor = string(vendor)
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

func (rpc *RPC) MediaReadyChannel(ctx context.Context) <-chan bool {
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		replyPrefix = "AXX+MEA+RDY"
	}

	outputChan := make(chan bool)
	inputChan := make(chan []byte)
	rpc.transport.RegisterPersistentReader(replyPrefix, inputChan)

	go func() {
		defer func() {
			rpc.transport.UnregisterPersistentReader(replyPrefix, inputChan)
			close(outputChan)
			close(inputChan)
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case <-inputChan:
				select {
				case outputChan <- true:
				// Cool, send worked
				default:
					// just pass on send fails
				}
			}
		}
	}()

	return outputChan
}

func (rpc *RPC) VolumeChannel(ctx context.Context) <-chan float32 {
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		replyPrefix = "AXX+VOL+"
	}

	outputChan := make(chan float32)
	inputChan := make(chan []byte)
	rpc.transport.RegisterPersistentReader(replyPrefix, inputChan)

	go func() {
		defer func() {
			rpc.transport.UnregisterPersistentReader(replyPrefix, inputChan)
			close(outputChan)
			close(inputChan)
		}()

		parser := regexp.MustCompile(`VOL\+(\d+)`)

		for {
			select {
			case <-ctx.Done():
				return
			case input := <-inputChan:
				matches := parser.FindSubmatch(input)
				if matches == nil {
					log.Printf("could not determine volume from string: %s\n" + string(input))
					continue
				}
				parsed, parseErr := strconv.Atoi(string(matches[1]))
				if parseErr != nil {
					continue
				}

				floatCast := float32(parsed)
				if parseErr == nil {
					select {
					case outputChan <- floatCast / 100:
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

func (rpc *RPC) MuteChannel(ctx context.Context) <-chan bool {
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		replyPrefix = "AXX+MUT+"
	}

	outputChan := make(chan bool)
	inputChan := make(chan []byte)
	rpc.transport.RegisterPersistentReader(replyPrefix, inputChan)

	go func() {
		defer func() {
			rpc.transport.UnregisterPersistentReader(replyPrefix, inputChan)
			close(outputChan)
			close(inputChan)
		}()

		parser := regexp.MustCompile(`MUT\+(\d+)`)

		for {
			select {
			case <-ctx.Done():
				return
			case input := <-inputChan:
				matches := parser.FindSubmatch(input)
				if matches == nil {
					log.Printf("could not determine mute from string: %s\n" + string(input))
					continue
				}
				parsed, parseErr := strconv.Atoi(string(matches[1]))
				if parseErr != nil {
					continue
				}

				if parseErr == nil {
					select {
					case outputChan <- parsed == 1:
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

func (rpc *RPC) PlayChannel(ctx context.Context) <-chan bool {
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		replyPrefix = "AXX+PLY+"
	}

	outputChan := make(chan bool)
	inputChan := make(chan []byte)
	rpc.transport.RegisterPersistentReader(replyPrefix, inputChan)

	go func() {
		defer func() {
			rpc.transport.UnregisterPersistentReader(replyPrefix, inputChan)
			close(outputChan)
			close(inputChan)
		}()

		parser := regexp.MustCompile(`PLY\+(\d+)`)

		for {
			select {
			case <-ctx.Done():
				return
			case input := <-inputChan:
				matches := parser.FindSubmatch(input)
				if matches == nil {
					log.Printf("could not determine play from string: %s\n" + string(input))
					continue
				}
				parsed, parseErr := strconv.Atoi(string(matches[1]))
				if parseErr != nil {
					continue
				}

				if parseErr == nil {
					select {
					case outputChan <- parsed == 1:
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
