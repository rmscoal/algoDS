package string

func findMin(x, y int32) int32 {
	if x < y {
		return x
	} else {
		return y
	}
}

func makeAnagrams(a, b string) int32 {
	// Given a and b. We can create two hash maps
	// that stores the number for each of the letters.

	// We can iterate all of these and the get the minimum
	// for each of them. Then we can substract the value of
	// each amount with the minimum.

	// Example, a = "cde" and b = "dcf"
	// we will have:
	// hashMapA = {c: 1, d: 1, e: 1}
	// hashMapB = {d: 1, c: 1, f: 1}
	// We are going to iterate hashMapA, for key, value := range hashMapA
	// min := minFunction(hashMapA[key], hashMapB[key])
	// in the cases where hashMapA's key is not in hashMapB's key.. Then
	// we substract those value from hashMapA[key]

	// We can do another loop to hashMapB... Here we are checking.. If there
	// are some keys in hashMapB that are not in A. If that is present

	A := make(map[rune]int32)
	B := make(map[rune]int32)

	var count int32 = 0

	for _, char := range a {
		A[char]++
	}

	// Can do last step here also actually
	for _, char := range b {
		if _, found := A[char]; !found {
			count++
		} else {
			B[char]++
		}
	}

	for char, value := range A {
		if _, found := B[char]; !found {
			count += value
			continue
		}

		min := findMin(A[char], B[char])
		A[char] -= min
		B[char] -= min
		count += A[char] + B[char]
	}

	return count
}
