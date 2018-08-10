package httpbz

import (
	"bytes"
	"net/http"

	"github.com/pkg/errors"
)

// Post 包装一下
func Post(url string, jsonData string, client *http.Client) (data string, err error) {
	if client == nil {
		client = &http.Client{}
	}
	var jsonStr = []byte(jsonData)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	response, err := client.Do(req)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	defer closeBody(response.Body)
	data, err = readBody(response.Body)
	return
}
