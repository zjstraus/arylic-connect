package serialMediaControl

import (
	"arylic-connect/rpcWrapper"
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

	if rpc.transport == nil {
		return EndpointVersion{}, rpcWrapper.ErrTransportNotConnected
	}

	request := ""
	replyPrefix := ""
	switch rpc.transport.Flavor() {
	case transport.Flavor_TCP:
		request = "MCU+PAS+RAKOIT:VER&"
		replyPrefix = "MCU+PAS+RAKOIT:VER:"
	}

	if request == "" {
		return EndpointVersion{}, rpcWrapper.ErrUnknownTransportFlavor
	}

	returnChan := make(chan []byte)
	defer close(returnChan)
	rpc.transport.RegisterOneshotReader(replyPrefix, returnChan)
	rpc.transport.SendMessage(ctx, request)

	select {
	case value := <-returnChan:
		parser := regexp.MustCompile(`VER:(\d+)-(\w+)-(\d+)`)
		matches := parser.FindSubmatch(value)
		if matches == nil {
			return EndpointVersion{}, errors.New("could not determine version from string: " + string(value))
		}
		version := EndpointVersion{
			Firmware: string(matches[1]),
			Git:      string(matches[2]),
			API:      string(matches[3]),
		}
		rpc.endpointVersion = version
		return version, nil
	case <-ctx.Done():
		return EndpointVersion{}, ctx.Err()
	}
}
