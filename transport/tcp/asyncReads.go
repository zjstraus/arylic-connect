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

package tcp

import (
	"errors"
	"net"
	"strings"
	"time"
)

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
					for prefix, receivers := range t.persistentRequests {
						if strings.HasPrefix(string(message), prefix) {
							var validReceivers []chan<- []byte
							for _, receiver := range receivers {
								select {
								case receiver <- message:
									validReceivers = append(validReceivers, receiver)
								default:
									continue
								}
							}
							t.persistentRequests[prefix] = validReceivers
						}
					}
					for prefix, receivers := range t.oneshotRequests {
						if strings.HasPrefix(string(message), prefix) {
							for _, receiver := range receivers {
								select {
								case receiver <- message:
									continue
								default:
									continue
								}
							}
							delete(t.oneshotRequests, prefix)
						}
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

func (t *Transport) RegisterPersistentReader(prefix string, channel chan<- []byte) {
	t.requestLocker.Lock()
	defer t.requestLocker.Unlock()

	currentListeners := t.persistentRequests[prefix]
	currentListeners = append(currentListeners, channel)
	t.persistentRequests[prefix] = currentListeners
}

func (t *Transport) RegisterOneshotReader(prefix string, channel chan<- []byte) {
	t.requestLocker.Lock()
	defer t.requestLocker.Unlock()

	currentListeners := t.oneshotRequests[prefix]
	currentListeners = append(currentListeners, channel)
	t.oneshotRequests[prefix] = currentListeners
}
