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
	"arylic-connect/transport"
	"context"
	"github.com/gorilla/websocket"
	"sync"
)

type atomicListenerPair struct {
	ctx           context.Context
	listenCommand string
	outMessage    interface{}
	outchan       chan<- []byte
}

type Transport struct {
	conn *websocket.Conn

	listenerCloser     chan int
	requestLocker      sync.Mutex
	persistentRequests map[string][]chan<- []byte
	oneshotRequests    map[string][]chan<- []byte

	// use an unbuffered channel to queue up commands to enable the context-aware
	// sending, as sending to a channel is cancel-able, unlike the socket write
	outgoingQueue chan interface{}

	atomicOutgoingQueue chan atomicListenerPair
}

func New() (*Transport, error) {
	return &Transport{
		persistentRequests:  make(map[string][]chan<- []byte),
		oneshotRequests:     make(map[string][]chan<- []byte),
		atomicOutgoingQueue: make(chan atomicListenerPair),
	}, nil
}

func (t *Transport) Connect(target string) error {
	closeErr := t.Close()
	if closeErr != nil {
		return closeErr
	}

	t.listenerCloser = make(chan int)
	t.outgoingQueue = make(chan interface{})

	conn, _, err := websocket.DefaultDialer.Dial(target, nil)
	if err != nil {
		return err
	}
	t.conn = conn
	go t.asyncReadLoop()
	go t.asyncWriteLoop()
	return nil
}

func (t *Transport) Close() error {
	if t.conn != nil {
		return t.conn.Close()
	}
	if t.listenerCloser != nil {
		close(t.listenerCloser)
	}

	return nil
}

func (t *Transport) Flavor() transport.InterfaceFlavor {
	return transport.Flavor_WS
}

func (t *Transport) Target() string {
	if t.conn == nil {
		return ""
	}
	return t.conn.RemoteAddr().String()
}
