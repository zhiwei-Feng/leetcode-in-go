package n0011

func maxArea(height []int) int {
    i, j:=0, len(height)-1
    ans := 0
    for i<j {
        ans = max(ans, (j-i)*min(height[i],height[j]))
        if height[i]<height[j] {
            i++
        }else {
            j--
        }
    }
    return ans
}

func max(i, j int) int {
    if i > j {
        return i
    }
    return j    
}

func min(i, j int) int {
    if i < j {
        return i
    }
    return j
}