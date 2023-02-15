package handler

import (
	"reflect"
	"testing"
)

func TestGetSliceWithoutNesting(t *testing.T) {
	original := []interface{}{1, 2, 3, []int{4, 5}}
	got := getSliceWithoutNesting(original)
	want := []interface{}{1, 2, 3}

	if reflect.DeepEqual(got, want) == false {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
