package websocketbz

import (
	"encoding/json"
	"log"

	melody "gopkg.in/olahol/melody.v1"
)

// Register 注册 websocket 的 json 结构
type Register struct {
	Oper string `json:"oper"`
	Key  string `json:"key"`
}

// DoRegister 注册申请的 websocket
func DoRegister(s *melody.Session, msg []byte) {
	var register Register
	err := json.Unmarshal(msg, &register)
	if err != nil {
		log.Println(err)
		return
	}
	s.Set(register.Key, "bigzhu")
	log.Println("register new websocket", s.Request.URL.Path, "key=", register.Key)
	return
}
