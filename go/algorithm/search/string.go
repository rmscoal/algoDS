package search

func appendUniqueHashmap(hash map[string]bool, x string) map[string]bool {
	if _, ok := hash[x]; ok {
		return hash
	}

	hash[x] = true
	return hash
}

func appendUniqueArray(arr []string, x string) []string {
	for _, v := range arr {
		if v == x {
			return arr
		}
	}

	arr = append(arr, x)
	return arr
}
