package msg

import(
	"golang.org/x/net/websocket"
	"../helper"
	"fmt"
	"../users"
	"../models"
)

func init(){
	
}

func MessageHubHandler(ws *websocket.Conn) {

    var err error
 
    for {
 
        var reply models.Message

        key :=ws.Request().Header.Get("Sec-WebSocket-Key")

        token :=ws.Request().Header.Get("Sec-Websocket-Protocol")

        cmbKey := token+"-"+key

		if _, ok := helper.OnlineUsers[cmbKey]; ok {  
			 helper.OnlineUsers[cmbKey] = ws
		}  else{
			  user:=users.QueryUser(token)
			  if user.UserId >0 {
			  		helper.OnlineUsers[cmbKey] = ws
			  		fmt.Println(user.UserName,user.NickName)
			  } else {
			  		msg :=  models.Message{
				        Prefix:"SR",
				        From: models.User{},
				        Hub: "",
				        Channel: "",
				        To: models.User{},
				        Body: "Error Token",
				        Time: "",
				        Nostr: "",
				        CSC: 0,
				    }
			  		websocket.JSON.Send(ws, msg);
			  		ws.Close()
			  }
		}

        if err = websocket.JSON.Receive(ws, &reply); err != nil {

            delete(helper.OnlineUsers, cmbKey)
            break
        }
 
        switch(reply.Prefix){
            case "UM":
                HandleUserMessage(ws,reply,cmbKey)
                break
            case "MM":
                HandleMulticast(ws,reply,cmbKey)
                break
                 
        }
    }
}