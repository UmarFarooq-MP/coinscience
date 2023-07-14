package http

import (
	"encoding/json"
	"github.com/imroc/req/v3"
	"log"
	"time"
)

type Client struct {
	client *req.Client //req.DevMode()
}

func (c *Client) Get(url string) (interface{}, error) {

	resp, err := c.client.R(). // Use R() to create a request.
					Get(url)
	if err != nil {
		return nil, err
	}

	var data interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Println("Error while Decoding the response")
		return nil, err
	}
	return data, nil
}

func NewHttpClient() *Client {
	return &Client{
		client: req.DefaultClient(),
	}
}

func NewHttpClientWithTimeout(second int) *Client {
	return &Client{
		client: req.DefaultClient().SetTimeout(time.Duration(second) * time.Second),
	}
}
