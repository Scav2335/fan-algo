package sorts

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	d := []Comparable[Int]{Int(18), Int(29), Int(32), Int(78), Int(69), Int(3), Int(6), Int(9), Int(8),
		Int(5), Int(4), Int(83), Int(3), Int(2), Int(19), Int(17), Int(34), Int(29), Int(83), Int(52),
		Int(3)}
	QuickSort(d)
	fmt.Println(d)

	c := []Comparable[String]{String("def"), String("abc"), String("ghi"), String("jkl"),
		String("mno"), String("pqrs")}
	QuickSort(c)
	fmt.Println(c)
}
