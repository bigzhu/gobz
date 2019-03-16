package modelsbz

// UpdateOrCreate if not exist create else update
func UpdateOrCreate(where interface{}, o interface{}) (err error) {
	if DB.Model(o).Where(where).Update(o).RowsAffected == 0 {
		err = DB.Create(o).Error
	} else {
		err = DB.Model(o).Where(where).First(o).Error
	}
	return
}
