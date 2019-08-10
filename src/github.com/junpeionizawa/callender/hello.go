
package main

import (
	"fmt"
    "net/http"
    "html/template"
    //"strings"
    "time"
)
func main() {
    fmt.Println("The Server runs with http://localhost:3000/")
    // ここで静的ファイルであるCSSを適用
    // Handler : 第一引数で与えたパターンに対して、第二引数Handlerを登録
    // StripPrefix : 第一引数で与えたURLのパスを取り除いて、第二引数Handlerを発動させる
    // FileServer : 、HTTPリクエストに対して、第一引数のrootを起点とするファイルシステムのコンテンツを返すハンドラを返す
    // ここでいうrootは、resources

    // つまり、resourcesディレクトリ以下の静的ファイル(ここでいうcss/index.css)を探して返す
    http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources/"))))
    http.HandleFunc("/", Handler)
    http.ListenAndServe(":3000", nil)
}

type Time struct {
    Year int
    Month int
    Day1 int
    Day2 interface{}
    Day3 interface{}
    Day4 interface{}
    Day5 interface{}
    Day6 interface{}
    Day7 interface{}
    Weekday string
}

func Handler(w http.ResponseWriter, r *http.Request) {

    t := Time{
        Year:time.Now().Year(),
        Month:int(time.Now().Month()),
        Day1:time.Now().Day(),
        Day2:time.Now().Add(24 * time.Hour).Day(),
        Day3 :time.Now().Add(48 * time.Hour).Day(),
        Day4 :time.Now().Add(72 * time.Hour).Day(),
        Day5 :time.Now().Add(96 * time.Hour).Day(),
        Day6 :time.Now().Add(120 * time.Hour).Day(),
        Day7 :time.Now().Add(144 * time.Hour).Day(),
        Weekday :time.Now().Weekday().String(),
    }



    tmpl := template.Must(template.ParseFiles("./resources/index.html"))
    tmpl.Execute(w, t)

} 