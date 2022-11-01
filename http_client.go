package openinsider

import (
	"net/http"

	gocache "github.com/mrod502/go-cache"
	"github.com/mrod502/htmlmarshal"
)

func NewClient(opts Options) (Client, error) {

	return &HttpClient{
		opts: opts,
		data: gocache.New[string, interface{}]().WithExpiration(opts.Ttl),
		cli:  http.DefaultClient,
	}, nil
}

type HttpClient struct {
	opts Options
	data *gocache.Cache[string, interface{}]
	ds   htmlmarshal.Deserializer
	cli  *http.Client
}

func (c *HttpClient) LatestClusterBuys() ([]Disclosure, error) {
	b, err := c.get("/latest-cluster-buys")
	if err != nil {
		return nil, err
	}
	var buys = make([]Disclosure, 0)
	err = c.ds.Deserialize(b, &buys)
	return buys, err
}

func (c *HttpClient) Screen(f *Filter) ([]Disclosure, error) {

	return nil, nil
}

func (c HttpClient) get(path string) ([]byte, error) {
	return nil, nil
}
