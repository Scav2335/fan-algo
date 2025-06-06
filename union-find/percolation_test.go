package union_find

import (
	"fmt"
	"math/rand/v2"
	"reflect"
	"testing"
)

func TestNewPercolator(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want *Percolator
	}{
		{
			name: "TestNewPercolator",
			args: args{
				N: 2,
			},
			want: &Percolator{
				Sites: []*Site{
					{
						0, true,
					},
					{
						1, false,
					},
					{
						2, false,
					},
					{
						3, false,
					},
					{
						4, false,
					},
					{
						5, true,
					},
				},
				Sz: []int{1, 1, 1, 1, 1, 1},
				N:  2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPercolator(tt.args.N); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPercolator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPercolator_Connected(t *testing.T) {
	type fields struct {
		Sites []*Site
		Sz    []int
		N     int
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "TestPercolator_Connected",
			fields: fields{
				Sites: []*Site{
					{
						0, true,
					},
					{
						1, false,
					},
					{
						2, false,
					},
					{
						3, false,
					},
					{
						4, false,
					},
					{
						5, true,
					},
				},
				Sz: []int{1, 1, 1, 1, 1, 1},
				N:  2,
			},
			args: args{0, 5},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Percolator{
				Sites: tt.fields.Sites,
				Sz:    tt.fields.Sz,
				N:     tt.fields.N,
			}
			if got := p.Connected(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Connected() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPercolator_Open(t *testing.T) {
	type fields struct {
		Sites []*Site
		Sz    []int
		N     int
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantOk bool
	}{
		{
			name: "TestPercolator_Open0",
			fields: fields{
				Sites: []*Site{
					{
						0, true,
					},
					{
						1, false,
					},
					{
						2, false,
					},
					{
						3, false,
					},
					{
						4, false,
					},
					{
						5, true,
					},
				},
				Sz: []int{1, 1, 1, 1, 1, 1},
				N:  2,
			},
			args:   args{0},
			wantOk: false,
		},
		{
			name: "TestPercolator_Open3",
			fields: fields{
				Sites: []*Site{
					{
						0, true,
					},
					{
						1, false,
					},
					{
						2, false,
					},
					{
						3, false,
					},
					{
						4, false,
					},
					{
						5, true,
					},
				},
				Sz: []int{1, 1, 1, 1, 1, 1},
				N:  2,
			},
			args:   args{3},
			wantOk: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Percolator{
				Sites: tt.fields.Sites,
				Sz:    tt.fields.Sz,
				N:     tt.fields.N,
			}
			if gotOk := p.Open(tt.args.i); gotOk != tt.wantOk {
				t.Errorf("Open() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestPercolator_Root(t *testing.T) {
	type fields struct {
		Sites []*Site
		Sz    []int
		N     int
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "TestPercolator_Root0",
			fields: fields{
				Sites: []*Site{
					{
						0, true,
					},
					{
						1, false,
					},
					{
						2, false,
					},
					{
						3, false,
					},
					{
						4, false,
					},
					{
						5, true,
					},
				},
				Sz: []int{1, 1, 1, 1, 1, 1},
				N:  2,
			},
			args: args{0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Percolator{
				Sites: tt.fields.Sites,
				Sz:    tt.fields.Sz,
				N:     tt.fields.N,
			}
			if got := p.Root(tt.args.i); got != tt.want {
				t.Errorf("Root() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPercolator_Union(t *testing.T) {
	type fields struct {
		Sites []*Site
		Sz    []int
		N     int
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "TestPercolator_Union",
			fields: fields{
				Sites: []*Site{
					{
						0, true,
					},
					{
						1, false,
					},
					{
						2, false,
					},
					{
						3, false,
					},
					{
						4, false,
					},
					{
						5, true,
					},
				},
				Sz: []int{1, 1, 1, 1, 1, 1},
				N:  2,
			},
			args: args{
				i: 1,
				j: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Percolator{
				Sites: tt.fields.Sites,
				Sz:    tt.fields.Sz,
				N:     tt.fields.N,
			}
			p.Union(tt.args.i, tt.args.j)
		})
	}
}

const N = 1000

func TestPercolator_UnionFind(t *testing.T) {
	var sumOpened float64

	for i := 0; i < N; i++ {
		p := NewPercolator(N)
		for !p.Connected(0, N*N+1) {
			// Random open sites
			if p.Open(rand.IntN(N*N + 1)) {
				sumOpened++
			}
		}
	}

	// Average probability
	fmt.Printf("The average probability is %v", sumOpened/(N*N*N))
}
