package util

import (
	"crypto/sha1"
	"fmt"
)

func EncryptSHA1(text string, passSault ...string) string {
	if len(passSault) > 0 && passSault[0] != "" {
		text += passSault[0]
	}
	h := sha1.New()
	h.Write([]byte(text))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
