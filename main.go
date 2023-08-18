package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

var mainDirPath string

func EncryptUserAuth(path string, password string) ([]byte, error) {

	// Load pem file
	pub, err := os.ReadFile(path)
	if err != nil {
		log.Println("Error reading certificate file:", err)
		return nil, err
	}

	// Decode public key struct from pem string
	block, _ := pem.Decode([]byte(pub))
	if block == nil {
		log.Println("Failed to decode PEM block")
		return nil, err
	}

	// Parse the certificate
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Println("Cannot parse the public key:", err)
		return nil, err
	}

	// Get the public key
	rsaPublicKey := cert.PublicKey.(*rsa.PublicKey)

	// Encrypt password using the public key
	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		rsaPublicKey,
		[]byte(password),
		nil)
	if err != nil {
		log.Println("Cannot encrypt the password given the public key:", err)
		return nil, err
	}

	return encryptedBytes, nil

}

func main() {

	password := "2018!!"
	certfilePathabs := "C:\\Users\\kalik\\GolandProjects\\rsaEncryptAuthProject\\base64.cer"
	log.Println(certfilePathabs)
	encryptedPass, _ := EncryptUserAuth(certfilePathabs, password)

	// Print the encrypted bytes as hexadecimal values
	for _, b := range encryptedPass {
		fmt.Printf("%02X", b)
	}
}
