package n0440

func findKthNumber(n int, k int) int {
    cur := 1
    k--
    for k > 0 {
        step := getSteps(cur, n)
        if step<=k {
            k-=step
            cur++
        }else {
            cur = cur*10
            k--
        }
    }
    return cur
}

func getSteps(cur, n int) int {
    last := cur
    count := 0
    for cur<=n {
        count += min(last, n)-cur+1
        cur, last = cur*10, last*10+9
    }
    return count
}

func min(i, j int) int {
    if i < j {
        return i
    }
    return j
}