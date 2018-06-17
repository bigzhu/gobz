package httpbz

import (
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// PostForm 包装一下
func PostForm(theURL string, form url.Values, client *http.Client) (data string, err error) {
	if client == nil {
		client = &http.Client{}
	}
	response, err := client.PostForm(theURL, form)
	defer closeBody(response.Body)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	data, err = readBody(response.Body)
	return
}
