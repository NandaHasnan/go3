package struck

import "fmt"

type bio struct {
	results []struct {
		name map[string]string
	}
}

func struck() {
	data := bio{
		results: []struct {
			name map[string]string
		}{
			{name: map[string]string{"firsName": "fazztrat"}},
			{name: map[string]string{"firsName": "fazztrac"}},
			{name: map[string]string{"firsName": "fazztrct"}},
			{name: map[string]string{"firsName": "fazzract"}},
			{name: map[string]string{"firsName": "fazzract"}},
			{name: map[string]string{"firsName": "fazztract"}},
		},
	}
	fmt.Println(data.results[5].name["firsName"])
}
