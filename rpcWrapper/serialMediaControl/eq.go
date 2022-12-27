package serialMediaControl

import (
	"arylic-connect/transport"
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func (rpc *RPC) GetBass(ctx context.Context) (float32, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:BAS&"
		replyPrefix = "MCU+PAS+RAKOIT:BAS:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)

	if reqErr != nil {
		return 0, reqErr
	}

	parser := regexp.MustCompile(`BAS:(-?\d+)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return 0, errors.New("could not determine bass from string: " + string(data))
	}
	parsed, parseErr := strconv.Atoi(string(matches[1]))
	if parseErr != nil {
		return 0, parseErr
	}

	floatCast := float32(parsed)

	return floatCast / 10, nil
}

func (rpc *RPC) SetBass(ctx context.Context, state float32) (float32, error) {
	request := ""
	replyPrefix := ""
	formattedState := int(state * 10)
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = fmt.Sprintf("MCU+PAS+RAKOIT:BAS:%d&", formattedState)
		replyPrefix = "MCU+PAS+RAKOIT:BAS:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return 0, reqErr
	}
	parser := regexp.MustCompile(`BAS:(-?\d+)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return 0, errors.New("could not determine bass from string: " + string(data))
	}
	parsed, parseErr := strconv.Atoi(string(matches[1]))
	if parseErr != nil {
		return 0, parseErr
	}

	floatCast := float32(parsed)

	return floatCast / 10, nil
}

func (rpc *RPC) GetTreble(ctx context.Context) (float32, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:TRE&"
		replyPrefix = "MCU+PAS+RAKOIT:TRE:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)

	if reqErr != nil {
		return 0, reqErr
	}

	parser := regexp.MustCompile(`TRE:(-?\d+)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return 0, errors.New("could not determine treble from string: " + string(data))
	}
	parsed, parseErr := strconv.Atoi(string(matches[1]))
	if parseErr != nil {
		return 0, parseErr
	}

	floatCast := float32(parsed)

	return floatCast / 10, nil
}

func (rpc *RPC) SetTreble(ctx context.Context, state float32) (float32, error) {
	request := ""
	replyPrefix := ""
	formattedState := int(state * 10)
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = fmt.Sprintf("MCU+PAS+RAKOIT:TRE:%d&", formattedState)
		replyPrefix = "MCU+PAS+RAKOIT:TRE:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return 0, reqErr
	}
	parser := regexp.MustCompile(`TRE:(-?\d+)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return 0, errors.New("could not determine treble from string: " + string(data))
	}
	parsed, parseErr := strconv.Atoi(string(matches[1]))
	if parseErr != nil {
		return 0, parseErr
	}

	floatCast := float32(parsed)

	return floatCast / 10, nil
}

func (rpc *RPC) SetVirtualBass(ctx context.Context, state bool) (bool, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		if state {
			request = "MCU+PAS+RAKOIT:VBS:1&"
		} else {
			request = "MCU+PAS+RAKOIT:VBS:0&"
		}
		replyPrefix = "MCU+PAS+RAKOIT:VBS:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return false, reqErr
	}

	parser := regexp.MustCompile(`VBS:(\d)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return false, errors.New("could not determine status from string: " + string(data))
	}

	return string(matches[1]) == "1", nil
}

func (rpc *RPC) GetVirtualBass(ctx context.Context) (bool, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:VBS&"
		replyPrefix = "MCU+PAS+RAKOIT:VBS:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return false, reqErr
	}

	parser := regexp.MustCompile(`VBS:(\d)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return false, errors.New("could not determine status from string: " + string(data))
	}

	return string(matches[1]) == "1", nil
}

func (rpc *RPC) ToggleVirtualBass(ctx context.Context) (bool, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:VBS:T&"
		replyPrefix = "MCU+PAS+RAKOIT:VBS:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return false, reqErr
	}

	parser := regexp.MustCompile(`VBS:(\d)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return false, errors.New("could not determine status from string: " + string(data))
	}

	return string(matches[1]) == "1", nil
}
