package helper

import(
	"time"
)

func Now() (string) {
   return  time.Now().Format("2006-01-02 15:04:05")
}


func Delay(sec time.Duration){
	us := 1000000*1000*sec; 
	time.Sleep(us)
}
