package sorts

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	d := []Comparable[Int]{Int(18), Int(32), Int(78), Int(69), Int(6),
		Int(5), Int(4), Int(3), Int(2)}
	SelectionSort(d)
	fmt.Println(d)

	c := []Comparable[String]{String("def"), String("abc"), String("ghi")}
	SelectionSort(c)
	fmt.Println(c)
}
