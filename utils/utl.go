package utils

func IF(condition bool,trueval,falseval interface{}) interface{} {
	if condition {
		return trueval
	}
	return falseval
}
func SliceStringContains(a []string,b string) bool {
	for _,t := range a {
		if t == b {
			return true
		}
	}
	return false
}