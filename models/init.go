package models

// MigrateAll 同步模型
func MigrateAll() (err error) {
	err = DB.Debug().AutoMigrate(
	//&AnkiMessage{},
	//&FollowWho{},
	//&User{},
	//&InfluencerSocial{},
	//&Influencer{},
	//&Message{},
	).Error
	return
}
