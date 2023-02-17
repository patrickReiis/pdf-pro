package handler

import (
	"reflect"
	"sort"
	"strings"
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

/*
Table Driven test for the `getFormattedString` function
The function is called and the result is split into new lines
and then the resulting slice is sorted in alphabetical order.

This is done because maps in go are unordered so this workaround allows to test the `getFormattedString` function
*/
func TestGetFormattedString(t *testing.T) {
	/*
		`arg1` stand for the argument 1 and so on
		`expected` is the result of calling the function `getFormattedString` and splitting it in new lines and sorting the result slice in alphabetical order
	*/
	type getFormattedStringTest struct {
		arg1       map[string]interface{}
		arg2, arg3 int
		expected   []string
	}

	var getFormattedStringTests = []getFormattedStringTest{
		{
			map[string]interface{}{
				"number1": 1,
				"number2": 2,
				"parent": map[string]interface{}{
					"key": "value",
				},
				"parentSlice": []int{1, 2, 3},
			},
			4, 3,
			[]string{"", "       key = value", "   number1 = 1", "   number2 = 2", "   parent:", "   parentSlice = [1 2 3]"},
		},
	}

	for _, elmnt := range getFormattedStringTests {
		result := getFormattedString(elmnt.arg1, elmnt.arg2, elmnt.arg3) // full string result

		got := strings.Split(result, "\n") // splitting the string result in new lines

		sort.Strings(got) // sorting the slice of strings

		if reflect.DeepEqual(got, elmnt.expected) == false {
			t.Errorf("got %q, wanted %q", got, elmnt.expected)
		}
	}
}
