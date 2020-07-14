package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	c := getCode("edwin@testing.com")
	fmt.Println(c)

	c = getCode("edwi@testing.com")
	fmt.Println(c)
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("some_key"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}