package arry

/*
数组查询包含
*/
func Contains(s []string, e string) bool {
	_rt := false
	for _, a := range s {
		if a == e {
			_rt = true
		}
	}
	return _rt
}
