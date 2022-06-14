package n0653

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	var (
		l, r       *TreeNode
		lstk, rstk []*TreeNode
	)
	// init
	l, r = root, root
	lstk = append(lstk, l)
	rstk = append(rstk, r)
	for l.Left != nil {
		lstk = append(lstk, l.Left)
		l = l.Left
	}
	for r.Right != nil {
		rstk = append(rstk, r.Right)
		r = r.Right
	}

	// main process
	for l != r {
		sum := l.Val + r.Val
		if sum == k {
			return true
		}

		if sum > k {
			r, rstk = rstk[len(rstk)-1], rstk[:len(rstk)-1]
			for cur := r.Left; cur != nil; cur = cur.Right {
				rstk = append(rstk, cur)
			}
		} else {
			l, lstk = lstk[len(lstk)-1], lstk[:len(lstk)-1]
			for cur := l.Right; cur != nil; cur = cur.Left {
				lstk = append(lstk, cur)
			}
		}
	}
	return false
}

/* traversal tree to store inorder array
func findTarget(root *TreeNode, k int) bool {
    var nums = make([]int, 0, 1e4)
    p := root
    var stk []*TreeNode
    for p != nil || len(stk) > 0 {
        for p != nil {
            stk = append(stk, p)
            p = p.Left
        }

        p, stk = stk[len(stk)-1], stk[:len(stk)-1]
        nums = append(nums, p.Val)
        p = p.Right
    }

    i, j := 0, len(nums)-1
    for i < j {
        if nums[i]+nums[j] == k {
            return true
        } else if nums[i]+nums[j] > k {
            j--
        } else {
            i++
        }
    }
    return false
}
*/
