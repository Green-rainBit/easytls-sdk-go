package easytls

import (
	"crypto/tls"
	"easytls-sdk-go/esaylego"
	"easytls-sdk-go/esaylego/dns"
	"fmt"

	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/challenge/resolver"
)

type legoClient struct {
	client    *esaylego.EsayClient
	domains   []string
	dnsConfig map[string]string
}

func (l *legoClient) GetCertificate() func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
	return GetCertificate(l.GetTls)
}

func (l *legoClient) GetTls() (tls.Certificate, error) {
	provider, err := dns.NewDNSChallengeProviderByName(l.dnsConfig[DNS_CHALLENGE], l.dnsConfig)
	if err != nil {
		return tls.Certificate{}, err
	}
	solversManager := resolver.NewSolversManager(l.client.GetCore())
	solversManager.SetDNS01Provider(provider)
	prober := resolver.NewProber(solversManager)
	certifier := certificate.NewCertifier(l.client.GetCore(), prober, l.client.GetCertifierOptionse())
	if len(l.domains) == 0 {
		return tls.Certificate{}, fmt.Errorf("domains not is empty")
	}
	request := certificate.ObtainRequest{
		Domains: l.domains,
		Bundle:  true,
	}
	certificates, err := certifier.Obtain(request)
	if err != nil {
		return tls.Certificate{}, err
	}
	if certificates != nil {
		newCertificate, err := tls.X509KeyPair(certificates.Certificate, certificates.PrivateKey)
		if err != nil {
			return tls.Certificate{}, err
		}
		return newCertificate, nil
	}
	return tls.Certificate{}, fmt.Errorf("certificates not is empty")
}
