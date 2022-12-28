package serialmedia

import (
	"arylic-connect/rpcWrapper/serialMediaControl"
	"context"
	"errors"
)

func (wrapper *SerialMediaWrapper) GetMultiroomMode(ctx context.Context, target string) (serialMediaControl.MultiroomMode, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return serialMediaControl.Mode_None, errors.New("endpoint not found")
	}

	return connection.GetMultiroomMode(ctx)
}

func (wrapper *SerialMediaWrapper) GetChannelConfig(ctx context.Context, target string) (serialMediaControl.ChannelConfig, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return serialMediaControl.Channel_Stereo, errors.New("endpoint not found")
	}

	return connection.GetChannelConfig(ctx)
}

func (wrapper *SerialMediaWrapper) GetVolumeSync(ctx context.Context, target string) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.GetVolumeSync(ctx)
}

func (wrapper *SerialMediaWrapper) SetVolumeSync(ctx context.Context, target string, mode bool) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.SetVolumeSync(ctx, mode)
}
