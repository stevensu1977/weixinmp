package utils

import (
	"fmt"
	"testing"
	"time"
)

var token = "BzJESBGw8xYxMYCpvzm26cH7Qpwh6QR1"
var nonce = "123"
var msg = "HelloWorld"

func TestValidate(t *testing.T) {
	ts := fmt.Sprintf("%d", time.Now().Unix())
	fmt.Println(ts)
	sign := Signature(token, ts, nonce, "")
	if ValidateURL(token, ts, nonce, sign) {
		t.Log("ValidateURL successful")
	} else {
		t.Fatalf("ValidateURL fail, %s\n", sign)
	}

}
