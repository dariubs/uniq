package uniq

func Int(list []int) []int {
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
