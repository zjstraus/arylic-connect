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
