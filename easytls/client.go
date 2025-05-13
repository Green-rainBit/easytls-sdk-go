package easytls

import (
	"crypto/tls"
	"errors"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/Green-rainBit/easytls-sdk-go/esaylego"
	"github.com/go-acme/lego/v4/lego"
)

const DNS_CHALLENGE = "DNS_CHALLENGE"

type TlsClientConfig struct {
	/*
		NewLegoClientByDNS
		production or develop (default is develop)
	*/
	Env string
	/*
		NewLegoClientByDNS or NewLegoServerClientByServer Required
	*/
	Domains []string

	LegoConfigUser *esaylego.EsaylegoUser
	/*
		NewLegoClientByDNS
		LegoConfig.CADirURL defaults to "lego.LEDirectoryProduction" Please suggest switching to "lego. LEDirectoryStaging"
	*/
	LegoConfig *lego.Config
	/*
		NewLegoClientByDNS
		NewLegoClientByDNS
			DNS Carrier alydns Configuration Examples:
			DNS_CHALLENGE: alidns
			ALICLOUD_ACCESS_KEY:
			ALICLOUD_SECRET_KEY:
			ALICLOUD_REGION_ID:

			For more configurations, please refer to https://go-acme.github.io/lego/dns/active24/
	*/
	DnsConfig map[string]string

	/*
		NewFileClientByFile
	*/
	CertFile string
	/*
		NewFileClientByFile
	*/
	KeyFile string

	/*
		NewLegoServerClientByServer
			TimeDiff    time.Duration  user gozero Signature
			Fingerprint string         user gozero Signature
			Key         []byte         user gozero Signature
			PubKey      []byte         user gozero Signature
	*/
	Host        string
	Scheme      string
	TimeDiff    time.Duration
	Fingerprint string
	Key         []byte
	PubKey      []byte
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

func NewLegoClient(config TlsClientConfig) (Client, error) {
	if config.KeyFile != "" && config.CertFile != "" {
		return NewFileClientByFile(config)
	} else if config.LegoConfig != nil || config.LegoConfigUser != nil {
		return NewLegoClientByDNS(config)
	} else if config.Host != "" && config.Scheme != "" {
		return NewLegoServerClientByServer(config)
	}
	return nil, errors.New("config error")
}

func NewLegoClientByDNS(config TlsClientConfig) (*legoClient, error) {
	if config.LegoConfigUser == nil {
		config.LegoConfigUser = &esaylego.EsaylegoUser{
			Email: config.LegoConfig.User.GetEmail(),
			Key:   config.LegoConfig.User.GetPrivateKey(),
		}
	}
	if config.LegoConfig == nil {
		config.LegoConfig = lego.NewConfig(config.LegoConfigUser)
	}
	if config.Env != "production" {
		config.LegoConfig.CADirURL = lego.LEDirectoryStaging
	}
	esaylegoClient, err := esaylego.NewEsayClient(config.LegoConfig)
	if err != nil {
		return nil, err
	}
	req, err := esaylegoClient.Registration.ResolveAccountByKey()
	if err != nil {
		return nil, err
	}
	config.LegoConfigUser.Registration = req
	return &legoClient{
		client:    esaylegoClient,
		domains:   config.Domains,
		dnsConfig: config.DnsConfig,
	}, nil
}

func NewFileClientByFile(config TlsClientConfig) (*fileClient, error) {
	return &fileClient{
		certFile: config.CertFile,
		keyFile:  config.KeyFile,
	}, nil
}

func NewLegoServerClientByServer(config TlsClientConfig) (*legoServerClient, error) {
	values := url.Values{}
	for _, domain := range config.Domains {
		values.Add("domains", domain)
	}
	u := url.URL{
		Path:     "/v1/get/tls",
		Host:     config.Host,
		Scheme:   config.Scheme,
		RawQuery: values.Encode(),
	}
	return &legoServerClient{
		fingerprint: config.Fingerprint,
		key:         config.Key,
		pubKey:      config.PubKey,
		timeDiff:    config.TimeDiff,
		u:           u,
		client:      http.Client{},
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
