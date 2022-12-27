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
