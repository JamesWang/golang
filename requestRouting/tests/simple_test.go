package tests

import "testing"

func TestSum(t *testing.T) {
	testValues := []int{10, 20, 30}
	_, sum := SortAndTotal(testValues)
	expected := 60

	if sum != expected {
		t.Fatalf("Expected %v, Got %v", expected, sum)
	}
}
