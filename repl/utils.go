package repl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func  fetch[T any](url string, structData *T) error {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return fmt.Errorf("New error: %v", err)
	}
	client := http.DefaultClient
	res, err := client.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()
	
	data, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, structData); err != nil {
		return err
	}

	return nil

}
