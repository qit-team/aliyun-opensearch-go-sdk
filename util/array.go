package util

type ArrayUtil struct {
}

func (util *ArrayUtil) InArray(data []string, target string) bool {
	if data == nil || len(data) == 0 {
		return false
	}
	for _, v := range data {
		if v == target {
			return true
		}
	}
	return false
}
