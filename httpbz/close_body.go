package httpbz

import (
	"io"
	"log"
)

// 定义一个复用的关闭 http body 的函数, 屏蔽 error
func closeBody(body io.ReadCloser) {
	if err := body.Close(); err != nil {
		log.Println(err)
	}
}
