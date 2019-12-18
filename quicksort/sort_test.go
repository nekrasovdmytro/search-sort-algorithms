package quicksort

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Sort success path",
			args: args{in: []int{3, 4, 5, 6, 5, 7}},
			want: []int{3, 4, 5, 5, 6, 7},
		},
		{
			name: "Sort all the same",
			args: args{in: []int{1, 1, 1, 1}},
			want: []int{1, 1, 1, 1},
		},
		{
			name: "Sort mixed up",
			args: args{in: []int{4234, 45345, 23423, 55, 2323, 5343, 2344, 5, 3, 5, 6, 1}},
			want: []int{1, 3, 5, 5, 6, 55, 2323, 2344, 4234, 5343, 23423, 45345},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkQuickSort(b *testing.B) {
	benchmarks := []struct {
		name string
		numberOfSortedEl int
	}{
		{
			  name: "1000",
			  numberOfSortedEl:1000,
		},
		{
			name: "2000",
			numberOfSortedEl:2000,
		},
		{
			name: "3000",
			numberOfSortedEl:3000,
		},
		{
			name: "10000",
			numberOfSortedEl:10000,
		},
		{
			name: "50000",
			numberOfSortedEl:50000,
		},
		{
			name: "100000",
			numberOfSortedEl:100000,
		},
	}
	for _, bm := range benchmarks {
		var in []int
		for j := 0; j < 10000; j++ {
			in = append(in, rand.Int())
		}
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				QuickSort(in)
			}
		})
	}
}

