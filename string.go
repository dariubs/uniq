package uniq

func UniqString(list []string) []string {
	check := make(map[string]bool)
	var uinque []string

	for _, val := range list {
		if _, ok := check[val]; !ok {
			check[val] = true
			uinque = append(uinque, val)
		}
	}

	return uinque
}
