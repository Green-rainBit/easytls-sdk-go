package esaylego

import (
	"crypto"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"log"
	"net/url"

	"github.com/go-acme/lego/v4/acme/api"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
)

type EsaylegoUser struct {
	Email        string
	Key          crypto.PrivateKey
	Registration *registration.Resource
}

func (u *EsaylegoUser) GetEmail() string {
	return u.Email
}
func (u *EsaylegoUser) GetRegistration() *registration.Resource {
	return u.Registration
}
func (u *EsaylegoUser) GetPrivateKey() crypto.PrivateKey {
	return u.Key
}

func NewLegoConfigUser(email, privateKeyStrey string) *EsaylegoUser {
	bytes, err := base64.StdEncoding.DecodeString(privateKeyStrey)
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := x509.ParseECPrivateKey(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return &EsaylegoUser{
		Email:        email,
		Key:          privateKey,
		Registration: &registration.Resource{},
	}
}

// Client is the user-friendly way to ACME
type EsayClient struct {
	core             *api.Core
	Registration     *registration.Registrar
	certifierOptions certificate.CertifierOptions
}

func NewEsayClient(config *lego.Config) (*EsayClient, error) {
	if config == nil {
		return nil, errors.New("a configuration must be provided")
	}

	_, err := url.Parse(config.CADirURL)
	if err != nil {
		return nil, err
	}

	if config.HTTPClient == nil {
		return nil, errors.New("the HTTP client cannot be nil")
	}

	privateKey := config.User.GetPrivateKey()
	if privateKey == nil {
		return nil, errors.New("private key was nil")
	}

	var kid string
	if reg := config.User.GetRegistration(); reg != nil {
		kid = reg.URI
	}

	core, err := api.New(config.HTTPClient, config.UserAgent, config.CADirURL, kid, privateKey)
	if err != nil {
		return nil, err
	}

	// prober := resolver.NewProber(solversManager)
	// certifier := certificate.NewCertifier(core, prober, certificate.CertifierOptions{KeyType: config.Certificate.KeyType, Timeout: config.Certificate.Timeout, OverallRequestLimit: config.Certificate.OverallRequestLimit})

	return &EsayClient{
		certifierOptions: certificate.CertifierOptions{KeyType: config.Certificate.KeyType, Timeout: config.Certificate.Timeout},
		Registration:     registration.NewRegistrar(core, config.User),
		core:             core,
	}, nil
}

func (e *EsayClient) GetCore() *api.Core {
	return e.core
}

func (e *EsayClient) GetCertifierOptionse() certificate.CertifierOptions {
	return e.certifierOptions
}
