package easytls

import (
	"crypto/tls"
)

type fileClient struct {
	certFile string
	keyFile  string
}

func (l *fileClient) GetCertificate() func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
	return GetCertificate(l.GetTls)
}

func (l *fileClient) GetTls() (tls.Certificate, error) {
	if l.certFile == "" || l.keyFile == "" {
		panic("cert file or key file is empty")
	}
	return tls.LoadX509KeyPair(l.certFile, l.keyFile)
}
