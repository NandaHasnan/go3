// package main

// import (
// 	"fmt"
// 	"test/go3/tanggal"
// )

// func main() {
// 	fmt.Println(tanggal.Tanggal(profile.name))
// }

package main

import (
	"fmt"
	"strings"
)

type data struct {
	name        string
	dateOfBirth string
}

func converDate(data string) string{

	var month = map[string]string{
		"01": "January",
		"02": "Februari",
		"03": "Maret",
		"04": "April",
		"05": "Mei",
		"06": "Juni",
		"07": "Juli",
		"08": "Agustus",
		"09": "September",
		"10": "Oktober",
		"11": "November",
		"12": "Desember",
	}

	date := strings.Split(data, "/")
	if i := 0; i < len(month); i++ {
		if month[i] == date[1] {
			date[1] = month[i+1]
			fmt.Println(strings.join(date, " "))
			return
		}
	}
}

func main() {
	profile := data{
		name:        "Fazz",
		dateOfBirth: "02/06/2000",
	}

	fmt.Println(profile.name)
	fmt.Println(profile.dateOfBirth)
	fmt.Println(profile.converDate)
}
