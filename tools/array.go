package tools

func InStringArray(list []string, findme string) bool {
	for _, b := range list {
		if b == findme {
			return true
		}
	}
	return false
}

func IndexOf(list []string, findme string) int {
	for i, b := range list {
		if b == findme {
			return i
		}
	}
	return -1
}
