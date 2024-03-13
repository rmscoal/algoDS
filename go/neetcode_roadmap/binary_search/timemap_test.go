package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Solution to: https://leetcode.com/problems/time-based-key-value-store/description/
// IMPORTANT NOTE: All the timestamps timestamp of set are strictly increasing.

type TimeMap struct {
	// NOTE: Can use this, for faster Set
	// hashmap map[string]struct {
	// 	val string
	// 	timestamp int
	// }
	hashmap_time map[string][]int
	hashmap_val  map[string][]string
}

func Constructor() TimeMap {
	return TimeMap{
		hashmap_time: make(map[string][]int),
		hashmap_val:  make(map[string][]string),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.hashmap_time[key]; !ok {
		this.hashmap_time[key] = make([]int, 0)
		this.hashmap_val[key] = make([]string, 0)
	}
	this.hashmap_time[key] = append(this.hashmap_time[key], timestamp)
	this.hashmap_val[key] = append(this.hashmap_val[key], value)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, ok := this.hashmap_time[key]; !ok {
		return ""
	}

	// We can return empty string early when our
	// timestamp is lesser that what was recorded
	if timestamp < this.hashmap_time[key][0] {
		return ""
	}

	var index int
	left, right := 0, len(this.hashmap_time[key])-1
	for left <= right {
		mid := (left + right) / 2

		if this.hashmap_time[key][mid] == timestamp {
			return this.hashmap_val[key][mid]
		} else if this.hashmap_time[key][mid] > timestamp {
			// Move to the left
			right = mid - 1
		} else {
			// Move to the right, and also record our previous mid
			left = mid + 1
			index = mid
		}
	}

	return this.hashmap_val[key][index]
}

func TestTimeMap_Case1(t *testing.T) {
	tm := Constructor()
	tm.Set("foo", "bar", 1)
	assert.Equal(t, "bar", tm.Get("foo", 1))
	assert.Equal(t, "bar", tm.Get("foo", 3))
	tm.Set("foo", "bar2", 4)
	assert.Equal(t, "bar2", tm.Get("foo", 4))
	assert.Equal(t, "bar2", tm.Get("foo", 5))
	assert.Equal(t, "bar", tm.Get("foo", 3))
}

func TestTimeMap_Case2(t *testing.T) {
	// Case
	// 1, 3, 5, 6, 8, 10

	tm := Constructor()
	tm.Set("foo", "bar1", 1)
	tm.Set("foo", "bar3", 3)
	assert.Equal(t, "bar1", tm.Get("foo", 1))
	assert.Equal(t, "bar1", tm.Get("foo", 2))
	assert.Equal(t, "bar3", tm.Get("foo", 3))
	tm.Set("foo", "bar5", 5)
	assert.Equal(t, "bar3", tm.Get("foo", 3))
	assert.Equal(t, "bar3", tm.Get("foo", 4))
	assert.Equal(t, "bar5", tm.Get("foo", 5))
	tm.Set("foo", "bar6", 6)
	assert.Equal(t, "bar6", tm.Get("foo", 6))
	assert.Equal(t, "bar3", tm.Get("foo", 4))
	assert.Equal(t, "bar5", tm.Get("foo", 5))
	tm.Set("foo", "bar8", 8)
	assert.Equal(t, "bar6", tm.Get("foo", 6))
	assert.Equal(t, "bar6", tm.Get("foo", 7))
	assert.Equal(t, "bar8", tm.Get("foo", 8))
	tm.Set("foo", "bar10", 10)
	assert.Equal(t, "bar8", tm.Get("foo", 9))
	assert.Equal(t, "bar10", tm.Get("foo", 10))
	assert.Equal(t, "", tm.Get("foo", 0))
}
