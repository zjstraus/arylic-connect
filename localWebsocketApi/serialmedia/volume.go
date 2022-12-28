/*
arylic-connect, an API broker for Arylic Audio devices
Copyright (C) 2022  Zach Strauss

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

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
