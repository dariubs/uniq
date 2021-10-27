package uniq

import "testing"

func TestUniqUint(t *testing.T) {
	list := []uint{1, 2, 2, 3, 4, 5, 5, 5, 6, 6, 7}
	unique := UniqUint(list)
	want := []uint{1, 2, 3, 4, 5, 6, 7}

	for i, v := range unique {
		if v != want[i] {
			t.Fatalf("item not found in unique list")
		}
	}
}
