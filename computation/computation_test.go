package computation

import (
	"strconv"
	"strings"
	"testing"
)

var testOkInput = `50, 55, 9, 20, 69, 77, 20, 24, 5, 73, 29, 41, 95, 18, 36, 27`

var testOkResult = `20`

func TestOK(t *testing.T) {
	in := testOkInput
	arrT := strings.Split(in, ",")
	var arrG []int
	for _, v := range arrT {
		i, _ := strconv.Atoi(strings.TrimSpace(v))
		arrG = append(arrG, i)
	}
	out, err := GetRepeatNumber(arrG)
	if err != nil {
		t.Errorf("Test OK failed: %s", err)
	}
	if out != testOkResult {
		t.Errorf("Test OK failed, result not match")
	}
}
