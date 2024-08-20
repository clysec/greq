package greq

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"os"

	"software.sslmate.com/src/go-pkcs12"
)

type ClientCertificateAuth struct {
	ClientCertificate  tls.Certificate
	CaCertificates     *x509.CertPool
	InsecureSkipVerify bool
}

func (ca ClientCertificateAuth) Prepare() error {
	return nil
}

func (ca ClientCertificateAuth) Apply(addHeaderFunc func(key, value string), setTransportFunc func(transport http.RoundTripper)) error {
	customTransport := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates:       []tls.Certificate{ca.ClientCertificate},
			InsecureSkipVerify: ca.InsecureSkipVerify,
		},
	}

	if ca.CaCertificates != nil {
		customTransport.TLSClientConfig.RootCAs = ca.CaCertificates
	}

	setTransportFunc(customTransport)
	return nil
}

func NewClientCertificateAuth() *ClientCertificateAuth {
	return &ClientCertificateAuth{}
}

func (ca *ClientCertificateAuth) FromX509(certFile, keyFile string) *ClientCertificateAuth {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		panic(err)
	}

	ca.ClientCertificate = cert
	return ca
}

func (ca *ClientCertificateAuth) FromX509Bytes(cert, key []byte) *ClientCertificateAuth {
	certPair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		panic(err)
	}

	ca.ClientCertificate = certPair

	return ca
}

func (ca *ClientCertificateAuth) WithCaCertificates(caCertificates *x509.CertPool) *ClientCertificateAuth {
	ca.CaCertificates = caCertificates
	return ca
}

func (ca *ClientCertificateAuth) WithInsecureSkipVerify(insecureSkipVerify bool) *ClientCertificateAuth {
	ca.InsecureSkipVerify = insecureSkipVerify
	return ca
}

func (ca *ClientCertificateAuth) FromPKCS12(pkcs12File, password string) *ClientCertificateAuth {
	contents, err := os.ReadFile(pkcs12File)
	if err != nil {
		panic(err)
	}

	ca.FromPKCS12Bytes(contents, password)

	return ca
}

func (ca *ClientCertificateAuth) FromPKCS12Bytes(pkcs12Data []byte, password string) *ClientCertificateAuth {
	privkey, certificate, cachain, err := pkcs12.DecodeChain(pkcs12Data, password)
	if err != nil {
		panic(err)
	}

	ca.ClientCertificate = tls.Certificate{
		Certificate: [][]byte{certificate.Raw},
		PrivateKey:  privkey,
	}

	ca.CaCertificates = x509.NewCertPool()
	for _, cert := range cachain {
		ca.CaCertificates.AddCert(cert)
	}

	return ca
}
