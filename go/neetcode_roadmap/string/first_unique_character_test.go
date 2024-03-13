package string

import "testing"

func firstUniqueCharacter(s string) string {
	if len(s) <= 1 {
		return s
	}

	hash := map[byte][]int{}
	for i := 0; i < len(s); i++ {
		if len(hash[s[i]]) <= 1 {
			hash[s[i]] = append(hash[s[i]], i)
		}
	}

	min := len(s)

	for key := range hash {
		if len(hash[key]) <= 1 {
			if min > hash[key][0] {
				min = hash[key][0]
			}
		}
	}

	return string(s[min])
}

func TestFirstUniqueCharacter(t *testing.T) {
	if result := firstUniqueCharacter("jakarta"); result != "j" {
		t.Fatalf("expected j but got %s", result)
	}

	if result := firstUniqueCharacter("karta"); result != "k" {
		t.Fatalf("expected k but got %s", result)
	}

	if result := firstUniqueCharacter("motherr"); result != "m" {
		t.Fatalf("expected m but got %s", result)
	}

	if result := firstUniqueCharacter("otter"); result != "o" {
		t.Fatalf("expected o but got %s", result)
	}

	if result := firstUniqueCharacter("abaccdbahakhhewbhakewbiv"); result != "d" {
		t.Fatalf("expected d but got %s", result)
	}
}
