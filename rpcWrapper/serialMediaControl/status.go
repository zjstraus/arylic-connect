package serialMediaControl

import (
	"arylic-connect/rpcWrapper"
	"arylic-connect/transport"
	"context"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type EndpointStatus struct {
	Source      InputSource
	Mute        bool
	Volume      float32
	Treble      float32
	Bass        float32
	Network     bool
	Internet    bool
	Playing     bool
	Led         bool
	Upgrading   bool
	ValidValues []string
}

func (rpc *RPC) GetStatus(ctx context.Context) (EndpointStatus, error) {
	status := EndpointStatus{}

	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:STA&"
		replyPrefix = "MCU+PAS+RAKOIT:STA:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return status, reqErr
	}

	if request == "" {
		return status, rpcWrapper.ErrUnknownTransportFlavor
	}

	parser := regexp.MustCompile(`STA:([\w,]+)&`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return status, errors.New("could not determine status from string: " + string(data))
	}
	entries := strings.Split(string(matches[1]), ",")

	if len(entries) >= 1 {
		status.Source.unmarshalApiText([]byte(entries[0]))
		status.ValidValues = append(status.ValidValues, "Source")
	}
	if len(entries) >= 2 {
		status.Mute = entries[1] == "1"
		status.ValidValues = append(status.ValidValues, "Mute")
	}
	if len(entries) >= 3 {
		parsedVal, parseErr := strconv.ParseFloat(entries[2], 32)
		if parseErr == nil {
			status.Volume = float32(parsedVal) / 100
			status.ValidValues = append(status.ValidValues, "Volume")
		}
	}
	if len(entries) >= 4 {
		parsedVal, parseErr := strconv.ParseFloat(entries[3], 32)
		if parseErr == nil {
			status.Treble = float32(parsedVal) / 10
			status.ValidValues = append(status.ValidValues, "Treble")
		}
	}
	if len(entries) >= 5 {
		parsedVal, parseErr := strconv.ParseFloat(entries[4], 32)
		if parseErr == nil {
			status.Bass = float32(parsedVal) / 10
			status.ValidValues = append(status.ValidValues, "Bass")
		}
	}
	if len(entries) >= 6 {
		status.Network = entries[5] == "1"
		status.ValidValues = append(status.ValidValues, "Network")
	}
	if len(entries) >= 7 {
		status.Internet = entries[6] == "1"
		status.ValidValues = append(status.ValidValues, "Internet")
	}
	if len(entries) >= 8 {
		status.Playing = entries[7] == "1"
		status.ValidValues = append(status.ValidValues, "Playing")
	}
	if len(entries) >= 9 {
		status.Led = entries[8] == "1"
		status.ValidValues = append(status.ValidValues, "Led")
	}
	if len(entries) >= 10 {
		status.Upgrading = entries[9] == "1"
		status.ValidValues = append(status.ValidValues, "Upgrading")
	}

	return status, nil
}
