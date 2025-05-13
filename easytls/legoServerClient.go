package easytls

import (
	"bytes"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/codec"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type legoServerClient struct {
	u           url.URL
	timeDiff    time.Duration
	fingerprint string
	client      http.Client
	key         []byte
	pubKey      []byte
}

func (l *legoServerClient) GetCertificate() func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
	return GetCertificate(l.GetTls)
}

func (l *legoServerClient) GetTls() (tls.Certificate, error) {
	rs := requestSettings{
		method:      http.MethodGet,
		url:         l.u.String(),
		fingerprint: l.fingerprint,
		timestamp:   time.Now().Add(l.timeDiff).Unix(),
	}
	req, err := l.buildRequest(rs)
	if err != nil {
		return tls.Certificate{}, err
	}
	resp, err := l.client.Do(req)
	if err != nil {
		return tls.Certificate{}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return tls.Certificate{}, err
		}
		certificates := struct {
			Domain            string `json:"domain"`
			CertURL           string `json:"certUrl"`
			CertStableURL     string `json:"certStableUrl"`
			PrivateKey        []byte `json:"privateKey"`
			Certificate       []byte `json:"certificate"`
			IssuerCertificate []byte `json:"issuerCertificate"`
			CSR               []byte `json:"cSR"`
		}{}
		if err = json.Unmarshal(body, &certificates); err != nil {
			return tls.Certificate{}, err
		}
		newCertificate, err := tls.X509KeyPair(certificates.Certificate, certificates.PrivateKey)
		if err != nil {
			return tls.Certificate{}, err
		}
		return newCertificate, nil
	}

	return tls.Certificate{}, err
}

type requestSettings struct {
	method      string
	url         string
	body        io.Reader
	crypt       bool
	requestUri  string
	timestamp   int64
	fingerprint string
	missHeader  bool
	signature   string
}

func (l *legoServerClient) buildRequest(rs requestSettings) (*http.Request, error) {
	var bodyStr string
	var err error

	if rs.crypt && rs.body != nil {
		var buf bytes.Buffer
		io.Copy(&buf, rs.body)
		bodyBytes, err := codec.EcbEncrypt(l.key, buf.Bytes())
		if err != nil {
			return nil, err
		}
		bodyStr = base64.StdEncoding.EncodeToString(bodyBytes)
	}

	r, err := http.NewRequest(rs.method, rs.url, strings.NewReader(bodyStr))
	if err != nil {
		return nil, err
	}
	if l.pubKey == nil {
		return r, nil
	}
	if len(rs.signature) == 0 {
		sha := sha256.New()
		sha.Write([]byte(bodyStr))
		bodySign := fmt.Sprintf("%x", sha.Sum(nil))
		var path string
		var query string
		if len(rs.requestUri) > 0 {
			u, err := url.Parse(rs.requestUri)
			if err != nil {
				return nil, err
			}

			path = u.Path
			query = u.RawQuery
		} else {
			path = r.URL.Path
			query = r.URL.RawQuery
		}
		contentOfSign := strings.Join([]string{
			strconv.FormatInt(rs.timestamp, 10),
			rs.method,
			path,
			query,
			bodySign,
		}, "\n")
		rs.signature = codec.HmacBase64(l.key, contentOfSign)
	}

	var mode string
	if rs.crypt {
		mode = "1"
	} else {
		mode = "0"
	}
	content := strings.Join([]string{
		"version=v1",
		"type=" + mode,
		fmt.Sprintf("key=%s", base64.StdEncoding.EncodeToString(l.key)),
		"time=" + strconv.FormatInt(rs.timestamp, 10),
	}, "; ")

	encrypter, err := codec.NewRsaEncrypter(l.pubKey)
	if err != nil {
		return nil, err
	}

	output, err := encrypter.Encrypt([]byte(content))
	if err != nil {
		return nil, err
	}

	encryptedContent := base64.StdEncoding.EncodeToString(output)
	if !rs.missHeader {
		r.Header.Set(httpx.ContentSecurity, strings.Join([]string{
			fmt.Sprintf("key=%s", rs.fingerprint),
			"secret=" + encryptedContent,
			"signature=" + rs.signature,
		}, "; "))
	}
	if len(rs.requestUri) > 0 {
		r.Header.Set("X-Request-Uri", rs.requestUri)
	}

	return r, nil
}
