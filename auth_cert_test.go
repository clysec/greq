package greq_test

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/clysec/greq"
)

type CaCertCollection struct {
	CACert  x509.Certificate
	CAKey   *rsa.PrivateKey
	CABytes []byte

	CaPEM        []byte
	CAPrivkeyPEM []byte
}

type CertCollection struct {
	Cert           tls.Certificate
	CertPEM        []byte
	CertPrivkeyPEM []byte
}

func (cac *CaCertCollection) GetSignedCert(cert *x509.Certificate) CertCollection {
	certPrivkey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, cert, &cac.CACert, &certPrivkey.PublicKey, cac.CAKey)
	if err != nil {
		panic(err)
	}

	certPEM := new(bytes.Buffer)
	err = pem.Encode(certPEM, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
	if err != nil {
		panic(err)
	}

	certPrivkeyPEM := new(bytes.Buffer)
	pem.Encode(certPrivkeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(certPrivkey),
	})

	crt, err := tls.X509KeyPair(certPEM.Bytes(), certPrivkeyPEM.Bytes())
	if err != nil {
		panic(err)
	}

	return CertCollection{
		Cert:           crt,
		CertPEM:        certPEM.Bytes(),
		CertPrivkeyPEM: certPrivkeyPEM.Bytes(),
	}
}

func (cac *CaCertCollection) CreateClientCert() CertCollection {
	return cac.GetSignedCert(&x509.Certificate{
		SerialNumber: big.NewInt(2025),
		Subject: pkix.Name{
			Country:       []string{"SE"},
			Organization:  []string{"Cloudyne Systems"},
			Province:      []string{"Stockholm"},
			Locality:      []string{"Stockholm"},
			StreetAddress: []string{"Kungsgatan 1"},
			PostalCode:    []string{"111 11"},
			CommonName:    "localhost",
		},
		DNSNames:     []string{"localhost"},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(0, 0, 10),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		KeyUsage:     x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
	})
}

func (cac *CaCertCollection) CreateServerCert() CertCollection {
	return cac.GetSignedCert(&x509.Certificate{
		SerialNumber: big.NewInt(2025),
		Subject: pkix.Name{
			Country:       []string{"SE"},
			Organization:  []string{"Cloudyne Systems"},
			Province:      []string{"Stockholm"},
			Locality:      []string{"Stockholm"},
			StreetAddress: []string{"Kungsgatan 1"},
			PostalCode:    []string{"111 11"},
			CommonName:    "localhost",
		},
		DNSNames:     []string{"localhost"},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(0, 0, 10),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		KeyUsage:     x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
	})
}

func PrepareCA() (certCollection CaCertCollection) {
	var err error

	certCollection.CACert = x509.Certificate{
		SerialNumber: big.NewInt(2024),
		Subject: pkix.Name{
			Country:       []string{"SE"},
			Organization:  []string{"Cloudyne Systems"},
			Province:      []string{"Stockholm"},
			Locality:      []string{"Stockholm"},
			StreetAddress: []string{"Kungsgatan 1"},
			PostalCode:    []string{"111 11"},
			CommonName:    "Cloudyne CA",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	certCollection.CAKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	caBytes, err := x509.CreateCertificate(rand.Reader, &certCollection.CACert, &certCollection.CACert, &certCollection.CAKey.PublicKey, certCollection.CAKey)
	if err != nil {
		panic(err)
	}

	certCollection.CABytes = caBytes

	caPEM := new(bytes.Buffer)

	err = pem.Encode(caPEM, &pem.Block{Type: "CERTIFICATE", Bytes: caBytes})
	if err != nil {
		panic(err)
	}

	caPrivkeyPEM := new(bytes.Buffer)
	pem.Encode(caPrivkeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(certCollection.CAKey),
	})
	certCollection.CaPEM = caPEM.Bytes()
	certCollection.CAPrivkeyPEM = caPrivkeyPEM.Bytes()

	return
}

func TestCertificateAuth(t *testing.T) {
	ca := x509.NewCertPool()

	caData, err := os.ReadFile("testfiles/rootca.crt")
	if err != nil {
		t.Fatal(err)
	}

	ca.AppendCertsFromPEM(caData)

	server := &http.Server{
		Addr: ":55811",
		TLSConfig: &tls.Config{
			ClientCAs:        ca,
			ClientAuth:       tls.RequireAndVerifyClientCert,
			MinVersion:       tls.VersionTLS12,
			CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			},
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	go func() {
		err := server.ListenAndServeTLS("testfiles/server.crt", "testfiles/server.key")
		if err != nil {
			t.Fatal(err)
		}

	}()

	time.Sleep(2 * time.Second)

	auth := greq.NewClientCertificateAuth().FromX509("testfiles/client.crt", "testfiles/client.key").WithCaCertificates(ca)

	resp, err := greq.GetRequest("https://localhost:55811").WithAuth(auth).Execute()
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Invalid response: not 200 but %d", resp.StatusCode)
	}

}
