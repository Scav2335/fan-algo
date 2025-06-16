package stacks

import (
	"math"
	"testing"
)

func TestDijkstra(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "TestDijkstra",
			args: args{
				[]string{"(", "101.23", "+", "(", "(", "2", "+", "3", ")", "*", "(",
					"4", "*", "5", ")", ")", ")"},
			},
			want: 201.23,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dijkstra(tt.args.input); Round(got, 6) != Round(tt.want, 6) {
				t.Errorf("Dijkstra() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Round(x float64, n int) float64 {
	pow10n := math.Pow10(n)
	return math.Round(x*pow10n) / pow10n
}
