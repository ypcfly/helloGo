package models

import "github.com/astaxie/beego/session"

// 定义
var customSession *session.Manager

func InitSession() {
	// session数据初始化
	sessionConfig := &session.ManagerConfig{
		CookieName:      "custom_sessionId",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "./tmp",
	}
	// 使用内存存储
	customSession, _ = session.NewManager("memory", sessionConfig)
	go customSession.GC()
}
