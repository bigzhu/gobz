package modelsbz

// UpdateOrCreate if not exist create else update
// tell user is exists
func UpdateOrCreate(where interface{}, o interface{}) (exists bool, err error) {
	if DB.Model(o).Where(where).Update(o).RowsAffected == 0 {
		exists = false
		err = DB.Create(o).Error
	} else {
		exists = true
		err = DB.Model(o).Where(where).First(o).Error
	}
	return
}
