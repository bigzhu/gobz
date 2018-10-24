package httpbz

import (
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// PostForm 包装一下
func PostForm(theURL string, form url.Values, client *http.Client) (data string, statusCode int, err error) {
	if client == nil {
		client = &http.Client{}
	}
	response, err := client.PostForm(theURL, form)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	defer closeBody(response.Body)
	statusCode = response.StatusCode
	data, err = readBody(response.Body)
	return
}
