

package main

import "fmt"
//timeモジュールの機能を使う
import "time"

func main() {
	//現在の時間を代入、:=は関数内のみ使える
    t := time.Now()
    fmt.Println(t)           
    fmt.Println(t.Year())   
    fmt.Println(t.Month())   
    fmt.Println(t.Day())   
    fmt.Println(t.Hour())    
    fmt.Println(t.Minute())  
    fmt.Println(t.Second())  
    fmt.Println(t.Weekday()) 
}