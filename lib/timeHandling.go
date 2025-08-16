package lib

import (
	"time"

  "github.com/piquette/finance-go/datetime"
)

// The difference in time. If the difference is one year, years = 1 and days = 365
type timeDiff struct {
	Years int
	Days int
	Hours int
}

// converts time objects to a time difference that's easier to figure out intervals for
func getApproxTimeDiff(t1, t2 time.Time) timeDiff {
	if t2.Before(t1) {
		t1, t2 = t2, t1
	}

	years := t2.Year() - t1.Year()
	tempDays := t2.YearDay() - t1.YearDay()
	tempHours := t2.Hour() - t1.Hour()
	

	if tempHours < 0 {
		tempHours += 24
		tempDays--
	}
	if tempDays < 0 {
		tempDays += 365 // don't have to be careful about leap years, since just for determining interval
		years--
	}

	days := tempDays + years * 365
	hours := tempHours + days * 365

	return timeDiff{
		Years: years,
		Days: days,
		Hours: hours,
	}
}

// Uses a start date and an end date to get the interval passed into piquette-finance go,
// effectively determining the amount of points received from the data
func DetermineInterval(startDate, endDate time.Time) datetime.Interval {
	timeDiff := getApproxTimeDiff(startDate, endDate)

	// Attempts to keep resolution between 40 and 100 points
	if timeDiff.Hours <= 1 {
		return datetime.OneMin
	} else if timeDiff.Hours <= 3 {
		return datetime.TwoMins
	} else if timeDiff.Days <= 1 {
		return datetime.FiveMins
	} else if timeDiff.Days <= 3 {
		return datetime.FifteenMins
	} else if timeDiff.Days <= 7 {
		return datetime.ThirtyMins
	} else if timeDiff.Days <= 23 {
		return datetime.OneHour
	} else if timeDiff.Days <= 100 {
		return datetime.OneDay
	} else if timeDiff.Years <= 2 {
		return datetime.FiveDay
	} else if timeDiff.Years <= 10 {
		return datetime.OneMonth
	} else if timeDiff.Years <= 20 {
		return datetime.ThreeMonth
	} else if timeDiff.Years <= 40 {
		return datetime.SixMonth
	} else if timeDiff.Years <= 80 {
		return datetime.OneYear
	} else if timeDiff.Years <= 200 {
		return datetime.TwoYear
	} else if timeDiff.Years <= 400 {
		return datetime.FiveYear
	}
	return datetime.TenYear
}
