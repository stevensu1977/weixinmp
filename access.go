package weixinmp

import (
	"log"
	"net/http"

	"github.com/stevensu1977/weixinmp/utils"
)

var token = "BzJESBGw8xYxMYCpvzm26cH7Qpwh6QR1"

// HandleAccess 接入微信公众平台开发，并接口来自微信服务器的消息
func HandleAccess(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s", r.URL)

	q := r.URL.Query()
	signature := q.Get("signature")
	timestamp := q.Get("timestamp")
	nonce := q.Get("nonce")

	// 每次都验证 URL，以判断来源是否合法
	if !utils.ValidateURL(token, timestamp, nonce, signature) {
		http.Error(w, "validate url error, request not from weixin?", http.StatusUnauthorized)
		return
	}

	switch r.Method {
	// 如果是 GET 请求，表示这是接入验证请求
	case "GET":
		w.Write([]byte(q.Get("echostr")))
	//case "POST":
	default:
		http.Error(w, "only GET method allowed", http.StatusUnauthorized)
	}
}
