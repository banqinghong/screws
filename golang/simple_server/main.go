package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Println("参数解析失败：", err)
		io.WriteString(w, err.Error())
		return
	}
	name := req.Form.Get("name")
	str := fmt.Sprintf("hello, %s", name)
	io.WriteString(w, str)
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Println("server start successful")
}
