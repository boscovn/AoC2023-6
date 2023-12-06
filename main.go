package main

import (
	"fmt"
)

func beatsRecord(distance int, timePressed int, timeAvailable int) bool {
	return distance < (timeAvailable-timePressed)*timePressed
}

func main() {
	// times := []int{7, 15, 30}
	// distances := []int{9, 40, 200}
	times := []int{50, 74, 86, 85}
	distances := []int{242, 1017, 1691, 1252}

	mul := 1
	for i, v := range times {
		sum := 0
		for j := 1; j < v; j++ {
			if beatsRecord(distances[i], j, v) {
				sum += 1
			}
		}
		if sum != 0 {
			fmt.Println(sum)
			mul *= sum
		}
	}

	fmt.Println(mul)
}
