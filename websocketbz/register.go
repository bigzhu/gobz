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

// Resp 注册的结果如何
type Resp struct {
	Result string `json:"result"`
}

// DoRegister 注册申请的 websocket
func DoRegister(s *melody.Session, msg []byte) {
	var register Register
	err := json.Unmarshal(msg, &register)
	if err != nil {
		log.Println(err)
		result, _ := json.Marshal(Resp{Result: err.Error()})
		s.Write(result)
		return
	}
	s.Set(register.Key, "bigzhu")
	log.Println("register new websocket", s.Request.URL.Path, "key=", register.Key)
	result, _ := json.Marshal(Resp{Result: "success"})
	s.Write(result)
	return
}
