package telegram

import "testing"

func TestSendToUser(t *testing.T) {
	err := SendToUser("bigzhu好啊, 你在干什么?")
	if err != nil {
		t.Fatal(err)
	}
}
