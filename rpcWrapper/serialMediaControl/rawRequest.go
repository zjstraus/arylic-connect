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

package serialMediaControl

import (
	"arylic-connect/rpcWrapper"
	"context"
)

type DirectCommand struct {
	Request  string
	Response string
}

func (rpc *RPC) DirectCommand(ctx context.Context, request string) (DirectCommand, error) {
	command := DirectCommand{Request: request}

	if rpc.transport == nil {
		return command, rpcWrapper.ErrTransportNotConnected
	}

	returnChan := make(chan []byte)
	defer close(returnChan)
	rpc.transport.RegisterOneshotReader("", returnChan)
	sendErr := rpc.transport.SendMessage(ctx, request)
	if sendErr != nil {
		return command, sendErr
	}

	select {
	case value := <-returnChan:
		command.Response = string(value)
		return command, nil
	case <-ctx.Done():
		return command, ctx.Err()
	}
}
