package transport

import "context"

type InterfaceFlavor int

const (
	Flavor_TCP InterfaceFlavor = iota
)

type AsyncLine interface {
	Connect(target string) error
	RegisterPersistentReader(prefix string, channel chan<- []byte)
	RegisterOneshotReader(prefix string, channel chan<- []byte)
	SendMessage(ctx context.Context, message string) error
	Flavor() InterfaceFlavor
	Close() error
}
