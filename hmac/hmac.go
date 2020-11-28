package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"log"
)

func main() {
	message := []byte("some random message")
	key := []byte{}
	for i := 0; i < sha512.Size; i++ {
		key = append(key, byte(i))
	}

	fmt.Printf("Message: %x\n", message)
	fmt.Printf("Key: %x\n", key)

	content, err := createHMAC(message, key)
	if err != nil {
		log.Fatalf("Main error: %s", err)
	}

	fmt.Printf("HMAC signature: %x\n", content)

	// hmac.Equal(messageHMAC, expectedHMAC)
}

func createHMAC(message, key []byte) ([]byte, error) {
	mac := hmac.New(sha512.New, key)
	_, err := mac.Write(message)
	if err != nil {
		return nil, fmt.Errorf("Error creating hmac message: %x", message)
	}

	return mac.Sum(nil), nil
}
