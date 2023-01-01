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
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

// GetBass queries the connected device for its current bass EQ setting.
//
// Value will be in the range of [-1,1] with a granularity of 0.1
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

// SetBass requests the connected device change its bass EQ setting and
// returns the result.
//
// Value will be in the range of [-1,1] with a granularity of 0.1
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

// GetTreble queries the connected device for its current treble EQ setting.
//
// Value will be in the range of [-1,1] with a granularity of 0.1
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

// SetTreble requests the connected device change its treble EQ setting and
// returns the result.
//
// Value will be in the range of [-1,1] with a granularity of 0.1
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

// SetVirtualBass requests the connected device change the state of its virtual
// bass booster and returns the result state.
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

// GetVirtualBass queries the connected device for the state of its virtual
// bass booster.
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

// ToggleVirtualBass requests the connected device invert the state of its virtual
// bass booster and returns the result.
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
