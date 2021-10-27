package uniq

// UniqInt make int slice items unique
// It returns an slice of int items
func UniqInt(list []int) []int {
	check := make(map[int]bool)
	var uinque []int

	for _, val := range list {
		if _, ok := check[val]; !ok {
			check[val] = true
			uinque = append(uinque, val)
		}
	}

	return uinque
}
