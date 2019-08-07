

package main

//import "fmt"
//timeモジュールの機能を使う
import "time"
import "html/template"
import "os"
import "log"
func main() {
	//現在の時間を代入、:=は関数内のみ使える

	//曜日、月を日本語変換
	//week := [7]string{"日曜日", "月曜日", "火曜日", "水曜日", "木曜日", "金曜日", "土曜日"}
	//month := [13]string{"","1月","2月","3月","4月","５月","6月","7月","8月","9月","10月","11月","12月"}
    //t := time.Now()
    //a := "/"
    //b := "日"
    //t1 := t.AddDate(0,0,-1 * t.Day())
    //t2 := t.AddDate(0,0,-1 * t.Day() + 1)
    //t3 := t.AddDate(0,1,-1 * t.Day())

    //fmt.Println(t1)   
    //fmt.Println(t2)
    //fmt.Println(t3)
    //fmt.Println(t.Add(24 * time.Hour))

    //fmt.Println(t.Year(),a,month[t.Month()],t.Day(),a,week[t.Weekday()])   
    //fmt.Println(month[t.Month()])   
    //fmt.Println(t.Day(),b)  
    //fmt.Println(week[t.Weekday()]) 
    //mt.Println(t.Hour())    
    //fmt.Println(t.Minute())  
    //fmt.Println(t.Second())  
    tt := template.Must(template.ParseFiles("./index.html"))

	// 値の埋め込み
	t := time.Now()
	if err := tt.Execute(os.Stdout, t); err != nil {
		log.Fatal(err)
	}
}
