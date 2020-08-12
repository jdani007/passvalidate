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
	flag.IntVar(&passlen, "len", 8, "Enter the password length.")
	flag.IntVar(&hashlen, "hash", 256, "Enter the SHA2 Hash length (256, 384, 512).")
	flag.StringVar(&password, "pass", "", "Enter complex password.")
}

func main(){
	flag.Parse()

	mongo := db.Mongo{
		Database: "user",
		Collection: "users",
		URI: "mongodb://127.0.0.1:27017",
	}

	if err := pw.Check(password, passlen); err != nil {
		fmt.Println(err)
		return
	}

	user := db.User{
		Firstname: "John",
		Lastname: "Doe",
		Username: "jdoe",
		Password: pw.Hash(password, hashlen),
		Email: "john.doe@hotmail.com",
	}

	if _, err := user.InsertCreds(mongo); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user.FindCreds("username",user.Username,mongo).Email)
	fmt.Println(user.FindCreds("username",user.Username,mongo).Password)
}