package serialmedia

import (
	"arylic-connect/rpcWrapper/serialMediaControl"
	"context"
	"errors"
)

func (wrapper *SerialMediaWrapper) DirectCommand(ctx context.Context, target string, command string) (serialMediaControl.DirectCommand, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return serialMediaControl.DirectCommand{}, errors.New("endpoint not found")
	}

	return connection.DirectCommand(ctx, command)
}
