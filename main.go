package main

import (
	"fmt"
	"log"

	"./computation"
)

func main() {
	// var numbers = []int{50, 55, 9, 20, 69, 77, 20, 24, 5, 73, 29, 41, 95, 18, 36, 27, 82, 45, 7, 2, 70, 8, 79, 50, 75, 1, 49, 17, 60, 5, 92, 67, 6, 32, 51, 31, 94, 69, 3, 11, 80, 33, 88, 86, 97, 11, 1, 48, 96, 40, 41, 41, 88, 59, 37, 8, 34, 57, 38, 28, 19, 78, 28, 37, 2, 87, 89, 37, 92, 44, 80, 10, 21, 53, 71, 12, 51, 24, 55, 41, 82, 66, 77, 48, 99, 12, 11, 65, 94, 19, 39, 94, 37, 60, 83, 91, 20, 88, 17, 39}
	var numbers = []int{50, 55, 9, 20, 69, 77, 20, 24, 5, 73, 29, 41, 95, 18, 36, 27}
	fmt.Println(numbers)
	r, err := computation.GetRepeatNumber(numbers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
	fmt.Println("Done")
}
