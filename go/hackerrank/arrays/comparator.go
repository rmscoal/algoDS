package arrays

// Returns if string `a` comes before `b` alphabetically
func compareAlphabet(a, b string) bool {
	min := 1000000000000
	if min >= len(a) {
		min = len(a)
	}

	if min >= len(b) {
		min = len(b)
	}

	for i := 0; i < min; i++ {
		if a[i] < b[i] {
			return true
		} else if a[i] > b[i] {
			return false
		}
	}

	return len(a) <= len(b)
}

func comparator(names []string, scores []int) {
	if len(names) != len(scores) {
		panic("Both lengths are not the same")
	}

	n := len(names)

	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if scores[j] < scores[j+1] {
				swapped = true
				scores[j], scores[j+1] = scores[j+1], scores[j]
				names[j], names[j+1] = names[j+1], names[j]
			} else if scores[j] == scores[j+1] && !compareAlphabet(names[j], names[j+1]) {
				swapped = true
				names[j], names[j+1] = names[j+1], names[j]
				scores[j], scores[j+1] = scores[j+1], scores[j]
			}
		}

		if !swapped {
			break
		}
	}
}

/*
func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    sizeInput := strings.TrimSpace(readLine(reader))
    size, err := strconv.Atoi(sizeInput)
    checkError(err)

    names := make([]string, 0, size)
    scores := make([]int, 0, size)

    for i := 0; i < int(size); i++ {
        input := strings.Split(readLine(reader), " ")
        names = append(names, input[0])

        score, err := strconv.Atoi(input[1])
        checkError(err)

        scores = append(scores, score)
    }

    comparator(names, scores)

    for i := 0; i < int(size); i++ {
        fmt.Printf("%s %d\n", names[i], scores[i])
    }
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
*/
