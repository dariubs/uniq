package uniq

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
