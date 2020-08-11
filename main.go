package main

import (
	"fmt"
	"flag"

	"github.com/james-daniels/passvalidate/check"
)

var passlen int
var hashlen int
var password string

func init(){
	flag.IntVar(&passlen, "len", 8, "Enter the password length (256, 384, 512).")
	flag.IntVar(&hashlen, "hash", 256, "Enter the SHA2 Hash length.")
	flag.StringVar(&password, "pass", "", "Enter complex password.")
}

func main(){
	flag.Parse()

	if err := check.Pass(password, passlen); err != nil {
		fmt.Println(err)
	} else {
		fmt.Print(check.Hash(password, hashlen))
	} 
}
