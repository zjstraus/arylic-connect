package serialMediaControl

import (
	"arylic-connect/transport"
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func (rpc *RPC) GetVolume(ctx context.Context) (float32, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:VOL&"
		replyPrefix = "MCU+PAS+RAKOIT:VOL:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return 0, reqErr
	}

	parser := regexp.MustCompile(`VOL:(\d+)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return 0, errors.New("could not determine volume from string: " + string(data))
	}
	parsed, parseErr := strconv.Atoi(string(matches[1]))
	if parseErr != nil {
		return 0, parseErr
	}

	if parsed == 0 {
		return 0, nil
	}

	floatCast := float32(parsed)

	return floatCast / 100, nil
}

func (rpc *RPC) SetVolume(ctx context.Context, state float32) (float32, error) {
	request := ""
	replyPrefix := ""
	formattedState := int(state * 100)
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = fmt.Sprintf("MCU+PAS+RAKOIT:VOL:%d&", formattedState)
		replyPrefix = "MCU+PAS+RAKOIT:VOL:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return 0, reqErr
	}

	parser := regexp.MustCompile(`VOL:(\d+)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return 0, errors.New("could not determine volume from string: " + string(data))
	}
	parsed, parseErr := strconv.Atoi(string(matches[1]))
	if parseErr != nil {
		return 0, parseErr
	}

	floatCast := float32(parsed)

	return floatCast / 100, nil
}

func (rpc *RPC) SetMute(ctx context.Context, state bool) (bool, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		if state {
			request = "MCU+PAS+RAKOIT:MUT:1&"
		} else {
			request = "MCU+PAS+RAKOIT:MUT:0&"
		}
		replyPrefix = "MCU+PAS+RAKOIT:MUT:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return false, reqErr
	}

	parser := regexp.MustCompile(`MUT:(\d)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return false, errors.New("could not determine status from string: " + string(data))
	}

	return string(matches[1]) == "1", nil
}

func (rpc *RPC) GetMute(ctx context.Context) (bool, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:MUT&"
		replyPrefix = "MCU+PAS+RAKOIT:MUT:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return false, reqErr
	}

	parser := regexp.MustCompile(`MUT:(\d)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return false, errors.New("could not determine status from string: " + string(data))
	}

	return string(matches[1]) == "1", nil
}

func (rpc *RPC) ToggleMute(ctx context.Context) (bool, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:MUT:T&"
		replyPrefix = "MCU+PAS+RAKOIT:MUT:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return false, reqErr
	}

	parser := regexp.MustCompile(`MUT:(\d)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return false, errors.New("could not determine status from string: " + string(data))
	}

	return string(matches[1]) == "1", nil
}

func (rpc *RPC) SetFixedVolume(ctx context.Context, state bool) (bool, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		if state {
			request = "MCU+PAS+RAKOIT:VOF:1&"
		} else {
			request = "MCU+PAS+RAKOIT:VOF:0&"
		}
		replyPrefix = "MCU+PAS+RAKOIT:VOF:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return false, reqErr
	}

	parser := regexp.MustCompile(`VOF:(\d)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return false, errors.New("could not determine status from string: " + string(data))
	}

	return string(matches[1]) == "1", nil
}

func (rpc *RPC) GetFixedVolume(ctx context.Context) (bool, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:VOF&"
		replyPrefix = "MCU+PAS+RAKOIT:VOF:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return false, reqErr
	}

	parser := regexp.MustCompile(`VOF:(\d)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return false, errors.New("could not determine status from string: " + string(data))
	}

	return string(matches[1]) == "1", nil
}

func (rpc *RPC) GetMaxVolume(ctx context.Context) (float32, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:MXV&"
		replyPrefix = "MCU+PAS+RAKOIT:MXV:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return 0, reqErr
	}

	parser := regexp.MustCompile(`MXV:(\d+)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return 0, errors.New("could not determine volume from string: " + string(data))
	}
	parsed, parseErr := strconv.Atoi(string(matches[1]))
	if parseErr != nil {
		return 0, parseErr
	}

	if parsed == 0 {
		return 0, nil
	}

	floatCast := float32(parsed)

	return floatCast / 100, nil
}

func (rpc *RPC) SetMaxVolume(ctx context.Context, state float32) (float32, error) {
	request := ""
	replyPrefix := ""
	formattedState := int(state * 100)
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = fmt.Sprintf("MCU+PAS+RAKOIT:MXV:%d&", formattedState)
		replyPrefix = "MCU+PAS+RAKOIT:MXV:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return 0, reqErr
	}

	parser := regexp.MustCompile(`MXV:(\d+)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return 0, errors.New("could not determine volume from string: " + string(data))
	}
	parsed, parseErr := strconv.Atoi(string(matches[1]))
	if parseErr != nil {
		return 0, parseErr
	}

	floatCast := float32(parsed)

	return floatCast / 100, nil
}

func (rpc *RPC) GetBalance(ctx context.Context) (float32, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:BAL&"
		replyPrefix = "MCU+PAS+RAKOIT:BAL:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return 0, reqErr
	}

	parser := regexp.MustCompile(`BAL:(\d+)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return 0, errors.New("could not determine volume from string: " + string(data))
	}
	parsed, parseErr := strconv.Atoi(string(matches[1]))
	if parseErr != nil {
		return 0, parseErr
	}

	if parsed == 0 {
		return 0, nil
	}

	floatCast := float32(parsed)

	return (floatCast / 100) - 1, nil
}

func (rpc *RPC) SetBalance(ctx context.Context, state float32) (float32, error) {
	request := ""
	replyPrefix := ""
	formattedState := int((state + 1) * 100)
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = fmt.Sprintf("MCU+PAS+RAKOIT:BAL:%d&", formattedState)
		replyPrefix = "MCU+PAS+RAKOIT:BAL:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return 0, reqErr
	}

	parser := regexp.MustCompile(`BAL:(\d+)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return 0, errors.New("could not determine volume from string: " + string(data))
	}
	parsed, parseErr := strconv.Atoi(string(matches[1]))
	if parseErr != nil {
		return 0, parseErr
	}

	floatCast := float32(parsed)

	return (floatCast / 100) - 1, nil
}
