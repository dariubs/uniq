package uniq

// UniqUint make uint slice items unique
// It returns an slice of uint items
func UniqUint(list []uint) []uint {
	check := make(map[uint]bool)
	var uinque []uint

	for _, val := range list {
		if _, ok := check[val]; !ok {
			check[val] = true
			uinque = append(uinque, val)
		}
	}

	return uinque
}
