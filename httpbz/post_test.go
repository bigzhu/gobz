package httpbz

import "testing"

func TestPost(t *testing.T) {
	params := `
		{
		"accountIndex": 3566163324,
		"walletId": "Ae2tdPwUPEZMKaLoy2cZZTS9TkNkd4HmKwaBB7U1UinRdHUU56hqqTw9fu5",
		"spendingPassword": ""
		}
	`
	data, _, err := Post("http://cardano.lorstone.com/api/v1/addresses", params, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("data:", data)
}
