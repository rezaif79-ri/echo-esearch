package config

import "github.com/elastic/go-elasticsearch/v8"

func InitElasticClient() (*elasticsearch.Client, error) {
	return elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
		Username:  "",
		Password:  "",
	})
}
