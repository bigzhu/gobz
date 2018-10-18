package rsabz

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"

	"github.com/pkg/errors"
)

// Sign 生成签名
func Sign(plainText string, privateKey string) (signatur string, err error) {
	p := []byte(plainText)
	hash := sha256.New()
	hash.Write(p)
	parsedPrivateKey, err := getPrivateKey(privateKey)
	if err != nil {
		return
	}
	s, err := rsa.SignPKCS1v15(rand.Reader, parsedPrivateKey, crypto.SHA256, hash.Sum(nil))
	if err != nil {
		return
	}
	signatur = base64.StdEncoding.EncodeToString(s)
	return
}

func getPemBytes(key string) (b []byte, err error) {
	p := []byte(key)
	thePem, _ := pem.Decode(p)
	if thePem == nil {
		return nil, errors.New("RSA PRIVATE KEY not found")
	}
	if thePem.Type != "RSA PRIVATE KEY" && thePem.Type != "PUBLIC KEY" && thePem.Type != "RSA PUBLIC KEY" {
		return nil, errors.Errorf("类型不是有效的RSA KEY: %v", thePem.Type)
	}
	b = thePem.Bytes
	return
}

func getPrivateKey(privateKey string) (p *rsa.PrivateKey, err error) {
	pemBytes, err := getPemBytes(privateKey)
	if err != nil {
		return
	}
	return x509.ParsePKCS1PrivateKey(pemBytes)
}
