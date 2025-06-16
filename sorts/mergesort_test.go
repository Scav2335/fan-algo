package sorts

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	d := []Comparable[Int]{Int(18), Int(32), Int(78), Int(69), Int(6), Int(9), Int(8),
		Int(5), Int(4), Int(3), Int(2), Int(19), Int(17), Int(34), Int(29), Int(83), Int(52)}
	MergeSort(d)
	fmt.Println(d)

	c := []Comparable[String]{String("def"), String("abc"), String("ghi"), String("jkl"),
		String("mno"), String("pqrs")}
	MergeSort(c)
	fmt.Println(c)
}

func TestBottomUpMergeSort(t *testing.T) {
	d := []Comparable[Int]{Int(18), Int(32), Int(78), Int(69), Int(6), Int(9), Int(8),
		Int(5), Int(4), Int(3), Int(2), Int(19), Int(17), Int(34), Int(29), Int(83), Int(52)}
	BottomUpMergeSort(d)
	fmt.Println(d)

	c := []Comparable[String]{String("def"), String("abc"), String("ghi"), String("jkl"),
		String("mno"), String("pqrs")}
	BottomUpMergeSort(c)
	fmt.Println(c)
}
