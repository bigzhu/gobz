package websocketbz

import melody "gopkg.in/olahol/melody.v1"

// 转换为按照 key 给对应的一系列 ws session 发送消息
func broadcastByKey(key string, msg string) {
	M.BroadcastFilter([]byte(msg), func(s *melody.Session) bool {
		_, exists := s.Get(key)
		return exists
	})
}
