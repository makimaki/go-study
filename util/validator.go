package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"log"
)

func Validate(channelSecret string, signature string, body string) bool {
	log.Printf(`
signature: %+v
body: %s
`, signature, body)

	mac := hmac.New(sha256.New, []byte(channelSecret))
	mac.Write([]byte(body))
	computed := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return signature == computed
}
