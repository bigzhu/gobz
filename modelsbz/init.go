package modelsbz

import "github.com/bigzhu/gobz/service/callback"

// MigrateAll 同步模型
func MigrateAll() (err error) {
	err = DB.Debug().AutoMigrate(
		&callback.CallbackModel{}, // 回调的模型
	//&FollowWho{},
	//&User{},
	//&InfluencerSocial{},
	//&Influencer{},
	//&Message{},
	).Error
	return
}
