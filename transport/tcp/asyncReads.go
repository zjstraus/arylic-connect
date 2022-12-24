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
