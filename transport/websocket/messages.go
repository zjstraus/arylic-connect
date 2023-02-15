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
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

type commandReturn struct {
	Command string `json:"cmd"`
}

func (t *Transport) writeMessage(payload interface{}) error {
	if t.conn == nil {
		return errors.New("t not connected")
	}

	var writeErr error
	switch castPayload := payload.(type) {
	case string:
		writeErr = t.conn.WriteMessage(websocket.TextMessage, []byte(castPayload))
	default:
		writeErr = t.conn.WriteJSON(payload)
	}

	return writeErr
}

func (t *Transport) readMessage(timeout time.Duration) ([]byte, error) {
	if t.conn == nil {
		return nil, errors.New("t not connected")
	}

	//deadlineErr := t.conn.SetReadDeadline(time.Now().Add(timeout))
	//if deadlineErr != nil {
	//	return nil, deadlineErr
	//}

	msgType, msg, msgErr := t.conn.ReadMessage()
	if msgErr != nil {
		return nil, msgErr
	}
	if msgType != websocket.TextMessage {
		return nil, fmt.Errorf("unknown websocket message type (%d)", msgType)
	}

	return msg, nil
}
