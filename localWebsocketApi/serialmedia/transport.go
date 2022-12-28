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
