package uniq

import "testing"

func TestInt(t *testing.T) {
	list := []int{1, 2, 2, 3, 4, 5, 5, 5, 6, 6, 7}
	unique := Int(list)
	want := []int{1, 2, 3, 4, 5, 6, 7}

	for i, v := range unique {
		if v != want[i] {
			t.Fatalf("item not found in unique list")
		}
	}
}
