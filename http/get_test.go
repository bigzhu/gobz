package http

import "testing"

func TestGet(t *testing.T) {
	data, err := Get("http://cardano.lorstone.com/api/v1/wallets/Ae2tdPwUPEZMKaLoy2cZZTS9TkNkd4HmKwaBB7U1UinRdHUU56hqqTw9fu5/accounts", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}
