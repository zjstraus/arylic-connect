package serialMediaControl

import (
	"arylic-connect/transport"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func (rpc *RPC) SetLED(ctx context.Context, state bool) (bool, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		if state {
			request = "MCU+PAS+RAKOIT:LED:1&"
		} else {
			request = "MCU+PAS+RAKOIT:LED:0&"
		}
		replyPrefix = "MCU+PAS+RAKOIT:LED:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return false, reqErr
	}

	parser := regexp.MustCompile(`LED:(\d)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return false, errors.New("could not determine status from string: " + string(data))
	}

	return string(matches[1]) == "1", nil
}

func (rpc *RPC) GetLED(ctx context.Context) (bool, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:LED&"
		replyPrefix = "MCU+PAS+RAKOIT:LED:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return false, reqErr
	}

	parser := regexp.MustCompile(`LED:(\d)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return false, errors.New("could not determine status from string: " + string(data))
	}

	return string(matches[1]) == "1", nil
}

func (rpc *RPC) ToggleLED(ctx context.Context) (bool, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:LED:T&"
		replyPrefix = "MCU+PAS+RAKOIT:LED:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return false, reqErr
	}

	parser := regexp.MustCompile(`LED:(\d)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return false, errors.New("could not determine status from string: " + string(data))
	}

	return string(matches[1]) == "1", nil
}

func (rpc *RPC) SetBeep(ctx context.Context, state bool) (bool, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		if state {
			request = "MCU+PAS+RAKOIT:BEP:1&"
		} else {
			request = "MCU+PAS+RAKOIT:BEP:0&"
		}
		replyPrefix = "MCU+PAS+RAKOIT:BEP:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return false, reqErr
	}

	parser := regexp.MustCompile(`BEP:(\d)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return false, errors.New("could not determine status from string: " + string(data))
	}

	return string(matches[1]) == "1", nil
}

func (rpc *RPC) GetBeep(ctx context.Context) (bool, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:BEP&"
		replyPrefix = "MCU+PAS+RAKOIT:BEP:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return false, reqErr
	}

	parser := regexp.MustCompile(`BEP:(\d)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return false, errors.New("could not determine status from string: " + string(data))
	}

	return string(matches[1]) == "1", nil
}

func (rpc *RPC) GetName(ctx context.Context) (string, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:NAM&"
		replyPrefix = "MCU+PAS+RAKOIT:NAM:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return "", reqErr
	}

	parser := regexp.MustCompile(`NAM:(\w+)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return "", errors.New("could not determine name from string: " + string(data))
	}

	decodedHex, decodeErr := hex.DecodeString(string(matches[1]))
	if decodeErr != nil {
		return "", decodeErr
	}

	return string(decodedHex), nil
}

func (rpc *RPC) SetName(ctx context.Context, name string) (string, error) {
	encoded := hex.EncodeToString([]byte(name))
	encoded = strings.ToUpper(encoded)
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = fmt.Sprintf("MCU+PAS+RAKOIT:NAM:%s&", encoded)
		replyPrefix = "MCU+PAS+RAKOIT:NAM:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return "", reqErr
	}

	parser := regexp.MustCompile(`NAM:(\w+)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return "", errors.New("could not determine name from string: " + string(data))
	}

	decodedHex, decodeErr := hex.DecodeString(string(matches[1]))
	if decodeErr != nil {
		return "", decodeErr
	}

	return string(decodedHex), nil
}

func (rpc *RPC) SetVoicePrompt(ctx context.Context, state bool) (bool, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		if state {
			request = "MCU+PAS+RAKOIT:PMT:1&"
		} else {
			request = "MCU+PAS+RAKOIT:PMT:0&"
		}
		replyPrefix = "MCU+PAS+RAKOIT:PMT:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return false, reqErr
	}

	parser := regexp.MustCompile(`PMT:(\d)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return false, errors.New("could not determine status from string: " + string(data))
	}

	return string(matches[1]) == "1", nil
}

func (rpc *RPC) GetVoicePrompt(ctx context.Context) (bool, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:PMT&"
		replyPrefix = "MCU+PAS+RAKOIT:PMT:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return false, reqErr
	}

	parser := regexp.MustCompile(`PMT:(\d)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return false, errors.New("could not determine status from string: " + string(data))
	}

	return string(matches[1]) == "1", nil
}
