package statistics

import (
	"fmt"
	"testing"
)

func TestGetState(t *testing.T) {
	var numbers = []float64{1.2353, 1.2356, 68.235, 23.6234, 2.346234}
	stat1 := GetState(numbers)
	fmt.Println(stat1)
	fmt.Println(stat1.numbers)
	fmt.Println(stat1.mean)
	fmt.Println(stat1.median)
}
