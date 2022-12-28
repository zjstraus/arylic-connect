package serialmedia

import (
	"arylic-connect/rpcWrapper/serialMediaControl"
	"context"
	"errors"
)

func (wrapper *SerialMediaWrapper) RequestPlayPause(ctx context.Context, target string) error {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return errors.New("endpoint not found")
	}

	return connection.RequestPlayPause(ctx)
}

func (wrapper *SerialMediaWrapper) RequestNext(ctx context.Context, target string) error {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return errors.New("endpoint not found")
	}

	return connection.RequestNext(ctx)
}

func (wrapper *SerialMediaWrapper) RequestPrevious(ctx context.Context, target string) error {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return errors.New("endpoint not found")
	}

	return connection.RequestPrevious(ctx)
}

func (wrapper *SerialMediaWrapper) RequestStop(ctx context.Context, target string) error {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return errors.New("endpoint not found")
	}

	return connection.RequestStop(ctx)
}

func (wrapper *SerialMediaWrapper) GetLoopMode(ctx context.Context, target string) (serialMediaControl.LoopMode, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return serialMediaControl.Loop_Sequence, errors.New("endpoint not found")
	}

	return connection.GetLoopMode(ctx)
}

func (wrapper *SerialMediaWrapper) SetLoopMode(ctx context.Context, target string, mode serialMediaControl.LoopMode) (serialMediaControl.LoopMode, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return serialMediaControl.Loop_Sequence, errors.New("endpoint not found")
	}

	return connection.SetLoopMode(ctx, mode)
}
