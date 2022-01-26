package onlinesim

import (
	"encoding/json"
	"fmt"
	"github.com/ddliu/go-httpclient"
	"net/http"
	"time"
)

const (
	baseURL   = "https://onlinesim.ru/api"
	rateLimit = 2
)

type Onlinesim struct {
	rateLimiter <-chan time.Time
	baseURL     string
	apiKey      string
	lang        string
	dev_id      string
	proxy       string
}

type Default struct {
	Response interface{} `json:"response"`
}

func NewClient(apiKey string, lang string, dev_id string, proxy string) *Onlinesim {
	if lang == "" {
		lang = "en"
	}

	return &Onlinesim{
		rateLimiter: time.Tick(time.Second / time.Duration(rateLimit)),
		apiKey:      apiKey,
		baseURL:     baseURL,
		lang:        lang,
		dev_id:      dev_id,
		proxy:       proxy,
	}
}

// SetRateLimit rate limit setter for custom usage
// Onlinesim limit is 5 requests per second (we use 2)
func (at *Onlinesim) SetRateLimit(customRateLimit int) {
	at.rateLimiter = time.Tick(time.Second / time.Duration(customRateLimit))
}

func (at *Onlinesim) rateLimit() {
	<-at.rateLimiter
}

func (at *Onlinesim) get(method string, params map[string]string) []byte {
	at.rateLimit()
	params["apikey"] = at.apiKey
	params["lang"] = at.lang
	params["dev_id"] = at.dev_id

	url := fmt.Sprintf("%s/%s.php", at.baseURL, method)

	var client *httpclient.HttpClient
	if at.proxy != "" {
		client = httpclient.WithOption(httpclient.OPT_PROXY_FUNC, func(*http.Request) (int, string, error) {
			return httpclient.PROXY_HTTP, at.proxy, nil
		}).Defaults(httpclient.Map{
			httpclient.OPT_USERAGENT:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36",
			"Accept-Language":         "en-us",
			httpclient.OPT_UNSAFE_TLS: true,
		})
	} else {
		client = httpclient.Defaults(httpclient.Map{
			httpclient.OPT_USERAGENT:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36",
			"Accept-Language":         "en-us",
			httpclient.OPT_UNSAFE_TLS: true,
		})
	}

	res, err := client.Get(url, params)
	if err != nil {
		panic(fmt.Errorf("request error: %w", err))
	}
	bodyString, err := res.ToString()
	if err != nil {
		panic(err)
	}

	return []byte(bodyString)
}

func (at *Onlinesim) getWithErr(method string, params map[string]string) ([]byte, error) {
	at.rateLimit()
	params["apikey"] = at.apiKey
	params["lang"] = at.lang
	params["dev_id"] = at.dev_id

	url := fmt.Sprintf("%s/%s.php", at.baseURL, method)

	var client *httpclient.HttpClient
	if at.proxy != "" {
		client = httpclient.WithOption(httpclient.OPT_PROXY_FUNC, func(*http.Request) (int, string, error) {
			return httpclient.PROXY_HTTP, at.proxy, nil
		}).Defaults(httpclient.Map{
			httpclient.OPT_USERAGENT:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36",
			"Accept-Language":         "en-us",
			httpclient.OPT_UNSAFE_TLS: true,
		})
	} else {
		client = httpclient.Defaults(httpclient.Map{
			httpclient.OPT_USERAGENT:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36",
			"Accept-Language":         "en-us",
			httpclient.OPT_UNSAFE_TLS: true,
		})
	}

	res, err := client.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("request error: %w", err)
	}
	bodyString, err := res.ToString()
	if err != nil {
		return nil, err
	}

	return []byte(bodyString), nil
}

func (at *Onlinesim) checkResponse(resp interface{}) error {
	if fmt.Sprintf("%v", resp) != "1" {
		return fmt.Errorf("%s", resp)
	}
	return nil
}

func (at *Onlinesim) checkEmptyResponse(resp []byte) error {
	__default := Default{}
	err := json.Unmarshal(resp, &__default)
	if err == nil {
		if __default.Response == nil {
			return nil
		}
		if fmt.Sprintf("%v", __default.Response) == "" {
			return nil
		}
		if fmt.Sprintf("%v", __default.Response) != "1" {
			return fmt.Errorf("%s", __default.Response)
		}
	}
	return nil
}

func (c *Onlinesim) Free() *GetFree {
	return &GetFree{
		client: c,
	}
}

func (c *Onlinesim) Numbers() *GetNumbers {
	return &GetNumbers{
		client: c,
	}
}

func (c *Onlinesim) Proxy() *GetProxy {
	return &GetProxy{
		client: c,
	}
}

func (c *Onlinesim) Rent() *GetRent {
	return &GetRent{
		client: c,
	}
}

func (c *Onlinesim) User() *GetUser {
	return &GetUser{
		client: c,
	}
}
