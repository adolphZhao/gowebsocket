package msg

import(
	"golang.org/x/net/websocket"
	"../helper"
	"fmt"
	"../models"
)

func init(){
	
}

func HandleMulticast(ws *websocket.Conn,request models.Message, key string) (models.Message) {

    reply := models.Message{
        Prefix:"MMR",
        From: models.User{},
        Hub: "",
        Channel: "",
        To: request.From ,
        Body: "OK",
        Time: "",
        Nostr: "",
        CSC: 0,
    }

    if err := websocket.JSON.Send(ws, reply); err != nil {
        fmt.Println(err);
    }

	for otk,ws := range helper.OnlineUsers{
		if err := websocket.JSON.Send(ws, request); err != nil {
        	delete(helper.OnlineUsers,otk)
    	}
	}
    
    return reply

}