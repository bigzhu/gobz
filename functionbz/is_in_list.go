package functionbz

//IsIn 是否在列表中
func IsIn(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
