package utils

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"
)

func Signature(token, timestamp, nonce, encrypt string) string {
	tmpArr := []string{token, timestamp, nonce, encrypt}
	sort.Strings(tmpArr)

	tmpStr := strings.Join(tmpArr, "")
	actual := fmt.Sprintf("%x", sha1.Sum([]byte(tmpStr)))
	return actual
}

func CheckSignature(token, timestamp, nonce, encrypt, sign string) bool {
	return Signature(token, timestamp, nonce, encrypt) == sign
}

func ValidateURL(token, timestamp, nonce, signature string) bool {
	return CheckSignature(token, timestamp, nonce, "", signature)
}
