package main

import (
	"errors"
	"fmt"
	"unicode"
	"crypto/sha256"
)

const passLength = 8
const password = "This%i5a"

func validatePW(pw string) error {

	pass := []rune(pw)

	if len(pass) < passLength {
		return errors.New("does not meet the password length requirements")
	} 

	var hasUpper, hasLower, hasNumber, hasSymbol bool

	for _, v := range pw {

		switch {
		case unicode.IsUpper(v):
			hasUpper = true
		case unicode.IsLower(v):
			hasLower = true
		case unicode.IsNumber(v):
			hasNumber = true
		case unicode.IsPunct(v) || unicode.IsSymbol(v):
			hasSymbol = true
		}
	}

	if !(hasUpper && hasLower && hasNumber && hasSymbol) {
		return errors.New("password does not meet the complexity requirements")
	}
	return nil
}

func hashPW(pw string) string {

	pass := sha256.Sum256([]byte(pw))

	return fmt.Sprintf("%x", pass)
}

func main(){

	if err := validatePW(password); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(hashPW(password))
	} 
	
}
