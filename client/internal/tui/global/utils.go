package global

import (
	"strconv"
	"time"
)

func CanConvertToInt32(strings ...string) bool {
	for _, str := range strings {
		_, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			return false
		}
	}
	return true
}

func ParseDate(dateString string) (time.Time, error) {
	layout := "2006-1-2"
	return time.Parse(layout, dateString)
}
