package websocketbz

import melody "gopkg.in/olahol/melody.v1"

// BroadcastByKey 转换为按照 key 给对应的一系列 ws session 发送消息
func BroadcastByKey(key string, msg string) {
	M.BroadcastFilter([]byte(msg), func(s *melody.Session) bool {
		_, exists := s.Get(key)
		return exists
	})
}
