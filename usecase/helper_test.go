package usecase

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase_old/helper"

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

	startTime, endTime, err := helper.TimeDate(6)
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
		startTime, endTime, err := helper.TimeDateByTypeFilter("hari-ini")
		assert.NoError(t, err)
		t.Logf("start : %v | end : %v", startTime, endTime)
	})
	t.Run(util.Kemarin, func(t *testing.T) {
		startTime, endTime, err := helper.TimeDateByTypeFilter("kemarin")
		assert.NoError(t, err)
		t.Logf("start : %v | end : %v", startTime, endTime)
	})
	t.Run(util.MingguNow, func(t *testing.T) {
		startTime, endTime, err := helper.TimeDateByTypeFilter("minggu-ini")
		assert.NoError(t, err)
		t.Logf("start : %v | end : %v", startTime, endTime)
	})
	t.Run(util.BulanNow, func(t *testing.T) {
		startTime, endTime, err := helper.TimeDateByTypeFilter("bulan-ini")
		assert.NoError(t, err)
		t.Logf("start : %v | end : %v", startTime, endTime)
	})
	t.Run(util.BulanNow, func(t *testing.T) {
		startTime, endTime, err := helper.TimeDateByTypeFilter("bulanasd-ini")
		assert.NoError(t, err)
		t.Logf("start : %v | end : %v", startTime, endTime)
	})
}

func TestUlid(t *testing.T) {
	ulid1 := ulid.Make()
	time.Sleep(1 * time.Millisecond)
	ulid2 := ulid.Make()
	fmt.Println(ulid1.Time())
	fmt.Println(ulid1.String())
	fmt.Println(ulid2.Time())
}
