package main

import (
	"log"
	"fmt"
	"crypto/cipher"
	"crypto/aes"
	"bytes"
	"golang.org/x/crypto/bcrypt"
) 

func main() {
	key, err := bcrypt.GenerateFromPassword([]byte("random key"), bcrypt.MinCost)
	if err != nil {
		log.Fatal("Error bcrypt password generation")
	}
	
	key = key[:aes.BlockSize]
	msg := "Random message"

	cipherMsg, err := encrypt(key, []byte(msg))
	if err != nil {
		log.Fatal("Error cipher operation")
	} 

	fmt.Println(string(cipherMsg))

	uncipherMsg, err := decrypt(key, []byte(cipherMsg))
	if err != nil {
		log.Fatal("Error uncipher operation")
	} 

	fmt.Println(string(uncipherMsg))
	
}

func encrypt(key []byte, msg []byte) ([]byte, error) {
	b, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("Error creating cipher")
	}

	iv := make([]byte, aes.BlockSize)
	strm := cipher.NewCTR(b, iv)

	buff := &bytes.Buffer{}
	sw := cipher.StreamWriter {
		S: strm, 
		W: buff, 
	}

	_, err = sw.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("Could not write block %s", err)
	}
	return buff.Bytes(), nil

}

func decrypt(key []byte, msg []byte) ([]byte, error){
	return encrypt(key, msg)
}