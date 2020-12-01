package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func main() {
	inputs := loadInput()

	var found int
	for _, v1 := range inputs {
		if found != 0 {
			break
		}
		for _, v2 := range inputs {
			if v1+v2 == 2020 {
				log.Printf("2 values that add to 2020 are %v * %v = %v", v1, v2, v1*v2)
				found = v1 * v2
			}
		}

	}

	found = 0
	for _, v1 := range inputs {
		if found != 0 {
			break
		}
		for _, v2 := range inputs {
			if found != 0 {
				break
			}
			for _, v3 := range inputs {
				if v1+v2+v3 == 2020 {
					log.Printf("3 values that add to 2020 are %v * %v * %v = %v", v1, v2, v3, v1*v2*v3)
					found = v1 * v2
				}
			}
		}

	}

}

func loadInput() []int {
	dat, err := ioutil.ReadFile("./inputs.json")
	if err != nil {
		panic(err)
	}

	type D struct {
		Data []int
	}

	var file D
	err = json.Unmarshal(dat, &file)
	if err != nil {
		panic(err)
	}

	return file.Data
}
