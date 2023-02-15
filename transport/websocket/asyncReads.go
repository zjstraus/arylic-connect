/*
arylic-connect, an API broker for Arylic Audio devices
Copyright (C) 2022  Zach Strauss

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

package websocket

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"time"
)

// asyncReadLoop spins off a new thread to wait for messages to come in, match
// them against any active readers, and forward them on.
func (t *Transport) asyncReadLoop() {
	scopedCloser := make(chan int)
	go func() {
		for {
			select {
			case <-scopedCloser:
				return
			default:
				message, messageErr := t.readMessage(5 * time.Second)
				if messageErr == nil {
					t.requestLocker.Lock()
					parsed := commandReturn{}
					jsonParseErr := json.Unmarshal(message, &parsed)
					if jsonParseErr == nil {
						for command, receivers := range t.persistentRequests {
							if command == parsed.Command {
								var validReceivers []chan<- []byte
								for _, receiver := range receivers {
									select {
									case receiver <- message:
										validReceivers = append(validReceivers, receiver)
									default:
										continue
									}
								}
								t.persistentRequests[command] = validReceivers
							}
						}

						for command, receivers := range t.oneshotRequests {
							if command == parsed.Command {
								for _, receiver := range receivers {
									select {
									case receiver <- message:
										continue
									default:
										continue
									}
								}
								delete(t.oneshotRequests, command)
							}
						}
					} else {
						log.Printf("Error in json parsing from message '%s': %s\n", message, jsonParseErr.Error())
					}
					t.requestLocker.Unlock()
				} else {
					if errors.Is(messageErr, net.ErrClosed) {
						return
					}
				}
			}
		}
	}()
	for {
		select {
		case <-t.listenerCloser:
			scopedCloser <- 1
			close(scopedCloser)
			return
		}
	}
}

func (t *Transport) RegisterPersistentReader(command string, channel chan<- []byte) {
	t.requestLocker.Lock()
	defer t.requestLocker.Unlock()

	currentListeners := t.persistentRequests[command]
	currentListeners = append(currentListeners, channel)
	t.persistentRequests[command] = currentListeners
}

func (t *Transport) UnregisterPersistentReader(command string, channel chan<- []byte) {
	t.requestLocker.Lock()
	defer t.requestLocker.Unlock()

	var newListeners []chan<- []byte
	for _, existing := range t.persistentRequests[command] {
		if channel != existing {
			newListeners = append(newListeners, existing)
		}
	}
	t.persistentRequests[command] = newListeners
}

func (t *Transport) RegisterOneshotReader(prefix string, channel chan<- []byte) bool {
	t.requestLocker.Lock()
	defer t.requestLocker.Unlock()

	currentListeners := t.oneshotRequests[prefix]
	listenerLength := len(currentListeners)
	currentListeners = append(currentListeners, channel)
	t.oneshotRequests[prefix] = currentListeners
	return listenerLength > 0
}
