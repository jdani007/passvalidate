package check

import (
	"errors"
	"fmt"
	"unicode"
	"crypto/sha256"
	"crypto/sha512"
)

//Pass performs the password length and complexity validation.
func Pass(pw string, plen int) error {

	pass := []rune(pw)

	if len(pass) < plen {
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

//Hash checks the SHA2 hash for the password.
func Hash(pw string, hlen int) string {

		switch hlen {
		case 256:
			return fmt.Sprintf("%x", sha256.Sum256([]byte(pw)))
		case 384:
			return fmt.Sprintf("%x", sha512.Sum384([]byte(pw)))
		case 512:
			return fmt.Sprintf("%x", sha512.Sum512([]byte(pw)))
		default:
			return "Only 256 (Default), 384, and 512 for SHA2 hash."
		}
}