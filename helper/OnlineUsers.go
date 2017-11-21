package helper

import(
	"golang.org/x/net/websocket"
)

var OnlineUsers map[string](*websocket.Conn)

func init(){
	 OnlineUsers = make(map[string](*websocket.Conn))
}