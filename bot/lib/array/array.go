package array

func Contains[K int | string](arr []K, value K) bool {
	for _, e := range arr {
		if e == value {
			return true
		}
	}

	return false
}
