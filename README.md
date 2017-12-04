# gowebsocket

go websocket proj

支持token 验证

支持单用户多点在线保证收到消息

支持用户上线下线

支持群发消息

支持私聊 

支持 ws  wss  协议 

支持配置服务地址和端口

可扩展支持渠道发送消息

支持docker  (docker build -t websocket .)  

#编译

make 即可  

如果要交叉编译 linux 服务器上的 执行 make linux 

如果要编译 其它平台的 自己补下Makefile

#消息体
```` go
type Message struct 
{
    Prefix string   //消息类型  UM 用户消息，MM  多播消息 ， UT  用户列表，
    From  User  
    To   User 
    Hub    string      // （未实现，目前消息为扁平的，没做层级）
    Channel   string  //  channel  用来区分 消息发送范围（未实现）
    Body   string       // 消息体
    Time  string        
    Nostr  string      //随机字符  （未实现）
    CSC    int          //消息验证串 （未实现）
} 

````
