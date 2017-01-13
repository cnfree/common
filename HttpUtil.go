package common

import (
	"sync"
	"github.com/cnfree/common/httpclient"
	"crypto/tls"
)

type httpUtil struct {
	mutex sync.Mutex
}

var HTTP = httpUtil{}

func (this httpUtil) NewHTTPClient() *httpclient.HTTPClient {
	return httpclient.NewHTTPClient()
}

func (this httpUtil) NewHTTPsClient(tlsConfig *tls.Config) *httpclient.HTTPClient {
	return httpclient.NewHTTPsClient(tlsConfig)
}
