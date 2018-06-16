package http

import (
	"net/http"

	"github.com/pkg/errors"
)

// Get Get请求
func Get(url string) (data string, err error) {
	response, err := http.Get(url)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	defer closeBody(response.Body)
	data, err = readBody(response.Body)
	return
}
