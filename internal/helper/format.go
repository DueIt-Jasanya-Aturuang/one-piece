package helper

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	resp "github.com/jasanya-tech/jasanya-response-backend-golang"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func FormatRupiah(num int) string {
	numStr := strconv.Itoa(num)

	formatted := ""

	if string(numStr[0]) == "-" {
		for i, char := range numStr {
			if (len(numStr)-i)%3 == 0 && i != 1 {
				formatted += "."
			}

			formatted += string(char)
		}
	} else {
		for i, char := range numStr {
			if (len(numStr)-i)%3 == 0 && i != 0 {
				formatted += "."
			}

			formatted += string(char)
		}
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

func TimeDate(periode int) (time.Time, time.Time, error) {
	now := time.Now().UTC()
	startTime := time.Date(now.Year(), now.Month(), periode, 0, 0, 0, 0, time.UTC)
	endTime := startTime.AddDate(0, 1, -1)

	return FormatDate(startTime, endTime)
}

func GetCodePqError(err error) pq.ErrorCode {
	var errPQ *pq.Error

	if errors.As(err, &errPQ) {
		return errPQ.Code
	}

	return ""
}

func TimeDateByTypeFilter(typeFilter string) (time.Time, time.Time, error) {
	now := time.Now().UTC()
	switch typeFilter {
	case util.DayNow:
		startTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
		endTime := startTime
		return startTime, endTime, nil
	case util.Kemarin:
		startTime := time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, time.UTC)
		endTime := startTime
		return startTime, endTime, nil
	case util.MingguNow:
		week := time.Monday - now.Weekday()
		startTime := now.AddDate(0, 0, int(week)).UTC()
		endTime := startTime.AddDate(0, 0, 6).UTC()
		return FormatDate(startTime, endTime)
	case util.BulanNow:
		startTime := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
		endTime := startTime.AddDate(0, 1, -1)
		return FormatDate(startTime, endTime)
	}

	startTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	endTime := startTime
	return FormatDate(startTime, endTime)
}

func FormatDate(startTime, endTime time.Time) (time.Time, time.Time, error) {
	startTimeStr := startTime.Format("2006-01-02")
	endTimeStr := endTime.Format("2006-01-02")

	startTimeNew, err := time.Parse("2006-01-02", startTimeStr)
	if err != nil {
		log.Warn().Msgf("error parse layout 2006-01-02 | err : %v", err)

		return time.Time{}, time.Time{}, resp.HttpErrString(string(resp.S500), resp.S500)
	}
	endTimeNew, err := time.Parse("2006-01-02", endTimeStr)
	if err != nil {
		log.Warn().Msgf("error parse layout 2006-01-02 | err : %v", err)
		return time.Time{}, time.Time{}, resp.HttpErrString(string(resp.S500), resp.S500)
	}

	return startTimeNew, endTimeNew, nil
}
