package config

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

func InitElasticClient() (*elasticsearch.Client, error) {
	client := http.Transport{
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * 2,
		DialContext:           (&net.Dialer{Timeout: time.Second * 2}).DialContext,
		TLSClientConfig: &tls.Config{
			MaxVersion:         tls.VersionTLS11,
			InsecureSkipVerify: true,
		}}

	return elasticsearch.NewClient(elasticsearch.Config{

		Addresses: []string{"http://localhost:9200"},
		Username:  "",
		Password:  "",
		Transport: &client,
	})
}
