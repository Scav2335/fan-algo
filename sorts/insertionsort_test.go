package sorts

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	type args[T any] struct {
		d []Comparable[T]
	}
	type testCase[T any] struct {
		name string
		args args[T]
	}
	tests := []testCase[Int]{
		{
			name: "InsertionSort",
			args: args[Int]{
				d: []Comparable[Int]{Int(5), Int(4), Int(3), Int(2)},
			},
		},
		{
			name: "InsertionSort",
			args: args[Int]{
				d: []Comparable[Int]{Int(18), Int(32), Int(78), Int(69), Int(6),
					Int(5), Int(4), Int(3), Int(2)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertionSort(tt.args.d)
			fmt.Println(tt.args.d)
		})
	}

	c := []Comparable[String]{String("def"), String("abc"), String("ghi")}
	InsertionSort(c)
	fmt.Println(c)
}
