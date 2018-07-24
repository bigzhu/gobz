package oauthbz

// OauthInfo Oauth 的用户信息
type OauthInfo struct {
	OID           uint   `gorm:"primary_key" json:"oid"`
	ID            string `json:"id" gorm:""`
	Email         string `json:"email" gorm:"not null"`
	VerifiedEmail bool   `json:"verified_email" gorm:"verified_email"`
	Name          string `json:"name" gorm:""`
	GivenName     string `json:"given_name" gorm:""`
	FamilyName    string `json:"family_name" gorm:""`
	Link          string `json:"link" gorm:""`
	Picture       string `json:"picture" gorm:""`
	Gender        string `json:"gender" gorm:""`
	Locale        string `json:"locale" gorm:""`
	Type          string `json:"type" gorm:"not null"`
}
