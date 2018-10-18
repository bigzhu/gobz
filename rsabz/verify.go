package rsabz

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"

	"github.com/pkg/errors"
)

// Verify 验证签名
// plainText 内容
// signatur 签名
// publicKey 公钥
func Verify(plainText string, signatur string, publicKey string) (err error) {
	parsedPublicKey, err := getPublicKey(publicKey)
	if err != nil {
		return
	}
	signBytes, err := base64.StdEncoding.DecodeString(signatur)
	if err != nil {
		err = errors.Wrap(err, "base64.StdEncoding.DecodeString signBase64")
		return
	}

	hashed := sha256.Sum256([]byte(plainText))
	return rsa.VerifyPKCS1v15(parsedPublicKey, crypto.SHA256, hashed[:], signBytes)
}

func getPublicKey(publicKey string) (p *rsa.PublicKey, err error) {
	pemBytes, err := getPemBytes(publicKey)
	if err != nil {
		return
	}
	keyInterface, err := x509.ParsePKIXPublicKey(pemBytes)
	if err != nil {
		return
	}
	p, ok := keyInterface.(*rsa.PublicKey)
	if !ok {
		err = errors.Errorf("Could not cast parsed key to *rsa.PublickKey")
	}
	return
}
