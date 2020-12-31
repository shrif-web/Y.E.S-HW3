package controller

import (
	"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(pwd []byte) (string,error) {
	if hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost);
	err != nil {
		return "", err
	}else{
		return string(hash),nil
	}
}
