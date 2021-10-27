package uniq

import "testing"

func TestUniqString(t *testing.T) {
	list := []string{"hello", "hello", "holla", "hello", "hi", "holla", "hi"}
	unique := UniqString(list)
	want := []string{"hello", "holla", "hi"}

	for i, v := range unique {
		if v != want[i] {
			t.Fatalf("item not found in unique list")
		}
	}
}
