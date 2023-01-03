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

// GetInternet queries if the device has an internet connection.
func (rpc *RPC) GetInternet(ctx context.Context) (bool, error) {
	if rpc.transport == nil {
		return false, rpcWrapper.ErrTransportNotConnected
	}

	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:WWW&"
		replyPrefix = "MCU+PAS+RAKOIT:WWW:"
	}

	if request == "" {
		return false, rpcWrapper.ErrUnknownTransportFlavor
	}

	returnChan := make(chan []byte)
	defer close(returnChan)
	rpc.transport.RegisterOneshotReader(replyPrefix, returnChan)
	sendErr := rpc.transport.SendMessage(ctx, request)
	if sendErr != nil {
		return false, sendErr
	}

	select {
	case value := <-returnChan:
		parser := regexp.MustCompile(`WWW:(\d)&`)
		matches := parser.FindSubmatch(value)
		if matches == nil {
			return false, errors.New("could not determine status from string: " + string(value))
		}

		return string(matches[1]) == "1", nil
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// SetInternet requests the device enable/disable its internet access and
// returns the result state.
func (rpc *RPC) SetInternet(ctx context.Context, state bool) (bool, error) {
	if rpc.transport == nil {
		return false, rpcWrapper.ErrTransportNotConnected
	}

	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		if state {
			request = "MCU+PAS+RAKOIT:WWW:1&"
		} else {
			request = "MCU+PAS+RAKOIT:WWW:0&"
		}
		replyPrefix = "MCU+PAS+RAKOIT:WWW:"
	}

	if request == "" {
		return false, rpcWrapper.ErrUnknownTransportFlavor
	}

	returnChan := make(chan []byte)
	defer close(returnChan)
	rpc.transport.RegisterOneshotReader(replyPrefix, returnChan)
	sendErr := rpc.transport.SendMessage(ctx, request)
	if sendErr != nil {
		return false, sendErr
	}

	select {
	case value := <-returnChan:
		parser := regexp.MustCompile(`WWW:(\d)&`)
		matches := parser.FindSubmatch(value)
		if matches == nil {
			return false, errors.New("could not determine status from string: " + string(value))
		}

		return string(matches[1]) == "1", nil
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// GetEthernet queries if the device has an ethernet connection.
func (rpc *RPC) GetEthernet(ctx context.Context) (bool, error) {
	if rpc.transport == nil {
		return false, rpcWrapper.ErrTransportNotConnected
	}

	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:ETH&"
		replyPrefix = "MCU+PAS+RAKOIT:ETH:"
	}

	if request == "" {
		return false, rpcWrapper.ErrUnknownTransportFlavor
	}

	returnChan := make(chan []byte)
	defer close(returnChan)
	rpc.transport.RegisterOneshotReader(replyPrefix, returnChan)
	sendErr := rpc.transport.SendMessage(ctx, request)
	if sendErr != nil {
		return false, sendErr
	}

	select {
	case value := <-returnChan:
		parser := regexp.MustCompile(`ETH:(\d)&`)
		matches := parser.FindSubmatch(value)
		if matches == nil {
			return false, errors.New("could not determine status from string: " + string(value))
		}

		return string(matches[1]) == "1", nil
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// SetEthernet requests the device enable/disable its ethernet connection and
// returns the result state.
func (rpc *RPC) SetEthernet(ctx context.Context, state bool) (bool, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		if state {
			request = "MCU+PAS+RAKOIT:ETH:1&"
		} else {
			request = "MCU+PAS+RAKOIT:ETH:0&"
		}
		replyPrefix = "MCU+PAS+RAKOIT:ETH:"
	}

	data, reqErr := atomicRequestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return false, reqErr
	}
	parser := regexp.MustCompile(`ETH:(\d)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return false, errors.New("could not determine status from string: " + string(data))
	}

	return string(matches[1]) == "1", nil
}

// GetWifi queries if the device has an wifi connection.
func (rpc *RPC) GetWifi(ctx context.Context) (bool, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:WIF&"
		replyPrefix = "MCU+PAS+RAKOIT:WIF:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return false, reqErr
	}
	if request == "" {
		return false, rpcWrapper.ErrUnknownTransportFlavor
	}

	parser := regexp.MustCompile(`WIF:(\d)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return false, errors.New("could not determine status from string: " + string(data))
	}

	return string(matches[1]) == "1", nil
}

// SetWifi requests the device enable/disable its wifi connection and
// returns the result state.
func (rpc *RPC) SetWifi(ctx context.Context, state bool) (bool, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		if state {
			request = "MCU+PAS+RAKOIT:WIF:1&"
		} else {
			request = "MCU+PAS+RAKOIT:WIF:0&"
		}
		replyPrefix = "MCU+PAS+RAKOIT:WIF:"
	}

	data, reqErr := atomicRequestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return false, reqErr
	}

	parser := regexp.MustCompile(`WIF:(\d)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return false, errors.New("could not determine status from string: " + string(data))
	}

	return string(matches[1]) == "1", nil
}

// RequestWifiReset requests the device restart its wifi stack.
func (rpc *RPC) RequestWifiReset(ctx context.Context) error {
	if rpc.transport == nil {
		return rpcWrapper.ErrTransportNotConnected
	}

	request := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:WRS&"
	}

	if request == "" {
		return rpcWrapper.ErrUnknownTransportFlavor
	}

	return rpc.transport.SendMessage(ctx, request)
}

// GetBluetooth queries if the device has a bluetooth connection.
func (rpc *RPC) GetBluetooth(ctx context.Context) (bool, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:BTC&"
		replyPrefix = "MCU+PAS+RAKOIT:BTC:"
	}
	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)

	if reqErr != nil {
		return false, reqErr
	}

	parser := regexp.MustCompile(`BTC:(\d)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return false, errors.New("could not determine status from string: " + string(data))
	}

	return string(matches[1]) == "1", nil
}

// SetBluetooth requests the device enable/disable its bluetooth connection
func (rpc *RPC) SetBluetooth(ctx context.Context, state bool) error {
	if rpc.transport == nil {
		return rpcWrapper.ErrTransportNotConnected
	}

	request := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		if state {
			request = "MCU+PAS+RAKOIT:BTC:1&"
		} else {
			request = "MCU+PAS+RAKOIT:BTC:0&"
		}
	}

	if request == "" {
		return rpcWrapper.ErrUnknownTransportFlavor
	}

	return rpc.transport.SendMessage(ctx, request)
}
