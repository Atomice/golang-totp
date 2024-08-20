package totp

import (
	"fmt"
)

type TOTPImpl struct {
	Username string `json:username`
	Password string `json:password`
	TOTP     string `json:totp`
	Key      string `json:key`
}

func NewTOTPImpl(username string, password string, key string) *TOTPImpl {
	return &TOTPImpl{
		Username: username,
		Password: password,
		Key:      key,
		TOTP:     "",
	}
}


func (impl *TOTPImpl) GenerateTOTP() {
	fmt.Println("DO SOME CALCULATING AND STORE VALUE")
}