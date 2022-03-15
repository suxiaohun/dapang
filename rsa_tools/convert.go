package rsa_tools

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
)

func ExportRsaPrivateKeyAsPemStr(privkey *rsa.PrivateKey) string {
	privkey_bytes := x509.MarshalPKCS1PrivateKey(privkey)
	privkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkey_bytes,
		},
	)
	return string(privkey_pem)
}

func ParseRsaPrivateKeyFromPemStr(privPEM string) (*rsa.PrivateKey, error) {
	if !strings.Contains(privPEM, "-----") {
		privPEM = strToPem(privPEM)
	}

	block, _ := pem.Decode([]byte(privPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

func ExportRsaPublicKeyAsPemStr(pubkey *rsa.PublicKey) (string, error) {
	pubkey_bytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", err
	}
	pubkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkey_bytes,
		},
	)

	return string(pubkey_pem), nil
}

func ParseRsaPublicKeyFromPemStr(pubPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		break // fall through
	}
	return nil, errors.New("Key type is not RSA")
}

func VerifyKeyPair(private, public string) (bool,error) {
	if !strings.Contains(private, "-----") {
		private = strToPem(private)
	}
	if !strings.Contains(public, "-----") {
		public = strToPem(public)
	}

	priKey, err := ParseRsaPrivateKeyFromPemStr(private)
	if err != nil {
		fmt.Printf("parse private key error: %v \n", err)
		return false, err
	}
	pubKey, err := ParseRsaPublicKeyFromPemStr(public)
	if err != nil {
		fmt.Printf("parse public key error: %v \n", err)
		return false, err
	}

	return priKey.PublicKey.Equal(pubKey),nil
}

func strToPem(str string) string {
	pemStr := "-----BEGIN RSA PRIVATE KEY-----"
	pemStr += "\n"

	result := splitByWidth(str, 64)

	for _, s := range result {
		pemStr += s
		pemStr += "\n"
	}

	pemStr += "-----END RSA PRIVATE KEY-----"
	pemStr += "\n"
	return pemStr
}

func splitByWidth(str string, size int) []string {
	strLength := len(str)
	var splited []string
	var stop int
	for i := 0; i < strLength; i += size {
		stop = i + size
		if stop > strLength {
			stop = strLength
		}
		splited = append(splited, str[i:stop])
	}
	return splited
}


