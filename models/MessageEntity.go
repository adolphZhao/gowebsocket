package models

type Message struct 
{
    Prefix string
    From  User 
    To   User 
    Hub    string 
    Channel   string 
    Body   string 
    Time  string
    Nostr  string 
    CSC    int    
} 

type User struct {
    UserId int
    UserName string
    NickName string
    Token string
    Index int
}