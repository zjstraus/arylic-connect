/*
arylic-connect, an API broker for Arylic Audio devices
Copyright (C) 2023  Zach Strauss

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

package httpmedia

import (
	"arylic-connect/rpcWrapper/httpControl"
	"context"
	"errors"
)

func (wrapper *HttpMediaWrapper) GetApList(ctx context.Context, target string) ([]httpControl.DetectedWLAN, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.HttpMediaCons[target]
	if !hasConnection {
		return nil, errors.New("endpoint not found")
	}

	return connection.GetApList(ctx)
}

func (wrapper *HttpMediaWrapper) GetWlanState(ctx context.Context, target string) (httpControl.WlanState, error) {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.HttpMediaCons[target]
	if !hasConnection {
		return httpControl.WLAN_FAIL, errors.New("endpoint not found")
	}

	return connection.GetWlanState(ctx)
}

func (wrapper *HttpMediaWrapper) ConnectToWifi(ctx context.Context, target string, ssid string, channel int, auth string, encryption string, password string) error {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.HttpMediaCons[target]
	if !hasConnection {
		return errors.New("endpoint not found")
	}

	return connection.ConnectToWifi(ctx, ssid, channel, auth, encryption, password)
}

func (wrapper *HttpMediaWrapper) ConnectToHiddenWifi(ctx context.Context, target string, ssid string, password string) error {
	wrapper.OpLock.RLock()
	defer wrapper.OpLock.RUnlock()

	connection, hasConnection := wrapper.HttpMediaCons[target]
	if !hasConnection {
		return errors.New("endpoint not found")
	}

	return connection.ConnectToHiddenWifi(ctx, ssid, password)
}
