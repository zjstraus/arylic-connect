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
	"context"
	"log"
)

// asyncWriteLoop picks an item off the outgoingQueue and puts it on the wire.
func (t *Transport) asyncWriteLoop() {
	for {
		select {
		case <-t.listenerCloser:
			return
		case atomicMsg := <-t.atomicOutgoingQueue:
			t.RegisterOneshotReader(atomicMsg.listenCommand, atomicMsg.outchan)
			err := t.writeMessage(atomicMsg.outMessage)
			if err != nil {
				log.Println(err.Error())
			}
			//time.Sleep(200 * time.Millisecond)
		case msg := <-t.outgoingQueue:
			err := t.writeMessage(msg)
			if err != nil {
				log.Println(err.Error())
			}
			//time.Sleep(200 * time.Millisecond)
		}
	}
}

// SendMessage queues an item to be sent out. This is safe to be called from
// multiple threads as the internal queue keeps the connection clean and allows
// for aborting the context if required.
func (t *Transport) SendMessage(ctx context.Context, message interface{}) error {
	select {
	case t.outgoingQueue <- message:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (t *Transport) SendMessageAtomic(ctx context.Context, message interface{}, command string, outchan chan<- []byte) error {
	wrapped := atomicListenerPair{
		ctx:           ctx,
		listenCommand: command,
		outMessage:    message,
		outchan:       outchan,
	}
	select {
	case t.atomicOutgoingQueue <- wrapped:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
