package provider

import (
	"io"
	"strings"
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
	serverName := strings.Split(path, "/")[1]
	serverPath := strings.Replace(path, "/" + serverName, "", 1)

	return HttpProxy {method,config.HttpProxyConfig[serverName] + serverPath, body, headers }
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
