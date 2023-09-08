package aerospike

import (
	"errors"

	as "github.com/aerospike/aerospike-client-go"
	"github.com/en666ki/urlime/internal/db"
)

type AerospikeDB struct {
	host      string
	port      int
	namespace string
	set       string
}

func New(host string, port int, namespace, set string) *AerospikeDB {
	return &AerospikeDB{host: host, port: port, namespace: namespace, set: set}
}

func (a *AerospikeDB) PutUrl(url *db.Url) error {
	c, err := as.NewClient(a.host, a.port)
	if err != nil {
		return err
	}
	if c == nil {
		return errors.New("client is nil")
	}
	defer c.Close()

	key, err := as.NewKey(a.namespace, a.set, url.Surl)
	if err != nil {
		return err
	}
	if key == nil {
		return errors.New("nil key")
	}

	policy := as.NewWritePolicy(0, 0)

	err = c.Put(policy, key, as.BinMap{"url": url.Url})
	if err != nil {
		return err
	}

	return nil
}

func (a *AerospikeDB) GetUrl(surl string) (*db.Url, error) {
	c, err := as.NewClient(a.host, a.port)
	if err != nil {
		return nil, err
	}
	if c == nil {
		return nil, errors.New("client is nil")
	}
	defer c.Close()

	key, err := as.NewKey(a.namespace, a.set, surl)
	if err != nil {
		return nil, err
	}
	if key == nil {
		return nil, errors.New("nil key")
	}

	policy := as.NewPolicy()

	bins, err := c.Get(policy, key)
	if err != nil {
		return nil, err
	}

	url, ok := bins.Bins["url"]
	if !ok {
		return nil, errors.New("no url bin in record")
	}

	fUrl, ok := url.(string)
	if !ok {
		return nil, errors.New("stored url is not string")
	}

	return db.NewUrl(surl, fUrl, a), nil
}
