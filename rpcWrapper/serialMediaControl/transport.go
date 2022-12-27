package serialMediaControl

import (
	"arylic-connect/rpcWrapper"
	"arylic-connect/transport"
	"context"
	"errors"
	"fmt"
	"regexp"
)

func (rpc *RPC) RequestPlayPause(ctx context.Context) error {
	if rpc.transport == nil {
		return rpcWrapper.ErrTransportNotConnected
	}

	request := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:POP&"
	}

	if request == "" {
		return rpcWrapper.ErrUnknownTransportFlavor
	}

	return rpc.transport.SendMessage(ctx, request)
}

func (rpc *RPC) RequestNext(ctx context.Context) error {
	if rpc.transport == nil {
		return rpcWrapper.ErrTransportNotConnected
	}

	request := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:NXT&"
	}

	if request == "" {
		return rpcWrapper.ErrUnknownTransportFlavor
	}

	return rpc.transport.SendMessage(ctx, request)
}

func (rpc *RPC) RequestPrevious(ctx context.Context) error {
	if rpc.transport == nil {
		return rpcWrapper.ErrTransportNotConnected
	}

	request := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:PRE&"
	}

	if request == "" {
		return rpcWrapper.ErrUnknownTransportFlavor
	}

	return rpc.transport.SendMessage(ctx, request)
}

func (rpc *RPC) RequestStop(ctx context.Context) error {
	if rpc.transport == nil {
		return rpcWrapper.ErrTransportNotConnected
	}

	request := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:STP&"
	}

	if request == "" {
		return rpcWrapper.ErrUnknownTransportFlavor
	}

	return rpc.transport.SendMessage(ctx, request)
}

type LoopMode int

const (
	Loop_RepeatAll LoopMode = iota
	Loop_RepeatOne
	Loop_RepeatShuffle
	Loop_Shuffle
	Loop_Sequence
)

func (mode LoopMode) MarshalText() ([]byte, error) {
	switch mode {
	case Loop_RepeatAll:
		return []byte("Repeat All"), nil
	case Loop_RepeatOne:
		return []byte("Repeat One"), nil
	case Loop_RepeatShuffle:
		return []byte("Repeat & Shuffle"), nil
	case Loop_Shuffle:
		return []byte("Shuffle"), nil
	case Loop_Sequence:
		return []byte("Sequence"), nil
	default:
		return []byte("Unknown"), errors.New("unknown loop mode")
	}
}

func (mode LoopMode) marshallApiText() ([]byte, error) {
	switch mode {
	case Loop_RepeatAll:
		return []byte("REPEATALL"), nil
	case Loop_RepeatOne:
		return []byte("REPEATONE"), nil
	case Loop_RepeatShuffle:
		return []byte("REPEATSHUFFLE"), nil
	case Loop_Shuffle:
		return []byte("SHUFFLE"), nil
	case Loop_Sequence:
		return []byte("SEQUENCE"), nil
	default:
		return []byte("Unknown"), errors.New("unknown loop mode")
	}
}

func (mode *LoopMode) UnmarshalText(text []byte) error {
	stringed := string(text)
	switch stringed {
	case "Repeat All":
		*mode = Loop_RepeatAll
	case "Repeat One":
		*mode = Loop_RepeatOne
	case "Repeat & Shuffle":
		*mode = Loop_RepeatShuffle
	case "Shuffle":
		*mode = Loop_Shuffle
	case "Sequence":
		*mode = Loop_Sequence
	default:
		*mode = Loop_Sequence
		return errors.New("unknown multiroom mode")
	}
	return nil
}

func (mode *LoopMode) unmarshalApiText(text []byte) error {
	stringed := string(text)
	switch stringed {
	case "REPEATALL":
		*mode = Loop_RepeatAll
	case "REPEATONE":
		*mode = Loop_RepeatOne
	case "REPEATSHUFFLE":
		*mode = Loop_RepeatShuffle
	case "SHUFFLE":
		*mode = Loop_Shuffle
	case "SEQUENCE":
		*mode = Loop_Sequence
	default:
		*mode = Loop_Sequence
		return errors.New("unknown loop mode")
	}
	return nil
}

func (rpc *RPC) GetLoopMode(ctx context.Context) (LoopMode, error) {
	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:LPM&"
		replyPrefix = "MCU+PAS+RAKOIT:LPM:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return Loop_Sequence, reqErr
	}

	parser := regexp.MustCompile(`LPM:(\w+)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return Loop_Sequence, errors.New("could not determine mode from string: " + string(data))
	}
	var mode LoopMode

	return mode, mode.unmarshalApiText(matches[1])
}

func (rpc *RPC) SetLoopMode(ctx context.Context, targetMode LoopMode) (LoopMode, error) {
	request := ""
	replyPrefix := ""
	formattedMode, modeFormatErr := targetMode.marshallApiText()
	if modeFormatErr != nil {
		return Loop_Sequence, modeFormatErr
	}
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = fmt.Sprintf("MCU+PAS+RAKOIT:LPM:%s&", formattedMode)
		replyPrefix = "MCU+PAS+RAKOIT:LPM:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return Loop_Sequence, reqErr
	}

	parser := regexp.MustCompile(`LPM:(\w+)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return Loop_Sequence, errors.New("could not determine loop mode from string: " + string(data))
	}
	var mode LoopMode

	return mode, mode.unmarshalApiText(matches[1])
}
