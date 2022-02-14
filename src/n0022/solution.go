package n0022

func generateParenthesis(n int) []string {
    posNum := n*2
    ans := make([]string, 0)
    var backtrace func(int)
    list := make([]byte,0)
    listLeftNum := 0

    backtrace = func(pos int) {
        if pos==posNum{
            ans = append(ans, string(list))
            return
        }

        if listLeftNum<n {
            list = append(list, '(')
            listLeftNum++
            backtrace(pos+1)
            list = list[:len(list)-1]
            listLeftNum--
        }
        if 2*listLeftNum>pos{
            list = append(list, ')')
            backtrace(pos+1)
            list = list[:len(list)-1]
        }
    }


    backtrace(0)
    return ans
}
