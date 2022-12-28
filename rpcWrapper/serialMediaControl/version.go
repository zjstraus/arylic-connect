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
	"regexp"
)

type EndpointVersion struct {
	Firmware string
	Git      string
	API      string
}

func (rpc *RPC) GetVersion(ctx context.Context) (EndpointVersion, error) {
	if rpc.endpointVersion.Git != "" {
		return rpc.endpointVersion, nil
	}

	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:VER&"
		replyPrefix = "MCU+PAS+RAKOIT:VER:"
	}

	data, reqErr := requestWithResponse(ctx, rpc.transport, request, replyPrefix)
	if reqErr != nil {
		return EndpointVersion{}, reqErr
	}

	parser := regexp.MustCompile(`VER:(\d+)-(\w+)-(\d+)`)
	matches := parser.FindSubmatch(data)
	if matches == nil {
		return EndpointVersion{}, errors.New("could not determine version from string: " + string(data))
	}
	version := EndpointVersion{
		Firmware: string(matches[1]),
		Git:      string(matches[2]),
		API:      string(matches[3]),
	}
	rpc.endpointVersion = version
	return version, nil
}
