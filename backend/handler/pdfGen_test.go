package handler

import (
	"reflect"
	"testing"
)

func TestGetSliceWithoutNesting(t *testing.T) {
	type getSliceWithoutNestingTest struct {
		originalSlice, expectedSlice []interface{}
	}

	var getSliceWithoutNestingTests = []getSliceWithoutNestingTest{
		{
			[]interface{}{1, 2, 3, []interface{}{4}}, []interface{}{1, 2, 3},
		},
		{
			[]interface{}{1, 2, 3, map[string]string{"any": "value"}, 4}, []interface{}{1, 2, 3, 4},
		},
	}

	for _, elmnt := range getSliceWithoutNestingTests {

		elmnt.originalSlice = getSliceWithoutNesting(elmnt.originalSlice) // convert the original slice to the expected slice

		if reflect.DeepEqual(elmnt.originalSlice, elmnt.expectedSlice) == false {
			t.Errorf("got %v, wanted %v", elmnt.originalSlice, elmnt.expectedSlice)
		}
	}
}
