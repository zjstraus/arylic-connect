package serialmedia

import (
	"context"
	"errors"
)

func (wrapper *SerialMediaWrapper) GetVolume(ctx context.Context, target string) (float32, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return 0, errors.New("endpoint not found")
	}

	return connection.GetVolume(ctx)
}

func (wrapper *SerialMediaWrapper) SetVolume(ctx context.Context, target string, level float32) (float32, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return 0, errors.New("endpoint not found")
	}

	return connection.SetVolume(ctx, level)
}

func (wrapper *SerialMediaWrapper) GetMute(ctx context.Context, target string) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.GetMute(ctx)
}

func (wrapper *SerialMediaWrapper) SetMute(ctx context.Context, target string, state bool) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.SetMute(ctx, state)
}

func (wrapper *SerialMediaWrapper) ToggleMute(ctx context.Context, target string) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.ToggleMute(ctx)
}

func (wrapper *SerialMediaWrapper) GetFixedVolume(ctx context.Context, target string) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.GetFixedVolume(ctx)
}

func (wrapper *SerialMediaWrapper) SetFixedVolume(ctx context.Context, target string, state bool) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.SetFixedVolume(ctx, state)
}

func (wrapper *SerialMediaWrapper) GetMaxVolume(ctx context.Context, target string) (float32, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return 0, errors.New("endpoint not found")
	}

	return connection.GetMaxVolume(ctx)
}

func (wrapper *SerialMediaWrapper) SetMaxVolume(ctx context.Context, target string, level float32) (float32, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return 0, errors.New("endpoint not found")
	}

	return connection.SetMaxVolume(ctx, level)
}

func (wrapper *SerialMediaWrapper) GetBalance(ctx context.Context, target string) (float32, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return 0, errors.New("endpoint not found")
	}

	return connection.GetBalance(ctx)
}

func (wrapper *SerialMediaWrapper) SetBalance(ctx context.Context, target string, level float32) (float32, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return 0, errors.New("endpoint not found")
	}

	return connection.SetBalance(ctx, level)
}
