package utils

import (
	"testing"
)

func TestAccessToken(t *testing.T) {

	appID := "wxe0a5a0ed81342a7c"
	appSecret := "798fd2b81dfd477825cad8089321719c"

	RefreshAccessToken(appID, appSecret)
	token, err := AccessToken()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(token)
	//time.Sleep(15 * time.Second)
}
