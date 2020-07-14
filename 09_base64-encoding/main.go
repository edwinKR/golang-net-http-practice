package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	t := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

	t64 := base64.StdEncoding.EncodeToString([]byte(t))
	fmt.Println("====Encoded String======", t64)

	bs, err := base64.StdEncoding.DecodeString(t64)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Decoded byte slice>>>>>>>", bs)
	fmt.Println("Decoded string version >>>>>>> ", string(bs))
}