package serialmedia

import (
	"context"
	"errors"
)

func (wrapper *SerialMediaWrapper) GetBass(ctx context.Context, target string) (float32, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return 0, errors.New("endpoint not found")
	}

	return connection.GetBass(ctx)
}

func (wrapper *SerialMediaWrapper) SetBass(ctx context.Context, target string, level float32) (float32, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return 0, errors.New("endpoint not found")
	}

	return connection.SetBass(ctx, level)
}

func (wrapper *SerialMediaWrapper) GetTreble(ctx context.Context, target string) (float32, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return 0, errors.New("endpoint not found")
	}

	return connection.GetTreble(ctx)
}

func (wrapper *SerialMediaWrapper) SetTreble(ctx context.Context, target string, level float32) (float32, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return 0, errors.New("endpoint not found")
	}

	return connection.SetTreble(ctx, level)
}

func (wrapper *SerialMediaWrapper) GetVirtualBass(ctx context.Context, target string) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.GetVirtualBass(ctx)
}

func (wrapper *SerialMediaWrapper) SetVirtualBass(ctx context.Context, target string, state bool) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.SetVirtualBass(ctx, state)
}

func (wrapper *SerialMediaWrapper) ToggleVirtualBass(ctx context.Context, target string) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.ToggleVirtualBass(ctx)
}
