package modelsbz

import (
	"errors"

	"gorm.io/gorm"
)

// UpdateOrCreate if not exist create else update
// tell user is exists
func UpdateOrCreate(where interface{}, o interface{}) (exists bool, err error) {
	var result interface{}
	err = DB.Model(o).Where(where).First(result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			exists = false
			err = DB.Create(o).Error
		}
	} else {
		exists = true
	}
	return
}
