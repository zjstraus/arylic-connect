package serialMediaControl

import "arylic-connect/transport"

type RPC struct {
	transport       transport.AsyncLine
	endpointVersion EndpointVersion
}

func New(t transport.AsyncLine) *RPC {
	return &RPC{transport: t}
}
