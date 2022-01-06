package n0071

import "strings"

func simplifyPath(path string) string {
    dirs := strings.Split(path, "/")
    ans := []string{}
    for _, dir := range dirs {
        if dir == "" || dir == "." {
            continue
        }else if dir == ".." {
            if len(ans)!=0 {
                ans = ans[:len(ans)-1]
            }
        }else {
            ans = append(ans, dir)
        }
    }

    var builder strings.Builder
    builder.WriteString("/")
    for i, v := range ans {
        builder.WriteString(v)
        if i!=len(ans)-1 {
            builder.WriteString("/")
        } 
    }
    return builder.String()
}