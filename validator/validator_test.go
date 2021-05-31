package validator

import (
	"testing"
)

func TestIsNumeric(t *testing.T) {
	//list_numbers_true := [...]float64{-999.63, -100, -1, 0, 1, 10, 1000, 999, 966.6}
	list_string_numbers_true := [...]string{"-99.6", "-1", "0", "1", "1000", "9999.99"}
	list_numbers_false := [...]string{"a1", "one", "two", "1a9", "0a1", "@", "#", "$", "?", "&", "`", ";", ","}

	for i := range list_string_numbers_true {
		number := list_string_numbers_true[i]
		msg := IsNumeric(number)
		if msg == false {
			t.Fatalf(`IsNumeric("%v") = must return true, but the return was false`, number)
		}
	}
	for i := range list_numbers_false {
		number := list_numbers_false[i]
		msg := IsNumeric(number)
		if msg == true {
			t.Fatalf(`Isnumber("%v") = must return false, but the return was true`, number)
		}
	}

}
