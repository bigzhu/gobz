package modelsbz

// MigrateAll 同步模型
func MigrateAll() (err error) {
	err = DB.AutoMigrate(
	//&FollowWho{},
	//&User{},
	//&InfluencerSocial{},
	//&Influencer{},
	//&Message{},
	)
	return
}
