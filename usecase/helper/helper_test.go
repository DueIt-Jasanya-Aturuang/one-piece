package helper

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func TestPersentase(t *testing.T) {
	num := 500000
	num1 := 100000

	res := Persentase(num, num1)
	t.Log(res)
}

func TestFormatRupiah(t *testing.T) {
	rp := 2000000000000

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

func TestTime(t *testing.T) {

	startTime, endTime, err := TimeDate(6)
	assert.NoError(t, err)
	t.Log(startTime)
	t.Log(endTime)
	startTime = startTime.Add(-24 * time.Hour)
	endTime = endTime.Add(-24 * time.Hour)
	t.Log(startTime)
	t.Log(endTime)
	t.Log(time.Now().Add(-24 * time.Hour).UTC())
}

func TestTimeDateByTypeFilter(t *testing.T) {
	t.Run(util.DayNow, func(t *testing.T) {
		startTime, endTime, err := TimeDateByTypeFilter("hari-ini")
		assert.NoError(t, err)
		t.Logf("start : %v | end : %v", startTime, endTime)
	})
	t.Run(util.Kemarin, func(t *testing.T) {
		startTime, endTime, err := TimeDateByTypeFilter("kemarin")
		assert.NoError(t, err)
		t.Logf("start : %v | end : %v", startTime, endTime)
	})
	t.Run(util.MingguNow, func(t *testing.T) {
		startTime, endTime, err := TimeDateByTypeFilter("minggu-ini")
		assert.NoError(t, err)
		t.Logf("start : %v | end : %v", startTime, endTime)
	})
	t.Run(util.BulanNow, func(t *testing.T) {
		startTime, endTime, err := TimeDateByTypeFilter("bulan-ini")
		assert.NoError(t, err)
		t.Logf("start : %v | end : %v", startTime, endTime)
	})
	t.Run(util.BulanNow, func(t *testing.T) {
		startTime, endTime, err := TimeDateByTypeFilter("bulanasd-ini")
		assert.NoError(t, err)
		t.Logf("start : %v | end : %v", startTime, endTime)
	})
}
