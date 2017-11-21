package users

import(
 	"database/sql"
    _ "github.com/go-sql-driver/mysql"
    "../helper"
    "../models"
)

var DB *sql.DB
var err error

func OpenDB(connStr string){
	DB, err = sql.Open("mysql", connStr)
    err = DB.Ping() 
    if err != nil {
        panic("database connect failure.")     
    }
}

func CloseDB(){
	DB.Close()
}

func Status(){
	for{
		tokens :=[]string{}

		for _,ws := range helper.OnlineUsers{
			token := ws.Request().Header.Get("Sec-Websocket-Protocol")
			tokens = append(tokens,token)
		}

		RefreshUserStatus(tokens)
		helper.Delay(15)
	}
}


func QueryUser(token string) (models.User) {

    row := DB.QueryRow("SELECT id,username,nickname FROM users where token = ?",token)

    var id int
    var userName string
    var nickName string

    row.Scan(&id,&userName,&nickName);

    user := models.User {
        UserId :id,
        UserName :userName,
        NickName :nickName,
    } 
        
    return  user;
}

func RefreshUserStatus(tokens []string) {

	stmt, _ := DB.Prepare(`DELETE FROM websocket_online`)

    stmt.Exec()

	if len(tokens) > 0 {
		
		for _,token := range tokens{

			stmt, _ = DB.Prepare(`INSERT INTO websocket_online (guid,username,nickname,user_id,created_at) SELECT token,username,nickname,id,NOW() FROM users WHERE token = ?`)

 			stmt.Exec(token)
		}
 	}
}

func OnLine(token string) {

	user:= QueryUser(token)

    stmt, _ := DB.Prepare(`DELETE FROM websocket_online WHERE user_id=?`)

    stmt.Exec(user.UserId)

    stmt,_ = DB.Prepare(`INSERT INTO websocket_online (guid,username,nickname,user_id,created_at) values (?,?,?,?,?)`)

    t:=helper.Now()

    stmt.Exec(token,user.UserName, user.NickName, user.UserId, t)
}

func OffLine(token string) {

    stmt, _ := DB.Prepare(`DELETE FROM websocket_online WHERE guid=?`)

    stmt.Exec(token)
}

func Users() ([](models.User)) {

    var count int;
    var id int
    var userName string
    var nickName string
    var guid string

    row := DB.QueryRow(`select count(user_id) as amount from  websocket_online`)
    row.Scan(&count)

    rows,_ := DB.Query(`select user_id,username,nickname,guid from  websocket_online`)
    users := make([](models.User),0)

    for rows.Next(){
        rows.Scan(&id,&userName,&nickName,&guid);
        user := models.User{
            UserId :id,
            UserName :userName,
            NickName :nickName,
        } 
       users = append(users,user)
    }

    return users;
}