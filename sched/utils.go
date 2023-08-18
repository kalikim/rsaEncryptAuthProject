package sched

//import (
//	"crypto/rand"
//	"crypto/rsa"
//	"crypto/x509"
//	"encoding/pem"
//	"fmt"
//	"os"
//	"path/filepath"
//)
//
//func GetMainDirPath() (string, error) {
//	ex, err := os.Executable()
//	if err != nil {
//		return "", err
//	}
//
//	exPath := filepath.Dir(ex)
//	mainDirPath := filepath.Join(exPath, "..")
//
//	mainDirPath, err = filepath.Abs(mainDirPath)
//	if err != nil {
//		return "", err
//	}
//
//	return mainDirPath, nil
//}
//
//func EncryptUserAuth(path string, password string) ([]byte, error) {
//	input := []byte(password)
//
//	certFile := path
//	certData, err := os.ReadFile(certFile)
//	if err != nil {
//		fmt.Println("Error reading certificate file:", err)
//		return nil, err
//	}
//
//	block, _ := pem.Decode(certData)
//
//	cert, _ := x509.ParseCertificate(block.Bytes)
//	publicKey, ok := cert.PublicKey.(*rsa.PublicKey)
//	if !ok {
//		fmt.Println("Failed to extract RSA public key")
//		return nil, err
//	}
//
//	encryptedBytes, _ := rsa.EncryptPKCS1v15(rand.Reader, publicKey, input)
//
//	return encryptedBytes, nil
//}
