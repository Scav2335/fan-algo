package stacks

import (
	"strconv"
)

func Dijkstra(input []string) float64 {
	values := NewArrayStack[float64]()
	operators := NewArrayStack[string]()

	for _, s := range input {
		switch s {
		case "(":
			break
		case "*", "+":
			operators.Push(s)
			break
		case ")":
			operator := operators.Pop()
			if operator == "*" {
				values.Push(values.Pop() * values.Pop())
			} else if operator == "+" {
				values.Push(values.Pop() + values.Pop())
			}
			break
		default:
			i, err := strconv.ParseFloat(s, 64)
			if err != nil {
				panic(err)
			}

			values.Push(i)
			break
		}
	}

	return values.Pop()
}
