package modelsbz

// OauthInfo Oauth 的用户信息
type OauthInfo struct {
	Base
	OutID     string `json:"out_id" gorm:""` // 外部系统的 id
	Email     string `json:"email" gorm:"not null"`
	Name      string `json:"name" gorm:""`
	AvatarURL string `json:"avatar_url"`           // 头像 url
	Link      string `json:"link" gorm:""`         // 个人页面的地址
	Type      string `json:"type" gorm:"not null"` // oauth 类型 google twitter github
}
