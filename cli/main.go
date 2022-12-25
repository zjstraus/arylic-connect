package main

import "arylic-connect/websocketApi"

func main() {
	err := websocketApi.RunWebsocketServer()
	if err != nil {
		panic(err)
	}
}
