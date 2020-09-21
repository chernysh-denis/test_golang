package computation

import (
	"fmt"
	"sort"
)

// function to find a repeating element
func GetRepeatNumber(arr []int) (string, error) {
	// fix the number of repetitions of each number
	mapTemp := make(map[int]int)
	for _, v := range arr {
		if _, ok := mapTemp[v]; ok {
			mapTemp[v]++
		} else {
			mapTemp[v] = 1
		}
	}
	// create an array with repetitions and sort it to get the maximum number of repetitions
	var arrT []int
	for _, v := range mapTemp {
		arrT = append(arrT, v)
	}
	sort.Ints(arrT)
	maxRep := arrT[len(arrT)-1]
	// get an array with numbers that are repeated the maximum number of times
	var arrL []int
	for i, v := range mapTemp {
		if v == maxRep {
			arrL = append(arrL, i)
		}
	}
	// sort to get the maximum element, if there are several
	sort.Ints(arrL)
	r := fmt.Sprintf("%d", arrL[len(arrL)-1])

	return r, nil
}
