package mergesort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success path",
			args: args{arr:[]int{2,1,4,5}},
			want: []int{1,2,4,5},
		},
		{
			name: "success path, mix values",
			args: args{arr:[]int{22,44,1,5, 123}},
			want: []int{1, 5, 22, 44, 123},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
