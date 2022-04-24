package offerii0004

func singleNumber(nums []int) int {
    var ans int32
    for i:=0;i<32;i++{
        var total int32
        for _, v := range nums {
            total += int32((v>>i)&1)
        }
        if total%3==1 {
            ans+=1<<i
        }
    }
    return int(ans)
}
