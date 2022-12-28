package serialmedia

import (
	"context"
	"errors"
)

func (wrapper *SerialMediaWrapper) GetName(ctx context.Context, target string) (string, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return "", errors.New("endpoint not found")
	}

	return connection.GetName(ctx)
}
