package easytls

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"easytls-sdk-go/esaylego"

	"github.com/go-acme/lego/v4/lego"
)

const DNS_CHALLENGE = "DNS_CHALLENGE"

type TlsClientConfig struct {
	/*
		必填

		Required
	*/
	Domains []string
	/*
		NewLegoClientByDNS 时使用
		LegoConfig.CADirURL 默认为 "lego.LEDirectoryProduction" 调试时请建议切换 "lego.LEDirectoryStaging"

		NewLegoClientByDNS
		LegoConfig.CADirURL defaults to "lego. LEDirectoryProduction" Please suggest switching to "lego. LEDirectoryStaging"
	*/
	LegoConfig *lego.Config
	/*
		NewLegoClientByDNS 时使用
		dns 运营商 alydns 配置 案例：
			DNS_CHALLENGE: alidns
			ALICLOUD_ACCESS_KEY:
			ALICLOUD_SECRET_KEY:
			ALICLOUD_REGION_ID:

			更多配置请参考 https://go-acme.github.io/lego/dns/active24/

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
		NewFileClientByFile 时使用
	*/
	CertFile string
	/*
		NewFileClientByFile 时使用
	*/
	KeyFile string

	/*
		NewLegoServerClientByFile 时使用
			TimeDiff    time.Duration  启用gozero 签名时使用
			Fingerprint string         启用gozero 签名时使用
			Key         []byte         启用gozero 签名时使用
			PubKey      []byte         启用gozero 签名时使用
	*/
	Host        string
	Scheme      string
	TimeDiff    time.Duration // 启用gozero 签名时使用
	Fingerprint string        // 启用gozero 签名时使用
	Key         []byte        // 启用gozero 签名时使用
	PubKey      []byte        // 启用gozero 签名时使用
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

func NewLegoServerClientByFile(config TlsClientConfig) (*legoServerClient, error) {
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
