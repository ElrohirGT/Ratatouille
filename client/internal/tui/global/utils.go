package global

import "strconv"

func CanConvertToInt32(strings ...string) bool {
	for _, str := range strings {
		_, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			return false
		}
	}
	return true
}
