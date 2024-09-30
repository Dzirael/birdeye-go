package birdeye

import (
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

var (
	baseURL = "https://public-api.birdeye.so"
)

type birdeye struct {
	client *resty.Client
}

func New(apiKey string, chain ...chain) Birdeye {
	client := resty.New().
		SetBaseURL(baseURL).
		SetHeader("X-API-KEY", apiKey)

	if len(chain) > 0 {
		client.SetHeader("x-chain", string(chain[0]))
	}

	return &birdeye{
		client: client,
	}
}

func (b *birdeye) SetXChain(xChain chain) *birdeye {
	b.client.SetHeader("x-chain", string(xChain))
	return b
}

func (b *birdeye) SetAPIKey(apiKey string) *birdeye {
	b.client.SetHeader("X-API-KEY", apiKey)
	return b
}

func (b *birdeye) call(req *resty.Request, method string, path string) (err error) {
	var resp *resty.Response
	switch method {
	case http.MethodGet:
		resp, err = req.Get(path)
	case http.MethodPost:
		resp, err = req.Post(path)
	}

	if err != nil {
		return errors.Wrap(err, "Birdeye: failed to make request")
	}

	if resp.IsError() {
		return errors.Errorf("Birdeye: request failed with status code: %d and body: %s", resp.StatusCode(), resp.String())
	}

	return nil
}
