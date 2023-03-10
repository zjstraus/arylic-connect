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
)

type InputSource int

const (
	Input_Net InputSource = iota
	Input_Usb
	Input_UsbDac
	Input_LineIn1
	Input_LineIn2
	Input_Bluetooth
	Input_Optical
	Input_Coax
	Input_I2s
	Input_Hdmi
	Input_None
	Input_Unknown
)

func (source InputSource) MarshalText() ([]byte, error) {
	switch source {
	case Input_Net:
		return []byte("Network"), nil
	case Input_Usb:
		return []byte("USB"), nil
	case Input_UsbDac:
		return []byte("USBDAC"), nil
	case Input_LineIn1:
		return []byte("Line-In"), nil
	case Input_LineIn2:
		return []byte("Line-In2"), nil
	case Input_Bluetooth:
		return []byte("Bluetooth"), nil
	case Input_Optical:
		return []byte("Optical"), nil
	case Input_Coax:
		return []byte("Coax"), nil
	case Input_I2s:
		return []byte("I2S"), nil
	case Input_Hdmi:
		return []byte("HDMI"), nil
	case Input_None:
		return []byte("None"), nil
	default:
		return []byte("Unknown"), nil
	}
}

func (source InputSource) marshallApiText() ([]byte, error) {
	switch source {
	case Input_Net:
		return []byte("NET"), nil
	case Input_Usb:
		return []byte("USB"), nil
	case Input_UsbDac:
		return []byte("USBDAC"), nil
	case Input_LineIn1:
		return []byte("LINE-IN"), nil
	case Input_LineIn2:
		return []byte("LINE-IN2"), nil
	case Input_Bluetooth:
		return []byte("BT"), nil
	case Input_Optical:
		return []byte("OPT"), nil
	case Input_Coax:
		return []byte("COAX"), nil
	case Input_I2s:
		return []byte("I2S"), nil
	case Input_Hdmi:
		return []byte("HDMI"), nil
	case Input_None:
		return []byte("NONE"), nil
	default:
		return nil, errors.New("no API string available for input")
	}
}

func (source *InputSource) UnmarshalText(text []byte) error {
	stringed := string(text)
	switch stringed {
	case "Network":
		*source = Input_Net
	case "USB":
		*source = Input_Usb
	case "USBDAC":
		*source = Input_UsbDac
	case "Line-In":
		*source = Input_LineIn1
	case "Line-In2":
		*source = Input_LineIn2
	case "Bluetooth":
		*source = Input_Bluetooth
	case "Optical":
		*source = Input_Optical
	case "Coax":
		*source = Input_Coax
	case "I2S":
		*source = Input_I2s
	case "HDMI":
		*source = Input_Hdmi
	case "None":
		*source = Input_None
	default:
		*source = Input_Unknown
	}
	return nil
}

func (source *InputSource) unmarshalApiText(text []byte) error {
	stringed := string(text)
	switch stringed {
	case "NET":
		*source = Input_Net
	case "USB":
		*source = Input_Usb
	case "USBDAC":
		*source = Input_UsbDac
	case "LINE-IN":
		*source = Input_LineIn1
	case "LINE-IN2":
		*source = Input_LineIn2
	case "BT":
		*source = Input_Bluetooth
	case "OPT":
		*source = Input_Optical
	case "COAX":
		*source = Input_Coax
	case "I2S":
		*source = Input_I2s
	case "HDMI":
		*source = Input_Hdmi
	case "NONE":
		*source = Input_None
	default:
		*source = Input_Unknown
		return errors.New("no API string available")
	}
	return nil
}

// GetSource queries the device for its current active source.
func (rpc *RPC) GetSource(ctx context.Context) (InputSource, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:SRC&"
		replyPrefix = "MCU+PAS+RAKOIT:SRC:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return Input_Unknown, reqErr
	}

	parser := regexp.MustCompile(`SRC:(\w+)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return Input_Unknown, errors.New("could not determine input from string: " + string(data))
	}
	var source InputSource

	return source, source.unmarshalApiText(matches[1])
}

// SetSource requests the device change its active source and returns
// the result.
//
// Input_None is not a valid source for use here.
func (rpc *RPC) SetSource(ctx context.Context, targetSource InputSource) (InputSource, error) {
	request := ""
	replyPrefix := ""
	formattedSource, sourceFormatErr := targetSource.marshallApiText()
	if sourceFormatErr != nil {
		return Input_Unknown, sourceFormatErr
	}
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = fmt.Sprintf("MCU+PAS+RAKOIT:SRC:%s&", formattedSource)
		replyPrefix = "MCU+PAS+RAKOIT:SRC:"
	}

	data, reqErr := atomicRequestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return Input_Unknown, reqErr
	}

	parser := regexp.MustCompile(`SRC:(\w+)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return Input_Unknown, errors.New("could not determine input from string: " + string(data))
	}
	var source InputSource

	return source, source.unmarshalApiText(matches[1])
}

// GetDefaultSource queries the device for what source will be
// active at power on, if any.
//
// A setting of Input_None signals that the device will power on to
// whatever source it had last.
func (rpc *RPC) GetDefaultSource(ctx context.Context) (InputSource, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:POM&"
		replyPrefix = "MCU+PAS+RAKOIT:POM:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return Input_Unknown, reqErr
	}

	parser := regexp.MustCompile(`POM:(\w+)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return Input_Unknown, errors.New("could not determine input from string: " + string(data))
	}
	var source InputSource

	return source, source.unmarshalApiText(matches[1])
}

// SetDefaultSource requests the device update what source will be
// active at power on, if any and returns the result.
//
// A setting of Input_None signals that the device will power on to
// whatever source it had last.
func (rpc *RPC) SetDefaultSource(ctx context.Context, targetSource InputSource) (InputSource, error) {
	request := ""
	replyPrefix := ""
	formattedSource, sourceFormatErr := targetSource.marshallApiText()
	if sourceFormatErr != nil {
		return Input_Unknown, sourceFormatErr
	}
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = fmt.Sprintf("MCU+PAS+RAKOIT:POM:%s&", formattedSource)
		replyPrefix = "MCU+PAS+RAKOIT:POM:"
	}

	data, reqErr := atomicRequestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return Input_Unknown, reqErr
	}

	parser := regexp.MustCompile(`POM:(\w+)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return Input_Unknown, errors.New("could not determine input from string: " + string(data))
	}
	var source InputSource

	return source, source.unmarshalApiText(matches[1])
}

// GetInputAutoswitch queries the device to see if automatically
// switching to new valid inputs is enabled.
func (rpc *RPC) GetInputAutoswitch(ctx context.Context) (bool, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:ASW&"
		replyPrefix = "MCU+PAS+RAKOIT:ASW:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return false, reqErr
	}

	parser := regexp.MustCompile(`ASW:(\d)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return false, errors.New("could not determine status from string: " + string(data))
	}

	return string(matches[1]) == "1", nil
}

// SetInputAutoswitch requests the device change if automatic
// switching is enabled and returns the result.
func (rpc *RPC) SetInputAutoswitch(ctx context.Context, state bool) (bool, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		if state {
			request = "MCU+PAS+RAKOIT:ASW:1&"
		} else {
			request = "MCU+PAS+RAKOIT:ASW:0&"
		}
		replyPrefix = "MCU+PAS+RAKOIT:ASW:"
	}

	data, reqErr := atomicRequestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return false, reqErr
	}

	parser := regexp.MustCompile(`ASW:(\d)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return false, errors.New("could not determine status from string: " + string(data))
	}

	return string(matches[1]) == "1", nil
}
