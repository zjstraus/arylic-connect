package serialMediaControl

import (
	"arylic-connect/rpcWrapper"
	"arylic-connect/transport"
	"context"
)

func requestWithResponse(ctx context.Context, t transport.AsyncLine, request string, replyPrefix string) ([]byte, error) {
	if t == nil {
		return nil, rpcWrapper.ErrTransportNotConnected
	}

	if request == "" {
		return nil, rpcWrapper.ErrUnknownTransportFlavor
	}

	returnChan := make(chan []byte)
	defer close(returnChan)
	t.RegisterOneshotReader(replyPrefix, returnChan)
	sendErr := t.SendMessage(ctx, request)
	if sendErr != nil {
		return nil, sendErr
	}

	select {
	case value := <-returnChan:
		return value, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
