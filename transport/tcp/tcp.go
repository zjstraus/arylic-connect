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
	"arylic-connect/transport"
	"net"
	"sync"
	"time"
)

// Transport is an AsyncLine implementation using a TCP stream encapsulated in a
// tunneling protocol to be proxied by the device's Linkplay module.
type Transport struct {
	conn net.Conn

	listenerCloser     chan int
	requestLocker      sync.Mutex
	persistentRequests map[string][]chan<- []byte
	oneshotRequests    map[string][]chan<- []byte

	// use an unbuffered channel to queue up commands to enable the context-aware
	// sending, as sending to a channel is cancel-able, unlike the socket write
	outgoingQueue chan string
}

func New() (*Transport, error) {
	return &Transport{
		persistentRequests: make(map[string][]chan<- []byte),
		oneshotRequests:    make(map[string][]chan<- []byte),
	}, nil
}

func (t *Transport) Connect(target string) error {
	t.Close()

	t.listenerCloser = make(chan int)
	t.outgoingQueue = make(chan string)

	conn, err := net.DialTimeout("tcp", target, 10*time.Second)
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
	return transport.Flavor_TCP
}

func (t *Transport) Target() string {
	if t.conn == nil {
		return ""
	}
	return t.conn.RemoteAddr().String()
}
