package serialmedia

import (
	"context"
	"errors"
)

func (wrapper *SerialMediaWrapper) GetLED(ctx context.Context, target string) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.GetLED(ctx)
}

func (wrapper *SerialMediaWrapper) SetLED(ctx context.Context, target string, enable bool) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.SetLED(ctx, enable)
}

func (wrapper *SerialMediaWrapper) ToggleLED(ctx context.Context, target string) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.ToggleLED(ctx)
}

func (wrapper *SerialMediaWrapper) GetBeep(ctx context.Context, target string) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.GetBeep(ctx)
}

func (wrapper *SerialMediaWrapper) SetBeep(ctx context.Context, target string, enable bool) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.SetBeep(ctx, enable)
}

func (wrapper *SerialMediaWrapper) GetName(ctx context.Context, target string) (string, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return "", errors.New("endpoint not found")
	}

	return connection.GetName(ctx)
}

func (wrapper *SerialMediaWrapper) SetName(ctx context.Context, target string, name string) (string, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return "", errors.New("endpoint not found")
	}

	return connection.SetName(ctx, name)
}

func (wrapper *SerialMediaWrapper) GetVoicePrompt(ctx context.Context, target string) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.GetVoicePrompt(ctx)
}

func (wrapper *SerialMediaWrapper) SetVoicePrompt(ctx context.Context, target string, enable bool) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.SetVoicePrompt(ctx, enable)
}
