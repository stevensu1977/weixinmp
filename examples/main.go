package main

import (
	"log"
	"net/http"
	"github.com/stevensu1977/weixinmp"
)

func main() {
	addr := ":10080"
	log.Println("Server listen on 10080")
	http.HandleFunc("/weixin",weixinmp.HandleAccess)
	log.Fatal(http.ListenAndServe(addr, nil))
}
