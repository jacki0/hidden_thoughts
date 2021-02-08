package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"io"
	"strings"
)

func getHash(message string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(message), 11)
    if err != nil {
    	panic(err)
    }
	return string(hash)
}

func removeBase64Padding(value string) string {
	return strings.Replace(value, "=", "", -1)
}

func Pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func encrypt(message string) (string, string, string, string) {
	hash := getHash(message)
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	msg := Pad([]byte(message))
	ciphertext := make([]byte, aes.BlockSize+len(msg))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(msg))
	finalMsg := removeBase64Padding(base64.URLEncoding.EncodeToString(ciphertext))
	res := strings.Split(hash, ".")
	if strings.Contains(res[0], "/"){
		res = strings.Split(res[0], "/")
	}
	return finalMsg, hash, string(key), res[0]
}


func main() {
	encryptedMessage, hash, key, pass := encrypt(input)
}
