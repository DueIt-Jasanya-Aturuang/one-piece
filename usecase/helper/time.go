package helper

import (
	"time"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func TimeDate(periode int) (time.Time, time.Time, error) {
	now := time.Now().UTC()
	startTime := time.Date(now.Year(), now.Month(), periode, 0, 0, 0, 0, time.UTC)
	endTime := startTime.AddDate(0, 1, -1)

	return FormatDate(startTime, endTime)
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

		return time.Time{}, time.Time{}, err
	}
	endTimeNew, err := time.Parse("2006-01-02", endTimeStr)
	if err != nil {
		log.Warn().Msgf("error parse layout 2006-01-02 | err : %v", err)
		return time.Time{}, time.Time{}, err
	}

	return startTimeNew, endTimeNew, nil
}
