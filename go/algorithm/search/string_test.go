package search

import "testing"

func BenchmarkAppendUnique(b *testing.B) {
	b.Run("Hashmap", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			appendUniqueHashmap(map[string]bool{
				"test@email.com":  true,
				"test1@email.com": true,
				"test2@email.com": true,
				"test3@email.com": false,
				"test4@email.com": true,
				"test5@email.com": false,
			}, "x")
		}
	})

	b.Run("Linear Array", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			appendUniqueArray([]string{
				"test@email.com", "test1@email.com", "test2@email.com",
				"test3@email.com", "test4@email.com", "test5@email.com",
			}, "x")
		}
	})
}
