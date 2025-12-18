package basaname2

import "strings"

func basaname2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {

		s = s[:dot]
	}
	return s
}
