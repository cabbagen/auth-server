package provider

import (
	"io"
	"net/http"
	"io/ioutil"
	"go-gateway/config"
)

type HttpProxy struct {
	Method              string
	RequestUrl          string
	Body                io.Reader
	Headers             map[string]string
}

type ProxyHeader struct {
	ContentType     string `header:"Content-Type"`
}

func NewHttpProxy(method, path string, body io.Reader, headers map[string]string) HttpProxy {
	return HttpProxy {method,config.HttpProxyConfig["http"] + path, body, headers }
}

func GetDefaultHttpClient() *http.Client {
	return &http.Client{}
}

func (hp HttpProxy) Request(headers ProxyHeader) ([]byte, error) {
	request, error := http.NewRequest(hp.Method, hp.RequestUrl, hp.Body)

	if error != nil {
		return nil, error
	}

	request.Close = true

	request.Header.Set("content-type", headers.ContentType)

	for key, value := range hp.Headers {
		request.Header.Set(key, value)
	}

	response, error := GetDefaultHttpClient().Do(request)

	if error != nil {
		return nil, error
	}

	return ioutil.ReadAll(response.Body)
}
