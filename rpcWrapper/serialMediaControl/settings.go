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
	"arylic-connect/transport"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// SetLED requests the device enable/disable any LEDs and returns the
// result state.
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

// GetLED queries the device to see if LEDs are enabled
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

// ToggleLED inverts the devices current LED enabled state and returns it.
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

// SetBeep requests the device enable/disable audible feedback when
// physical controls are used and returns the result state.
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

// GetBeep queries the device to see if audible feedback is enabled.
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

// GetName queries the device for its name.
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

// SetName requests the device change its name and returns the result
// state.
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

// SetVoicePrompt requests the device enable/disable any voice
// prompts and returns the result state.
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

// GetVoicePrompt queries the device to see if voice prompts are enabled
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
