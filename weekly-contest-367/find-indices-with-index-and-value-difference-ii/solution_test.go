package findindiceswithindexandvaluedifferenceii

import (
	"reflect"
	"testing"
)

func Test_findIndices(t *testing.T) {
	type args struct {
		nums            []int
		indexDifference int
		valueDifference int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findIndices(tt.args.nums, tt.args.indexDifference, tt.args.valueDifference); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}
