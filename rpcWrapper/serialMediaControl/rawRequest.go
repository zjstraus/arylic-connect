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
