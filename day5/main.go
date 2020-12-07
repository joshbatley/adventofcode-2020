package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
)

func main() {
	inputs := loadInputs()
	var seatID float64
	var seatIDs []float64
	for _, s := range inputs {
		row := string(s[0:7])
		col := string(s[7:])
		r := bsp(row, []float64{0, 127})
		c := bsp(col, []float64{0, 7})
		t := r*8 + c
		seatIDs = append(seatIDs, t)
		if t > seatID {
			seatID = t
		}
	}

	sort.Float64s(seatIDs)
	var posSeatID float64
	offset := seatIDs[0]
	len := float64(len(seatIDs)) + offset
	for i := offset; i < len; i++ {
		if float64(i) != seatIDs[int(i-offset)] {
			posSeatID = float64(i)
			break
		}
	}
	log.Println("you seatID is", posSeatID)
	log.Println("the largest seatID is", seatID)
}

func bsp(s string, r []float64) float64 {
	for _, l := range s {
		switch l {
		case 'F', 'L':
			r[1] = math.Floor((r[0] + r[1]) / 2)
		case 'B', 'R':
			r[0] = math.Round((r[0] + r[1]) / 2)
		}
	}
	return r[0]
}

func loadInputs() []string {
	file, err := os.Open("./inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var d []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		d = append(d, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return d
}
