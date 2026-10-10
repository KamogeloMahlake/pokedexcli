package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	client http.Client
}

func (c *Client) GetData(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return []byte{}, fmt.Errorf("New error: %v", err)
	}
	res, err := c.client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer res.Body.Close()
	
	data, err := io.ReadAll(res.Body)

	if err != nil {
		return []byte{}, err
	}
	return data, nil
	

}
