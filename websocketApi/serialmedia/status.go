package serialmedia

import (
	"arylic-connect/rpcWrapper/serialMediaControl"
	"context"
	"errors"
)

func (wrapper *SerialMediaWrapper) GetStatus(ctx context.Context, target string) (serialMediaControl.EndpointStatus, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return serialMediaControl.EndpointStatus{}, errors.New("endpoint not found")
	}

	return connection.GetStatus(ctx)
}
