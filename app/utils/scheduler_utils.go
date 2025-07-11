package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func CalculateCronExpression(interval_value int, interval_unit string) (string, error) {
	var expression string

	switch interval_unit {
	case
		"minutes":
		if interval_value < 0 || interval_value > 59 {
			return "", errors.New("Wrong minute value")
		}
		expression = fmt.Sprintf("*/%d * * * *", interval_value)
	case "hours":
		if interval_value < 0 || interval_value > 23 {
			return "", errors.New("Wrong hour value")
		}
		expression = fmt.Sprintf("0 */%d * * *", interval_value)
	case "days":
		if interval_value < 1 || interval_value > 31 {
			return "", errors.New("Wrong day value")
		}
		expression = fmt.Sprintf("0 0 */%d * *", interval_value)
	default:
		return "", errors.New("Invalid interval unit")
	}

	return expression, nil
}

func CalculateNextRun(schedule string, timezone string) (time.Time, error) {
	// Load location
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid timezone: %v", err)
	}

	// Parse schedule
	sched, err := cron.ParseStandard(schedule)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid schedule expression: %v", err)
	}

	// Calculate next run in the specified timezone
	now := time.Now().In(loc)
	return sched.Next(now), nil
}
