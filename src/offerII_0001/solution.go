package offerii0001

import "math"

func divide(a int, b int) int {
    // special case
    if a == 0 {
        return 0
    }
    if a == math.MinInt32 && b == -1 {
        return math.MaxInt32
    }
    if b == -1 {
        return -a
    }else if b == 1 {
        return a
    }

    flag := false
    if a > 0 {
        a = -a
        flag = !flag
    }
    if b > 0 {
        b = -b
        flag = !flag
    }

    if a>b {
        return 0
    }

    l, r := 1, math.MaxInt32
    for l < r {
        mid := l + (r-l+1)/2
        if fastAdd(b, mid, a) {
            // fmt.Println(b, mid, a)
            l = mid
        }else {
            r = mid-1
        }
    }
    if flag {
        l = -l
    }
    return l
}

// base*n >= t 因为base和t都是负数, 所以用>=
func fastAdd(base, n, t int) bool {
    add := base
    ans := 0
    for n>0 {
        if n&1 == 1 {
            ans += add
        }
        add += add
        n >>= 1
    }
    return ans >= t
}