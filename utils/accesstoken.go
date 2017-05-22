package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"errors"
	"log"
)

// tick := time.Tick(7 * time.Second)
//const refreshTimeout = 3 * time.Second
const refreshLimit = 30 * time.Minute

const refreshTimeout = 10 * time.Minute
const tokenURL = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"

type accessToken struct {
	AccessToken string `json:"access_token"` //	获取到的凭证
	ExpiresIn   int    `json:"expires_in"`   //	凭证有效时间，单位：秒
	AccessIn    int
	mutex       sync.RWMutex
}

// AccessToken 取最新的 access_token，必须使用这个接口取，内部已经加锁,需要先调用RefreshAccessToken,
var AccessToken = func() (string, error) {
	return "", errors.New("invoke RefreshAccessToken first")
}

// RefreshAccessToken 生成闭包函数
func RefreshAccessToken(appId, appSecret string) {
	// 内部变量，外部不可以调用
	var _token = &accessToken{}

	url := fmt.Sprintf(tokenURL, appId, appSecret)
	new := refresh(url)

	_token.mutex.Lock()
	_token.AccessToken = new.AccessToken
	_token.ExpiresIn = new.ExpiresIn
	_token.AccessIn = int(time.Now().Unix())
	_token.mutex.Unlock()

	AccessToken = func() (string, error) {
		_token.mutex.RLock()
		defer _token.mutex.RUnlock()
		return _token.AccessToken, nil
	}

	go func() {
		//time.Sleep(5 * time.Second)
		tick := time.Tick(refreshTimeout)
		for {
			//如果刷新的时间太频繁则不调用refresh
			if int(time.Now().Unix())-_token.AccessIn > int(refreshLimit) {
				new := refresh(url)
				log.Printf("old access token %v\n", _token)
				log.Printf("new access token %v\n", new)
				_token.mutex.Lock()
				_token.AccessToken = new.AccessToken
				_token.ExpiresIn = new.ExpiresIn
				_token.AccessIn = int(time.Now().Unix())
				_token.mutex.Unlock()
			} else {
				log.Println("ExpiresIn is too long", _token.ExpiresIn-(int(time.Now().Unix())-_token.AccessIn))
			}

			<-tick // 等待下一个时钟周期到来
		}
	}()
}

func refresh(url string, ns ...int) (new *accessToken) {
	n := 0
	if len(ns) > 0 {
		n = ns[0]
	}

	var err error
	defer func() {
		if err != nil {
			log.Println(err)
			time.Sleep(3 * time.Minute)
			if n < 9 {
				n++
				new = refresh(url, n)
			}
		}
	}()

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	resp.Body.Close()

	new = &accessToken{}
	err = json.Unmarshal(body, new)
	if err != nil {
		return
	}

	return new
}
