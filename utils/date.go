package utils

import "time"

func FormattedDate() string {
	now := time.Now()
	return now.Format("2006-01-02 15:04:05")
}

func ExpiryCardDate() string {
	now := time.Now()
	one_year := now.AddDate(5,0,0)
	
	return one_year.Format("01/06")
}