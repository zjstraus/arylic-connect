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

package websocketControl

import (
	"arylic-connect/rpcWrapper"
	"arylic-connect/transport"
	"context"
)

func requestWithResponse(ctx context.Context, t transport.AsyncMessage, request interface{}, replyCommand string) ([]byte, error) {
	if t == nil {
		return nil, rpcWrapper.ErrTransportNotConnected
	}

	if request == nil {
		return nil, rpcWrapper.ErrUnknownTransportFlavor
	}

	returnChan := make(chan []byte)
	defer close(returnChan)
	otherReaders := t.RegisterOneshotReader(replyCommand, returnChan)
	if !otherReaders {
		sendErr := t.SendMessage(ctx, request)
		if sendErr != nil {
			return nil, sendErr
		}
	}

	select {
	case value := <-returnChan:
		return value, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

//func atomicRequestWithResponse(ctx context.Context, t transport.AsyncLine, request string, replyPrefix string) ([]byte, error) {
//	if t == nil {
//		return nil, rpcWrapper.ErrTransportNotConnected
//	}
//
//	if request == "" {
//		return nil, rpcWrapper.ErrUnknownTransportFlavor
//	}
//
//	returnChan := make(chan []byte)
//	defer close(returnChan)
//
//	sendErr := t.SendMessageAtomic(ctx, request, replyPrefix, returnChan)
//	if sendErr != nil {
//		return nil, sendErr
//	}
//
//	select {
//	case value := <-returnChan:
//		return value, nil
//	case <-ctx.Done():
//		return nil, ctx.Err()
//	}
//}
