package test

import (
	"strings"
	"time"
)

func Splict(s, sep string) (result []string) {
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+1:]
		i = strings.Index(s, sep)
		time.Sleep(1000)
	}
	result = append(result, s)
	return
}
