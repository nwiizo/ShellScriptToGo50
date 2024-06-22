package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const passphrase = "your-secret-passphrase"

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run script.go {encrypt|decrypt} directory")
		return
	}

	action := os.Args[1]
	directory := os.Args[2]

	switch action {
	case "encrypt":
		encryptDirectory(directory)
	case "decrypt":
		decryptDirectory(directory)
	default:
		fmt.Println("Invalid action. Use 'encrypt' or 'decrypt'.")
	}
}

func encryptDirectory(dir string) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) != ".enc" {
			encryptFile(path)
		}
		return nil
	})
}

func decryptDirectory(dir string) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".enc" {
			decryptFile(path)
		}
		return nil
	})
}

func encryptFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		return
	}

	block, err := aes.NewCipher(generateKey(passphrase))
	if err != nil {
		fmt.Printf("Error creating cipher for %s: %v\n", filename, err)
		return
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Printf("Error creating GCM for %s: %v\n", filename, err)
		return
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Printf("Error generating nonce for %s: %v\n", filename, err)
		return
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	err = os.WriteFile(filename+".enc", ciphertext, 0644)
	if err != nil {
		fmt.Printf("Error writing encrypted file %s: %v\n", filename, err)
		return
	}

	os.Remove(filename)
	fmt.Printf("Encrypted: %s\n", filename)
}

func decryptFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		return
	}

	block, err := aes.NewCipher(generateKey(passphrase))
	if err != nil {
		fmt.Printf("Error creating cipher for %s: %v\n", filename, err)
		return
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Printf("Error creating GCM for %s: %v\n", filename, err)
		return
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		fmt.Printf("Ciphertext too short for %s\n", filename)
		return
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Printf("Error decrypting %s: %v\n", filename, err)
		return
	}

	decryptedFilename := filename[:len(filename)-4] // Remove .enc extension
	err = os.WriteFile(decryptedFilename, plaintext, 0644)
	if err != nil {
		fmt.Printf("Error writing decrypted file %s: %v\n", decryptedFilename, err)
		return
	}

	os.Remove(filename)
	fmt.Printf("Decrypted: %s\n", decryptedFilename)
}

func generateKey(passphrase string) []byte {
	hash := sha256.Sum256([]byte(passphrase))
	return hash[:]
}
