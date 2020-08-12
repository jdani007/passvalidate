# passvalidate
Currently reading [The Go Workshop](https://courses.packtpub.com/courses/go)

Here's my spin on [Exercise 3.01: Program to Measure Password Complexity](https://github.com/PacktWorkshops/The-Go-Workshop/blob/master/Chapter03/Exercise03.01/main.go)

Added SHA256 hash generation and command line options.

Added MongoDB support with insert and find operations.

### Usage:
Build the binary
```Bash
$ go build -o passvalidate
```

Run the program
```Bash
$ ./passvalidate --pass This%i5a
4b3c0857acebe278e38ef315e7656cd74dffc25ee8e25d7a10214443d3b7725a
```

Default password length is 8.

SHA2 options are 256 (Default), 384, and 512.

```GO
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
```