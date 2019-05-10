package helpers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/astaxie/beego"
	"math/big"
	"os"
	"time"
)

func CreateCertificateIfNotExists() {
	certFileName := "conf/beemail.crt"
	keyFileName := "conf/beemail.key"
	if !CheckIfFileExists(certFileName) || !CheckIfFileExists(keyFileName) {
		ca := createCertificateAuthority()
		privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
		publicKey := &privateKey.PublicKey
		cert, err := x509.CreateCertificate(rand.Reader, ca, ca, publicKey, privateKey)
		CheckError(err)
		certificateFile, err := os.Create(certFileName)
		err = pem.Encode(certificateFile, &pem.Block{Type: "CERTIFICATE", Bytes: cert})
		CheckError(err)
		err = certificateFile.Close()
		CheckError(err)
		beego.Info("Created new cert file " + certFileName)

		key, err := os.OpenFile(keyFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		CheckError(err)
		err = pem.Encode(key, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})
		CheckError(err)
		err = key.Close()
		CheckError(err)
		beego.Info("Created new key file " + keyFileName)
	}
}

func generateRandomBigNumber() *big.Int {
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(1024))
	CheckError(err)
	return randomNumber
}

func createCertificateAuthority() *x509.Certificate {
	ca := &x509.Certificate{
		SerialNumber:          generateRandomBigNumber(),
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}
	return ca
}
