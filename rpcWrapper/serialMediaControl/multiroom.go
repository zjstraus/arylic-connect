package serialMediaControl

import (
	"arylic-connect/transport"
	"context"
	"errors"
	"regexp"
)

type MultiroomMode int

const (
	Mode_None MultiroomMode = iota
	Mode_Master
	Mode_Slave
)

func (mode MultiroomMode) MarshalText() ([]byte, error) {
	switch mode {
	case Mode_None:
		return []byte("None"), nil
	case Mode_Master:
		return []byte("Master"), nil
	case Mode_Slave:
		return []byte("Slave"), nil
	default:
		return []byte("Unknown"), errors.New("unknown multiroom mode")
	}
}

func (mode MultiroomMode) marshallApiText() ([]byte, error) {
	switch mode {
	case Mode_None:
		return []byte("N"), nil
	case Mode_Master:
		return []byte("M"), nil
	case Mode_Slave:
		return []byte("S"), nil
	default:
		return []byte("Unknown"), errors.New("unknown multiroom mode")
	}
}

func (mode *MultiroomMode) UnmarshalText(text []byte) error {
	stringed := string(text)
	switch stringed {
	case "None":
		*mode = Mode_None
	case "Slave":
		*mode = Mode_Slave
	case "Master":
		*mode = Mode_Master
	default:
		*mode = Mode_None
		return errors.New("unknown multiroom mode")
	}
	return nil
}

func (mode *MultiroomMode) unmarshalApiText(text []byte) error {
	stringed := string(text)
	switch stringed {
	case "N":
		*mode = Mode_None
	case "S":
		*mode = Mode_Slave
	case "M":
		*mode = Mode_Master
	default:
		*mode = Mode_None
		return errors.New("unknown multiroom mode")
	}
	return nil
}

type ChannelConfig int

const (
	Channel_Stereo ChannelConfig = iota
	Channel_Left
	Channel_Right
)

func (channel ChannelConfig) MarshalText() ([]byte, error) {
	switch channel {
	case Channel_Stereo:
		return []byte("Stereo"), nil
	case Channel_Left:
		return []byte("Left"), nil
	case Channel_Right:
		return []byte("Right"), nil
	default:
		return []byte("Unknown"), errors.New("unknown multiroom channel")
	}
}

func (channel ChannelConfig) marshallApiText() ([]byte, error) {
	switch channel {
	case Channel_Stereo:
		return []byte("S"), nil
	case Channel_Left:
		return []byte("L"), nil
	case Channel_Right:
		return []byte("R"), nil
	default:
		return []byte("Unknown"), errors.New("unknown multiroom channel")
	}
}

func (channel *ChannelConfig) UnmarshalText(text []byte) error {
	stringed := string(text)
	switch stringed {
	case "None":
		*channel = Channel_Stereo
	case "Slave":
		*channel = Channel_Left
	case "Master":
		*channel = Channel_Right
	default:
		*channel = Channel_Stereo
		return errors.New("unknown multiroom channel")
	}
	return nil
}

func (channel *ChannelConfig) unmarshalApiText(text []byte) error {
	stringed := string(text)
	switch stringed {
	case "S":
		*channel = Channel_Stereo
	case "L":
		*channel = Channel_Left
	case "R":
		*channel = Channel_Right
	default:
		*channel = Channel_Stereo
		return errors.New("unknown multiroom channel")
	}
	return nil
}

func (rpc *RPC) GetMultiroomMode(ctx context.Context) (MultiroomMode, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:MRM&"
		replyPrefix = "MCU+PAS+RAKOIT:MRM:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return Mode_None, reqErr
	}

	parser := regexp.MustCompile(`MRM:(\w+)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return Mode_None, errors.New("could not determine mode from string: " + string(data))
	}
	var mode MultiroomMode

	return mode, mode.unmarshalApiText(matches[1])
}

func (rpc *RPC) GetChannelConfig(ctx context.Context) (ChannelConfig, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:CHN&"
		replyPrefix = "MCU+PAS+RAKOIT:CHN:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return Channel_Stereo, reqErr
	}

	parser := regexp.MustCompile(`CHN:(\w+)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return Channel_Stereo, errors.New("could not determine mode from string: " + string(data))
	}
	var mode ChannelConfig

	return mode, mode.unmarshalApiText(matches[1])
}

func (rpc *RPC) SetVolumeSync(ctx context.Context, state bool) (bool, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		if state {
			request = "MCU+PAS+RAKOIT:VOS:1&"
		} else {
			request = "MCU+PAS+RAKOIT:VOS:0&"
		}
		replyPrefix = "MCU+PAS+RAKOIT:VOS:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return false, reqErr
	}

	parser := regexp.MustCompile(`VOS:(\d)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return false, errors.New("could not determine status from string: " + string(data))
	}

	return string(matches[1]) == "1", nil
}

func (rpc *RPC) GetVolumeSync(ctx context.Context) (bool, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:VOS&"
		replyPrefix = "MCU+PAS+RAKOIT:VOS:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return false, reqErr
	}

	parser := regexp.MustCompile(`VOS:(\d)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return false, errors.New("could not determine status from string: " + string(data))
	}

	return string(matches[1]) == "1", nil
}
