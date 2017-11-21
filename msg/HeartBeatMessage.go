package msg

import(
	"golang.org/x/net/websocket"
	"../helper"
	"../models"
)

func init(){

}

func HandleHeartBeat() {
    for{
        for cmb,ws := range helper.OnlineUsers{
          
			reply := models.Message{
                Prefix:"SPR",
                From: models.User{},
                Hub: "",
                Channel: "",
                To: models.User{} ,
                Body: "汪",
                Time: helper.Now(),
                Nostr: "",
                CSC: 0,
            }

            if err := websocket.JSON.Send(ws, reply); err != nil {
                delete(helper.OnlineUsers,cmb)
            }
        }
        helper.Delay(15)
    }   
}