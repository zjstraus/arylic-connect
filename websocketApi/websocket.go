package websocketApi

import (
	"arylic-connect/websocketApi/serialmedia"
	"github.com/ethereum/go-ethereum/rpc"
	"net/http"
)

func RunWebsocketServer() error {
	serialMediaService := serialmedia.New()

	rpcServer := rpc.NewServer()
	serialMediaErr := rpcServer.RegisterName("serialmedia", serialMediaService)
	if serialMediaErr != nil {
		panic(serialMediaErr)
	}

	http.Handle("/ws", rpcServer.WebsocketHandler([]string{"*"}))
	return http.ListenAndServe(":8080", nil)
}
