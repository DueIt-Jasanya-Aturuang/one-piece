package helper

import (
	"testing"

	"github.com/google/uuid"
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

func TestUUID(t *testing.T) {
	// uuidSatori := uuidsatori.NewV4().String()
	// 5cba7dd8-30b2-4920-84b9-1049490b1b85
	randomUUID := "5cba7dd8-30b2-4920-84b9-104949012385"
	res, err := uuid.Parse(randomUUID)
	t.Log(err)
	t.Log(res)
}
