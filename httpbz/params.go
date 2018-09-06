package httpbz

import "net/url"

// AddParamsToUrl get 时候向 url 里面添加参数
func AddParamsToUrl(old string, params map[string]string) (new string, err error) {
	u, err := url.Parse(old)
	if err != nil {
		return
	}
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return
	}
	for k, v := range params {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()
	new = u.String()
	return
}
