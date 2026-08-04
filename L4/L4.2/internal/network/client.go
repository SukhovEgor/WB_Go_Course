package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"DistributedGrep/internal/app"
)

type Client struct {
	http *http.Client
}

func NewClient() *Client {

	return &Client{

		http: &http.Client{},
	}
}

func (c *Client) Send(

	addr string,

	req app.GrepRequest,

) (*app.GrepResponse, error) {

	body, err := json.Marshal(req)

	if err != nil {
		return nil, err
	}

	resp, err := c.http.Post(

		addr+"/grep",

		"application/json",

		bytes.NewReader(body),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		return nil,

			fmt.Errorf("server returned %d",

				resp.StatusCode)
	}

	var result app.GrepResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}