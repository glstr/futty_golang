package es

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	defaultClient *esClient
)

type ESConfig struct {
	URL string
}

type ESClient interface {
	WriteDoc(index string, doc interface{}) error
	GetDoc(index string, id string, doc interface{}) error
	Search(index string, query interface{}, res interface{}) error
}

func GetClient() ESClient {
	return defaultClient
}

type esClient struct {
	c *elasticsearch.Client
}

func InitESClient(c *ESConfig) error {
	cfg := elasticsearch.Config{
		Addresses: []string{c.URL},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return err
	}

	defaultClient = &esClient{c: client}
	return nil
}

func (c *esClient) WriteDoc(index string, doc interface{}) error {
	body, err := json.Marshal(doc)
	if err != nil {
		return err
	}
	res, err := c.c.Index(index, bytes.NewReader(body), c.c.Index.WithRefresh("true"))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.IsError() {
		b, _ := io.ReadAll(res.Body)
		return fmt.Errorf("es index error: %s", string(b))
	}
	return nil
}

func (c *esClient) GetDoc(index string, id string, doc interface{}) error {
	res, err := c.c.Get(index, id)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.IsError() {
		b, _ := io.ReadAll(res.Body)
		return fmt.Errorf("es get error: %s", string(b))
	}
	var r struct {
		Source json.RawMessage `json:"_source"`
	}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return err
	}
	if len(r.Source) == 0 {
		return fmt.Errorf("empty _source")
	}
	return json.Unmarshal(r.Source, doc)
}

func (c *esClient) Search(index string, query interface{}, res interface{}) error {
	qb, err := json.Marshal(query)
	if err != nil {
		return err
	}
	resp, err := c.c.Search(
		c.c.Search.WithIndex(index),
		c.c.Search.WithBody(bytes.NewReader(qb)),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.IsError() {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("es search error: %s", string(b))
	}
	var sr struct {
		Hits struct {
			Hits []struct {
				Source json.RawMessage `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&sr); err != nil {
		return err
	}
	// Build a JSON array from _source entries and unmarshal into res
	buf := bytes.NewBufferString("[")
	for i, h := range sr.Hits.Hits {
		buf.Write(h.Source)
		if i < len(sr.Hits.Hits)-1 {
			buf.WriteString(",")
		}
	}
	buf.WriteString("]")
	return json.Unmarshal(buf.Bytes(), res)
}
