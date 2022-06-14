package n0220

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	mp := make(map[int]int)
	for i, x := range nums {
		id := getBucketId(x, t+1)
		if _, has := mp[id]; has {
			return true
		}
		if y, has := mp[id-1]; has && abs(x-y) <= t {
			return true
		}
		if y, has := mp[id+1]; has && abs(x-y) <= t {
			return true
		}
		mp[id] = x
		if i >= k {
			delete(mp, getBucketId(nums[i-k], t+1))
		}
	}
	return false
}

func abs(i int) int {
	if i < 0 {
		i = -i
	}
	return i
}

// x: value, w: bucket size
// Especially, if w = 3, [-1, -2, -3] belong to id[-1] bucket
func getBucketId(x, w int) int {
	if x >= 0 {
		return x / w
	}
	return (x+1)/w - 1
}
