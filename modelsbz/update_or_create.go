package modelsbz

// UpdateOrCreate if not exist create else update
func UpdateOrCreate(where interface{}, o interface{}) (err error) {
	run := DB.Where(where).Update(o)
	err = run.Error
	if err != nil {
		return
	}
	if run.RowsAffected == 0 {
		err = DB.Create(o).Error
	}
	return
}
