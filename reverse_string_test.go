package reverse_string

import "testing"

func TestReverseString(t *testing.T) {
	initial := "abcdefghאב"
	actual := ReverseString(initial)
	expected := "באhgfedcba"
	if actual != expected {
		t.Errorf("Result is not correct, expected:%s, but was:%s", expected, actual)
	}
}
