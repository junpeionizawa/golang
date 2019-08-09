
package main

import (
    "fmt"
    "net/http"
    "html/template"
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

type Person struct {
    Name string
    From string
}

func Handler(w http.ResponseWriter, r *http.Request) {

    p := Person{
        Name:"hogehoge",
        From : "千葉",
    }

    tmpl := template.Must(template.ParseFiles("./view/index.html"))
    tmpl.Execute(w, p)
}