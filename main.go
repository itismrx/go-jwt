package main

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	bcrypt.Cost([]byte(""))
	log.Println("Hello World")
}
