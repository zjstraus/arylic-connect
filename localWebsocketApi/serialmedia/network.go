package serialmedia

import (
	"context"
	"errors"
)

func (wrapper *SerialMediaWrapper) GetInternet(ctx context.Context, target string) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.GetInternet(ctx)
}

func (wrapper *SerialMediaWrapper) SetInternet(ctx context.Context, target string, enable bool) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.SetInternet(ctx, enable)
}

func (wrapper *SerialMediaWrapper) GetEthernet(ctx context.Context, target string) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.GetEthernet(ctx)
}

func (wrapper *SerialMediaWrapper) SetEthernet(ctx context.Context, target string, enable bool) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.SetEthernet(ctx, enable)
}

func (wrapper *SerialMediaWrapper) GetWifi(ctx context.Context, target string) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.GetWifi(ctx)
}

func (wrapper *SerialMediaWrapper) SetWifi(ctx context.Context, target string, enable bool) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.SetWifi(ctx, enable)
}

func (wrapper *SerialMediaWrapper) RequestWifiReset(ctx context.Context, target string) error {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return errors.New("endpoint not found")
	}

	return connection.RequestWifiReset(ctx)
}

func (wrapper *SerialMediaWrapper) GetBluetooth(ctx context.Context, target string) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.GetBluetooth(ctx)
}

func (wrapper *SerialMediaWrapper) SetBluetooth(ctx context.Context, target string, enable bool) error {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return errors.New("endpoint not found")
	}

	return connection.SetBluetooth(ctx, enable)
}

func (wrapper *SerialMediaWrapper) GetWifiPlayback(ctx context.Context, target string) (bool, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.SerialMediaCons[target]
	if !hasConnection {
		return false, errors.New("endpoint not found")
	}

	return connection.GetWifiPlayback(ctx)
}
