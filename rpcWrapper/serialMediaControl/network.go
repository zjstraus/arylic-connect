package serialMediaControl

import (
	"arylic-connect/rpcWrapper"
	"arylic-connect/transport"
	"context"
	"errors"
	"regexp"
)

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

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
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

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
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
