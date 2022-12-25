package serialmedia

import (
	"arylic-connect/rpcWrapper/serialMediaControl"
	"context"
	"errors"
)

func (wrapper *SerialMediaWrapper) Version(ctx context.Context, target string) (serialMediaControl.EndpointVersion, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return serialMediaControl.EndpointVersion{}, errors.New("endpoint not found")
	}

	return connection.GetVersion(ctx)
}
