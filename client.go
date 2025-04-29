package easytls

import (
	"crypto/tls"
	"easytls-sdk-go/esaylego"
	"log"
	"sync"
	"time"

	"github.com/go-acme/lego/v4/lego"
)

const DNS_CHALLENGE = "DNS_CHALLENGE"

type TlsClientConfig struct {
	LegoConfig *lego.Config
	Domains    []string
	DnsConfig  map[string]string

	CertFile string
	KeyFile  string
}

type Client interface {
	GetTls() (tls.Certificate, error)
	GetCertificate() func(*tls.ClientHelloInfo) (*tls.Certificate, error)
}

func getTheUpdateTime(cert tls.Certificate) time.Time {
	if cert.Leaf == nil {
		return time.Now()
	}
	return cert.Leaf.NotAfter.Add(-7 * 24 * time.Hour)
}

func NewLegoClientByDNS(config TlsClientConfig) (*legoClient, error) {
	esaylegoClient, err := esaylego.NewEsayClient(config.LegoConfig)
	return &legoClient{
		client:    esaylegoClient,
		domains:   config.Domains,
		dnsConfig: config.DnsConfig,
	}, err
}

func NewFileClientByFile(config TlsClientConfig) (*fileClient, error) {
	return &fileClient{
		certFile: config.CertFile,
		keyFile:  config.KeyFile,
	}, nil
}

func GetCertificate(getTls func() (tls.Certificate, error)) func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
	certificate := tls.Certificate{}
	mu := sync.RWMutex{}
	go func() {
		for {
			newCertificate, err := getTls()
			if err != nil {
				log.Fatalln(err)
				continue
			}
			mu.Lock()
			certificate = newCertificate
			mu.Unlock()
			updateTime := getTheUpdateTime(newCertificate)
			sleepDuration := time.Until(updateTime)
			if sleepDuration < 0 {
				sleepDuration = time.Minute // 默认间隔
			}
			time.Sleep(sleepDuration)
		}
	}()
	return func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
		mu.RLock()
		defer mu.RUnlock()
		return &certificate, nil
	}
}
