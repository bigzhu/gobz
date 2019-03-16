package modelsbz

import "github.com/jinzhu/gorm"

// UpdateOrCreate if not exist create else update
func UpdateOrCreate(where interface{}, o *interface{}) (err error) {
	var i interface{}
	err = DB.Where(where).First(i).Error
	// exist update
	if err == nil {
		err = DB.Where(where).Update(o).Error
		return
	}
	// not exist create
	if gorm.IsRecordNotFoundError(err) {
		err = DB.Create(&o).Error
	}
	return
}
