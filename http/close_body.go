package http

import (
	"io"
	"log"
)

func closeBody(body io.ReadCloser) {
	if err := body.Close(); err != nil {
		log.Println(err)
	}
}
