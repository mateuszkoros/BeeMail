package helpers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"github.com/astaxie/beego"
	"math/big"
	"os"
	"time"
)

var (
	cryptographyDirectory = "cryptography"
	certFileName          = cryptographyDirectory + "/BeeMail.crt"
	keyFileName           = cryptographyDirectory + "/BeeMail.key"
)

// CreateCertificateIfNotExists checks if user certificate has been created before.
// If there is no certificate it creates a new one.
func CreateCertificateIfNotExists() {
	if !CheckIfFileExists(certFileName) || !CheckIfFileExists(keyFileName) {
		certificateBytes, privateKey, _ := createCertificate()
		savePublicKey(certificateBytes)
		savePrivateKey(privateKey)
	}
}

// generateRandomBigNumber is used to create a serial number for a new certificate.
func generateRandomBigNumber() *big.Int {
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(1024))
	CheckError(err)
	return randomNumber
}

// createCertificate creates a new certificate and returns private and public key.
func createCertificate() ([]byte, *rsa.PrivateKey, *rsa.PublicKey) {
	certificate := &x509.Certificate{
		SerialNumber: generateRandomBigNumber(),
		Subject: pkix.Name{
			Organization: []string{"BeeMail"},
		},
		DNSNames:    []string{"BeeMail"},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().AddDate(100, 0, 0),
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageDigitalSignature,
	}
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := &privateKey.PublicKey
	certificateBytes, err := x509.CreateCertificate(rand.Reader, certificate, certificate, publicKey, privateKey)
	CheckError(err)
	return certificateBytes, privateKey, publicKey
}

// savePublicKey stores newly generated public key on disk.
func savePublicKey(certificateBytes []byte) {
	CheckError(os.MkdirAll(cryptographyDirectory, 0700))
	certificateFile, err := os.Create(certFileName)
	CheckError(err)
	CheckError(pem.Encode(certificateFile, &pem.Block{Type: "CERTIFICATE", Bytes: certificateBytes}))
	CheckError(certificateFile.Close())
	beego.Info("Created new cert file " + certFileName)
}

// savePrivateKey stores newly generated private key on disk,
func savePrivateKey(privateKey *rsa.PrivateKey) {
	CheckError(os.MkdirAll(cryptographyDirectory, 0700))
	keyFile, err := os.OpenFile(keyFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	CheckError(err)
	CheckError(pem.Encode(keyFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}))
	CheckError(keyFile.Close())
	beego.Info("Created new key file " + keyFileName)
}
