package main

import (
	"golang.org/x/crypto/bcrypt"
)

func getHash(message string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(message), 11)
    if err != nil {
    	panic(err)
    }
	return string(hash)
}

func encrypt(message string) string {
	pass := getHash(message)
	return pass
}

func main() {
	encrypt(input)
}
