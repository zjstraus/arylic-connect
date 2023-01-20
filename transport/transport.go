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

// Package transport defines interfaces for Arylic APIs that are shared across
// multiple connection mechanisms.
package transport

import (
	"context"
)

// InterfaceFlavor is an enum used to differentiate which interface implementation
// is in use, as commands take different forms between connection types.
type InterfaceFlavor int

const (
	Flavor_TCP InterfaceFlavor = iota // UART over a TCP connection to a Linkplay module
	Flavor_HTTP
)

// AsyncLine is an interface for a family of connections using a string request -
// reply protocol. (Direct UART and TCP tunneled UART)
//
// Unless otherwise specified an implementation of AsyncLine should be safe for
// use by multiple threads.
//
// Actual commands on this transport take the form of [prefix]{command}[:param], so
// the general flow of implementation is to register a listener for the expected
// return prefix and command for your request, then make the request.
type AsyncLine interface {
	// Connect establishes a connection to a given target
	Connect(target string) error

	// RegisterPersistentReader sets up a channel to receive a message off the line
	// every time the given prefix is received.
	RegisterPersistentReader(prefix string, channel chan<- []byte)
	UnregisterPersistentReader(prefix string, channel chan<- []byte)

	// RegisterOneshotReader sets up a channel to receive a message off the line
	// the first time a given prefix is received.
	// Returns true if there was already at least one reader queued for that prefix
	RegisterOneshotReader(prefix string, channel chan<- []byte) bool

	SendMessageAtomic(ctx context.Context, message string, prefix string, outchan chan<- []byte) error

	// SendMessage puts a message out on the connection.
	SendMessage(ctx context.Context, message string) error

	// Flavor returns the InterfaceFlavor corresponding to this implementation
	// so that it can be used to switch command formats.
	Flavor() InterfaceFlavor

	Close() error

	// Target returns the connection target string.
	Target() string
}

type HTTP interface {
	Connect(target string) error

	MakeRequest(ctx context.Context, command string, params ...string) ([]byte, error)

	Flavor() InterfaceFlavor

	Close() error
	Target() string
}
