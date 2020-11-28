package sort

import "testing"

// TestInts test the sort.Ints func
func TestInts(t *testing.T) {
	a := []int{3, 4, 5, 1, -1, 5, 3}
	Ints(a)
	t.Log(a)
}

func TestStrings(t *testing.T) {
	a := []string{"b", "c", "a", "z", "zz", "ff", "bb", "ac", "adf"}
	Strings(a)
	t.Log(a)
}
