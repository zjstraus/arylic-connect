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

package serialMediaControl

import (
	"arylic-connect/rpcWrapper"
	"arylic-connect/transport"
	"context"
	"errors"
	"regexp"
)

func (rpc *RPC) RequestReboot(ctx context.Context) error {
	if rpc.transport == nil {
		return rpcWrapper.ErrTransportNotConnected
	}

	request := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:SYS:REBOOT&"
	}

	if request == "" {
		return rpcWrapper.ErrUnknownTransportFlavor
	}

	return rpc.transport.SendMessage(ctx, request)
}

func (rpc *RPC) RequestStandby(ctx context.Context) error {
	if rpc.transport == nil {
		return rpcWrapper.ErrTransportNotConnected
	}

	request := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:SYS:STANDBY&"
	}

	if request == "" {
		return rpcWrapper.ErrUnknownTransportFlavor
	}

	return rpc.transport.SendMessage(ctx, request)
}

func (rpc *RPC) RequestReset(ctx context.Context) error {
	if rpc.transport == nil {
		return rpcWrapper.ErrTransportNotConnected
	}

	request := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:SYS:RESET&"
	}

	if request == "" {
		return rpcWrapper.ErrUnknownTransportFlavor
	}

	return rpc.transport.SendMessage(ctx, request)
}

func (rpc *RPC) RequestRecover(ctx context.Context) error {
	if rpc.transport == nil {
		return rpcWrapper.ErrTransportNotConnected
	}

	request := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:SYS:RECOVER&"
	}

	if request == "" {
		return rpcWrapper.ErrUnknownTransportFlavor
	}

	return rpc.transport.SendMessage(ctx, request)
}

func (rpc *RPC) GetWifiPlayback(ctx context.Context) (bool, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:PLA&"
		replyPrefix = "MCU+PAS+RAKOIT:PLA:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return false, reqErr
	}

	parser := regexp.MustCompile(`PLA:(\d)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return false, errors.New("could not determine status from string: " + string(data))
	}

	return string(matches[1]) == "1", nil
}
