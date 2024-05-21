package kata

import (
	"fmt"
	"strings"
)

func SortMyString(s string) string {
	var even string
	var odd string
	str := strings.Split(s, "")
	for i := 0; i < len(str); i++ {
		if i%2 == 0 {
			even += str[i]
		} else {
			odd += str[i]
		}

	}
	result := fmt.Sprintf("%s %s", even, odd)
	return result
}
