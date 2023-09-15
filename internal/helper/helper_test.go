package helper

import (
	"testing"
)

func TestPersentase(t *testing.T) {
	num := 500000
	num1 := 100000

	res := Persentase(num, num1)
	t.Log(res)
}

func TestFormatRupiah(t *testing.T) {
	rp := -1000000000

	res := FormatRupiah(rp)
	t.Log(res)
}
