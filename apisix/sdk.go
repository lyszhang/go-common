package apisix_sdk

import (
	"errors"
	"fmt"
	"sync"

	"github.com/twwch/gin-sdk/httpclient/base"
)

var once sync.Once
var __client *apisixClient
var (
	URIORURLSChooseOneError = errors.New("URI和URIS不能同时使用")
)

func newErrors(msg string) error {
	return errors.New(msg)
}

func GetApiSixClient() *apisixClient {
	return __client
}

type apisixClient struct {
	client base.HttpClient
}

func (apisix *apisixClient) GetRoute() *Route {
	return &Route{apisixClient{client: apisix.client}}
}

func (apisix *apisixClient) GetUpstream() *Upstream {
	return &Upstream{apisixClient{client: apisix.client}}
}

func NewApiSixClient(host, key string) {
	once.Do(func() {
		httpClient, err := base.NewClient(base.SetHost(host), base.SetHeaders(map[string]string{"X-API-KEY": key}))
		if err != nil {
			panic(fmt.Sprintf("create apisix client failed, err: %s", err.Error()))
		}
		__client = &apisixClient{
			client: httpClient,
		}
	})
}
