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
	if rpc.endpointVersion.Git != "" {
		return command, nil
	}

	if rpc.transport == nil {
		return command, rpcWrapper.ErrTransportNotConnected
	}

	returnChan := make(chan []byte)
	defer close(returnChan)
	rpc.transport.RegisterOneshotReader("", returnChan)
	rpc.transport.SendMessage(ctx, request)

	select {
	case value := <-returnChan:
		command.Response = string(value)
		return command, nil
	case <-ctx.Done():
		return command, ctx.Err()
	}
}
