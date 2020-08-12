package main

import (
	"fmt"
	"flag"

	"github.com/james-daniels/passvalidate/pw"
	"github.com/james-daniels/passvalidate/db"
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

	user := db.User{
		Firstname: "John",
		Lastname: "Doe",
		Username: "jdoe",
		Password: pw.Hash(password, hashlen),
		Email: "john.doe@hotmail.com",
	}

	if err := pw.Check(password, passlen); err != nil {
		fmt.Println(err)
		return 
	}

	if err := user.InsertCreds(); err != nil {
		fmt.Println(err)
		return
	}

	creds, err := user.FindCreds()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(creds.Email)
	fmt.Println(creds.Password)
}