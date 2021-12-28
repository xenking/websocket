package main

import (
	"fmt"

	"github.com/valyala/fasthttp"

	"github.com/xenking/websocket"
)

func OnMessage(c *websocket.Conn, isBinary bool, data []byte) {
	c.Write(data)
}

func main() {
	wS := websocket.Server{}
	wS.HandleData(OnMessage)

	s := fasthttp.Server{
		Handler: wS.Upgrade,
	}

	fmt.Println(s.ListenAndServe(":8081"))
}
