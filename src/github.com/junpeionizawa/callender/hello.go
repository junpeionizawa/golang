

package main

import "fmt"
//timeモジュールの機能を使う
import "time"

func main() {
	//現在の時間を代入、:=は関数内のみ使える

	//曜日、月を日本語変換
	week := [...]string{"日曜日", "月曜日", "火曜日", "水曜日", "木曜日", "金曜日", "土曜日"}
	month := [...]string{"","1月","2月","3月","4月","５月","6月","7月","8月","9月","10月","11月","12月"}
    t := time.Now()
    fmt.Println(t)           
    fmt.Println(t.Year())   
    fmt.Println(month[t.Month()])   
    fmt.Println(t.Day())   
    fmt.Println(week[t.Weekday()]) 
    fmt.Println(t.Hour())    
    fmt.Println(t.Minute())  
    fmt.Println(t.Second())  
}