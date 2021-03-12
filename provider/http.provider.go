package provider

import (
	"io"
	"net/http"
	"io/ioutil"
	"auth-go/config"
)

type HttpProxy struct {
	Method              string
	RequestUrl          string
	Body                io.Reader
}

type ProxyHeader struct {
	ContentType     string `header:"Content-Type"`
}

func NewHttpProxy(method, path string, body io.Reader) HttpProxy {
	return HttpProxy {method,config.HttpProxyConfig["http"] + path, body }
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

	response, error := GetDefaultHttpClient().Do(request)

	// defer response.Body.Close()

	if error != nil {
		return nil, error
	}

	return ioutil.ReadAll(response.Body)
}
