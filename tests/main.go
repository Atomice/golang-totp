package main

import (
	"fmt"

	"github.com/Atomice/golang-totp/totp"
)

func main() {
	totp.NewTOTPImpl("test", "test", "test")
	fmt.Println("test")
}