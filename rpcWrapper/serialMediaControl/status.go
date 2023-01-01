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
	"arylic-connect/rpcWrapper"
	"arylic-connect/transport"
	"context"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// EndpointStatus is a large struct the API returns all at once. It
// has a lot of entries and a vague definition of what's included
// at what API levels, so it also includes a slice of field names
// to show what items were parsed.
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

// GetStatus queries the device for its current status summary.
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
