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
