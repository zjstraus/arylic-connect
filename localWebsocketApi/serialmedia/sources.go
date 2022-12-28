package serialmedia

import (
	"arylic-connect/rpcWrapper/serialMediaControl"
	"context"
	"errors"
)

func (wrapper *SerialMediaWrapper) GetSource(ctx context.Context, target string) (serialMediaControl.InputSource, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return serialMediaControl.Input_Unknown, errors.New("endpoint not found")
	}

	return connection.GetSource(ctx)
}

func (wrapper *SerialMediaWrapper) SetSource(ctx context.Context, target string, source serialMediaControl.InputSource) (serialMediaControl.InputSource, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return serialMediaControl.Input_Unknown, errors.New("endpoint not found")
	}

	return connection.SetSource(ctx, source)
}

func (wrapper *SerialMediaWrapper) GetDefaultSource(ctx context.Context, target string) (serialMediaControl.InputSource, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return serialMediaControl.Input_Unknown, errors.New("endpoint not found")
	}

	return connection.GetDefaultSource(ctx)
}

func (wrapper *SerialMediaWrapper) SetDefaultSource(ctx context.Context, target string, source serialMediaControl.InputSource) (serialMediaControl.InputSource, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return serialMediaControl.Input_Unknown, errors.New("endpoint not found")
	}

	return connection.SetDefaultSource(ctx, source)
}

func (wrapper *SerialMediaWrapper) GetInputAutoswitch(ctx context.Context, target string) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.GetInputAutoswitch(ctx)
}

func (wrapper *SerialMediaWrapper) SetInputAutoswitch(ctx context.Context, target string, enable bool) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.SetInputAutoswitch(ctx, enable)
}
