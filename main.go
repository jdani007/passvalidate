package main

import (
	"errors"
	"fmt"
	"unicode"
)

const passLength = 8

func validate(pw string) error {
	pass := []rune(pw)

	if len(pass) < passLength {
		return errors.New("does not meet password length requirements")
	} 
	return complexChecker(pass)
}

func complexChecker(pw []rune) error {

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
		return errors.New("password does not meet the required complexity")
	}

	return nil
}

func main(){

	if err := validate("This!i5a"); err != nil {
		fmt.Println(err)
	}
}

