package oauthbz

import (
	"time"

	"github.com/bigzhu/gobz/modelsbz"
)

// OauthInfo Oauth 的用户信息
type OauthInfo struct {
	modelsbz.Base
	OutID     string `json:"out_id" gorm:""` // 外部系统的 id
	Email     string `json:"email" gorm:"not null"`
	Name      string `json:"name" gorm:""`
	AvatarURL string `json:"avatar_url"`           // 头像 url
	Link      string `json:"link" gorm:""`         // 个人页面的地址
	Type      string `json:"type" gorm:"not null"` // oauth 类型 google twitter github
}

// GoogleOauthInfo Oauth 的用户信息
type GoogleOauthInfo struct {
	ID            string `json:"id" `
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name" `
	FamilyName    string `json:"family_name" `
	Link          string `json:"link" `
	Picture       string `json:"picture" `
	Gender        string `json:"gender" `
	Locale        string `json:"locale" `
}

// GithubOauthInfo github 返回的信息
type GithubOauthInfo struct {
	Login             string      `json:"login"`
	ID                int         `json:"id"`
	NodeID            string      `json:"node_id"`
	AvatarURL         string      `json:"avatar_url"`
	GravatarID        string      `json:"gravatar_id"`
	URL               string      `json:"url"`
	HTMLURL           string      `json:"html_url"`
	FollowersURL      string      `json:"followers_url"`
	FollowingURL      string      `json:"following_url"`
	GistsURL          string      `json:"gists_url"`
	StarredURL        string      `json:"starred_url"`
	SubscriptionsURL  string      `json:"subscriptions_url"`
	OrganizationsURL  string      `json:"organizations_url"`
	ReposURL          string      `json:"repos_url"`
	EventsURL         string      `json:"events_url"`
	ReceivedEventsURL string      `json:"received_events_url"`
	Type              string      `json:"type"`
	SiteAdmin         bool        `json:"site_admin"`
	Name              string      `json:"name"`
	Company           string      `json:"company"`
	Blog              string      `json:"blog"`
	Location          string      `json:"location"`
	Email             string      `json:"email"`
	Hireable          interface{} `json:"hireable"`
	Bio               string      `json:"bio"`
	PublicRepos       int         `json:"public_repos"`
	PublicGists       int         `json:"public_gists"`
	Followers         int         `json:"followers"`
	Following         int         `json:"following"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}
