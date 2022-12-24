package main

import (
	"arylic-connect/rpcWrapper/serialMediaControl"
	"arylic-connect/transport/tcp"
	"context"
	"log"
)

func main() {
	transport, _ := tcp.New()
	transportErr := transport.Connect("Living.apt.horsegunwhiskyfist.com:8899")
	if transportErr != nil {
		panic(transportErr)
	}
	defer transport.Close()

	//catchallChan := make(chan []byte)
	//go func() {
	//	for {
	//		select {
	//		case msg := <-catchallChan:
	//			log.Println(string(msg))
	//		}
	//	}
	//}()
	//
	//transport.RegisterPersistentReader("", catchallChan)
	//
	//writeErr := transport.SendMessage(context.Background(), "MCU+PAS+RAKOIT:VER&")
	////writeErr := transport.SendMessage(context.Background(), "MCU+PAS+EQGet&")
	//if writeErr != nil {
	//	panic(writeErr)
	//}

	rpc := serialMediaControl.New(transport)
	version, versionErr := rpc.GetVersion(context.Background())
	if versionErr != nil {
		panic(versionErr)
	}
	log.Printf("%v", version)
}
