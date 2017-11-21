package main

import (
    "golang.org/x/net/websocket"
    "fmt"
    "net/http"
    "./msg"
    "./configs"
    "./users"
)

var CONFIGS map[string]string
var CONNSTR = ""


func init(){
    CONFIGS = configs.Load()
    CONNSTR = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",CONFIGS["dbuser"],CONFIGS["dbpwd"],CONFIGS["dbhost"],CONFIGS["dbport"],CONFIGS["dbname"])
    users.OpenDB(CONNSTR)
}

func main() {

    go msg.HandleHeartBeat()

    go users.Status()

    http.Handle("/hub", websocket.Handler(msg.MessageHubHandler))

    defer users.CloseDB()

    if(CONFIGS["protocol"] == "HTTPS"){
        RunHttpTLS()
    }else{
        RunHttp()
    }
}

func  RunHttpTLS(){
    err := http.ListenAndServeTLS(fmt.Sprintf("%s:%s",CONFIGS["host"],CONFIGS["port"]),"server.crt", "server.key", nil)
    CatchError(err)
}

func  RunHttp(){
     err := http.ListenAndServe(fmt.Sprintf("%s:%s",CONFIGS["host"],CONFIGS["port"]), nil)
     CatchError(err)
}

func CatchError(err error){
    if err != nil {
        fmt.Println(err.Error())
    } 
    // panic("ListenAndServe: " + err.Error())     
}

