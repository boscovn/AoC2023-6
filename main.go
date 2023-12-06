package main

import (
	"fmt"
)

func beatsRecord(distance int, timePressed int, timeAvailable int) bool {
	return distance < (timeAvailable-timePressed)*timePressed
}

func main() {
	// time := 71530
	// distance := 940200
	time := 50748685
	distance := 242101716911252

	sum := 0
	for j := 1; j < time; j++ {
		if beatsRecord(distance, j, time) {
			sum += 1
		}
	}

	fmt.Println(sum)
}
