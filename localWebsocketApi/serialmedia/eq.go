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
