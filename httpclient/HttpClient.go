package httpclient

import (
	"io/ioutil"
	"fmt"
	"strings"
	"net/url"
	"sync"
	"crypto/tls"
	"net/http"
)

type HTTPClient struct {
	mutex        sync.Mutex
	client       *http.Client
	transport    *http.Transport
	schemePrefix string
}

func NewHTTPClient() *HTTPClient {
	var instance = new(HTTPClient)
	instance.schemePrefix = "http://"
	instance.transport = &http.Transport{
		MaxIdleConnsPerHost: 1024,
		TLSClientConfig:     nil,
	}
	instance.client = &http.Client{Transport: instance.transport}
	return instance
}

func NewHTTPsClient(tlsConfig *tls.Config) *HTTPClient {
	var instance = new(HTTPClient)
	instance.schemePrefix = "https://"
	instance.transport = &http.Transport{
		MaxIdleConnsPerHost: 1024,
		TLSClientConfig:     tlsConfig,
	}
	instance.client = &http.Client{Transport: instance.transport}
	return instance
}

func (this HTTPClient) Post(url string, values url.Values) ([]byte, error) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	r, err := this.client.PostForm(url, values)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (this HTTPClient) Get(url string) ([]byte, error) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	r, err := this.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if r.StatusCode < 200 || r.StatusCode >= 300 {
		return nil, fmt.Errorf("%s: %s", url, r.Status)
	}
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (this HTTPClient) DownloadUrl(fileUrl string) (filename string, content []byte, e error) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	response, err := this.client.Get(fileUrl)
	if err != nil {
		return "", nil, err
	}
	defer response.Body.Close()
	contentDisposition := response.Header["Content-Disposition"]
	if len(contentDisposition) > 0 {
		if strings.HasPrefix(contentDisposition[0], "filename=") {
			filename = contentDisposition[0][len("filename="):]
			filename = strings.Trim(filename, "\"")
		}
	}
	content, e = ioutil.ReadAll(response.Body)
	return
}
