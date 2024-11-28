package tanggal

import (
	"fmt"
	"time"
)

type bio struct {
	name string
	// dateOfBirth string
}

var (
	dateOfBirth = "01/12/1999"
	converDate  = "1 Desember 1999"
)

func Tanggal() {
	profile := bio{
		name: "Fazz",
	}
	date := getDate(2, 6, 2000)
	fmt.Println(profile.name)
	fmt.Println(date.Format(dateOfBirth))
	fmt.Println(date.Format(converDate))
}

func getDate(day, month, year int) time.Time {
	date := time.Date(day, time.Month(month), year, 0, 0, 0, 0, time.UTC)
	return date
}
