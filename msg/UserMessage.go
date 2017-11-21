package msg

import(
	"golang.org/x/net/websocket"
	"../helper"
	"fmt"
	"regexp"
	"../models"
)

func init(){
	
}

func HandleUserMessage(ws *websocket.Conn,request models.Message,key string) (models.Message) {

    reply := models.Message{
        Prefix:"UMR",
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
		if ok,_:= regexp.MatchString(request.To.Token, otk); ok {
			if err := websocket.JSON.Send(ws, request); err != nil {
            	delete(helper.OnlineUsers,otk)
        	}
		}
	}
    
    return reply
}