package bfs

// https://www.geeksforgeeks.org/minimum-jumps-to-same-value-or-adjacent-to-reach-end-of-array/
func minimumJumps(nums []int) int {
	/*
		[100,  -23,  -23,  404,  100,  23,  23,  23,  3,  404]
	*/

	mapper := make(map[int][]int)
	visited := make([]bool, len(nums))

	for i := 0; i < len(nums); i++ {
		mapper[nums[i]] = append(mapper[nums[i]], i)
	}

	// Count holds the count of jumps it is doing
	count := 0
	// queue holds the current index while doing BFS
	queue := []int{}

	// Appends the first element to queue and mark visited
	queue = append(queue, 0)
	visited[0] = true

	for len(queue) > 0 {
		// Take the current size of queue.
		size := len(queue)

		for i := 0; i < size; i++ {
			// Retrieve from queue
			curr := queue[0]
			queue = queue[1:]

			if curr == len(nums)-1 {
				return count
			}

			if curr+1 < len(nums) && !visited[curr+1] {
				queue = append(queue, curr+1)
				visited[curr+1] = true
			}

			if curr-1 >= 0 && !visited[curr-1] {
				queue = append(queue, curr-1)
				visited[curr-1] = true
			}

			// Finds the similar ones from our mapper
			if indexSlice, ok := mapper[nums[curr]]; ok {
				for k := 0; k < len(indexSlice); k++ {
					similarCurr := indexSlice[k]
					if similarCurr == curr {
						continue
					}

					if !visited[similarCurr] {
						queue = append(queue, similarCurr)
						visited[similarCurr] = true
					}
				}
			}

			// Delete from our map because we assumed that
			// all same value neighbours has been inserted
			// to the queue
			delete(mapper, nums[curr])
		}

		count++
	}

	return count
}
