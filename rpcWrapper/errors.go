package rpcWrapper

import "errors"

var (
	ErrTransportNotConnected  = errors.New("transport is not connected")
	ErrUnknownTransportFlavor = errors.New("transport flavor unknown")
)
