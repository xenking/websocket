package main

import (
	"github.com/valyala/fasthttp"

	"github.com/xenking/websocket"
)

func main() {
	ws := websocket.Server{}

	ws.HandleData(wsHandler)

	fasthttp.ListenAndServe(":9000", ws.Upgrade)
}

func wsHandler(c *websocket.Conn, isBinary bool, data []byte) {
	c.Write(data)
}
