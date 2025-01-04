package integers

import (
	"fmt"
	"testing"
)

// NOTE: Like test functions, example functions also reside inside _test.go files
// and they will show up as clickable examples in documentation at pkgsite

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
