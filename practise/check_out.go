package main 

import (    
    "fmt"
    "time"
)

func main(){
    var hours int = 8
    now := time.Now().Local().Add(time.Hour * time.Duration(hours))
    fmt.Println(now.Format("Go out of office by 15:04:05 PM"))
}