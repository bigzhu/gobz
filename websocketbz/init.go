package websocketbz

import melody "gopkg.in/olahol/melody.v1"

// M Melody 的实例
var M *melody.Melody

// 初始化
func init() {
	// 确保实例化
	if M == nil {
		M = melody.New()
		// 把注册的方法也设置了.
		// 这里不考虑客户端向服务端发数据的情况了. 基本需求都是服务端向客户端通知
		M.HandleMessage(DoRegister)
	}
}
