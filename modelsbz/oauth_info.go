package modelsbz

// OauthInfo Oauth 的用户信息
type OauthInfo struct {
	Base
	OutID     string // 外部系统的 id
	Email     string `gorm:"not null"`
	Name      string
	AvatarURL string // 头像 url
	Link      string // 个人页面的地址
	Type      string `gorm:"not null"` // oauth 类型 google twitter github
}
