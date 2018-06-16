package httpbz

import (
	"io"
	"io/ioutil"

	"github.com/pkg/errors"
)

func readBody(body io.ReadCloser) (data string, err error) {
	content, err := ioutil.ReadAll(body)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	data = string(content)
	return
}
