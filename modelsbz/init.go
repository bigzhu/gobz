package modelsbz


// MigrateAll 同步模型
func MigrateAll() (err error) {
	err = DB.Debug().AutoMigrate(
	//&FollowWho{},
	//&User{},
	//&InfluencerSocial{},
	//&Influencer{},
	//&Message{},
	).Error
	return
}
