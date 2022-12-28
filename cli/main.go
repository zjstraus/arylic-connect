package main

import "arylic-connect/localWebsocketApi"

func main() {
	err := localWebsocketApi.RunWebsocketServer()
	if err != nil {
		panic(err)
	}
}
