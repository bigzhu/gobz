package modelsbz

// UpdateOrCreate if not exist create else update
// tell user is exists
func UpdateOrCreate(where interface{}, o interface{}) (exists bool, err error) {
	count := int64(0)
	err = DB.Model(o).Where(where).Count(&count).Error
	if err != nil {
		return
	}
	if count == 0 {
		exists = false
		err = DB.Create(o).Error
	} else {
		exists = true
		err = DB.Debug().Model(o).Where(where).Updates(o).Error
	}
	return
}
