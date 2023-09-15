package helper

import (
	"fmt"
	"strconv"
)

func FormatRupiah(num int) string {
	numStr := strconv.Itoa(num)

	formatted := ""

	for i, char := range numStr {
		if (len(numStr)-i)%3 == 0 && i != 0 {
			formatted += "."
		}

		formatted += string(char)
	}

	return formatted
}

func Persentase(num int, num1 int) string {
	persentase := float64(num) / float64(num1) * 100.0

	if persentase < 1 && persentase != 0 {
		return fmt.Sprintf("%.1f%%", persentase)
	} else {
		return fmt.Sprintf("%d%%", int(persentase))
	}
}
