

package main

import (
	//"bytes"
	//"fmt"
	//"io"
    //"os"
    //"html/template"
    "log"
    "net/http"
    //"time"
)

func main() {
    //ディレクトリを指定する
    fs := http.FileServer(http.Dir("layout"))
    //ルーティング設定。"/"というアクセスがきたらstaticディレクトリのコンテンツを表示させる
    http.Handle("/", fs)

    log.Println("Listening...")
    // 3000ポートでサーバーを立ち上げる
    http.ListenAndServe(":3000", nil)

   
}
