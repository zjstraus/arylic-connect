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
