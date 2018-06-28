package modelsbz

import "testing"

func TestMigrateAll(t *testing.T) {
	err := MigrateAll()
	if err != nil {
		t.Error(err)
	}
}
