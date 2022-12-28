package serialmedia

import (
	"context"
	"errors"
)

func (wrapper *SerialMediaWrapper) RequestReboot(ctx context.Context, target string) error {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return errors.New("endpoint not found")
	}

	return connection.RequestReboot(ctx)
}

func (wrapper *SerialMediaWrapper) RequestStandby(ctx context.Context, target string) error {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return errors.New("endpoint not found")
	}

	return connection.RequestStandby(ctx)
}

func (wrapper *SerialMediaWrapper) RequestReset(ctx context.Context, target string) error {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return errors.New("endpoint not found")
	}

	return connection.RequestReset(ctx)
}

func (wrapper *SerialMediaWrapper) RequestRecover(ctx context.Context, target string) error {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return errors.New("endpoint not found")
	}

	return connection.RequestRecover(ctx)
}
