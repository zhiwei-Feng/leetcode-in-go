package n0101

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
    var dfs func(p, q *TreeNode) bool
    dfs = func(p, q *TreeNode) bool {
        if p==nil||q==nil {
            return p==nil&&q==nil
        }
        if p.Val!=q.Val{
            return false
        }
        judge1 := dfs(p.Left, q.Right)
        judge2 := dfs(p.Right, q.Left)
        return judge1&&judge2
    }
    return dfs(root, root)
}

func isSymmetric_(root *TreeNode) bool {
    u, v := root, root
    q := []*TreeNode{}
    q = append(q, u)
    q = append(q, v)
    for len(q) > 0 {
        u, v = q[0], q[1]
        q = q[2:]
        if u == nil && v == nil {
            continue
        }
        if u == nil || v == nil {
            return false
        }
        if u.Val != v.Val {
            return false
        }
        q = append(q, u.Left)
        q = append(q, v.Right)

        q = append(q, u.Right)
        q = append(q, v.Left)
    }
    return true
}