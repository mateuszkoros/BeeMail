package helpers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"github.com/astaxie/beego"
	"math/big"
	"os"
	"sync"
	"time"
)

var (
	certificateAuthority  *x509.Certificate
	once                  sync.Once
	rootCACert            = "conf/rootCA.crt"
	rootCAKey             = "conf/rootCA.key"
	cryptographyDirectory = "cryptography"
	certFileName          = cryptographyDirectory + "/BeeMail.crt"
	keyFileName           = cryptographyDirectory + "/BeeMail.key"
)

func CreateCertificateIfNotExists() {
	if !CheckIfFileExists(certFileName) || !CheckIfFileExists(keyFileName) {
		certificateBytes, privateKey, _ := createCertificate()
		savePublicKey(certificateBytes)
		savePrivateKey(privateKey)
	}
}

func GetRootCA() *x509.Certificate {
	once.Do(func() {
		authorityKeyPair, err := tls.LoadX509KeyPair(rootCACert, rootCAKey)
		CheckError(err)
		certificateAuthority, err = x509.ParseCertificate(authorityKeyPair.Certificate[0])
		CheckError(err)
	})
	if certificateAuthority != nil {
		return certificateAuthority
	} else {
		panic("Failed to get certificate authority")
	}
}

func generateRandomBigNumber() *big.Int {
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(1024))
	CheckError(err)
	return randomNumber
}

func createCertificate() ([]byte, *rsa.PrivateKey, *rsa.PublicKey) {
	GetRootCA()
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
	certificateBytes, err := x509.CreateCertificate(rand.Reader, certificate, certificateAuthority, publicKey, privateKey)
	CheckError(err)
	return certificateBytes, privateKey, publicKey
}

func savePublicKey(certificateBytes []byte) {
	CheckError(os.MkdirAll(cryptographyDirectory, 0700))
	certificateFile, err := os.Create(certFileName)
	CheckError(err)
	CheckError(pem.Encode(certificateFile, &pem.Block{Type: "CERTIFICATE", Bytes: certificateBytes}))
	CheckError(certificateFile.Close())
	beego.Info("Created new cert file " + certFileName)
}

func savePrivateKey(privateKey *rsa.PrivateKey) {
	CheckError(os.MkdirAll(cryptographyDirectory, 0700))
	keyFile, err := os.OpenFile(keyFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	CheckError(err)
	CheckError(pem.Encode(keyFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}))
	CheckError(keyFile.Close())
	beego.Info("Created new key file " + keyFileName)
}
