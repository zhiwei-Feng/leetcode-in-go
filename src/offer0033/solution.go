package offer0033

func verifyPostorder(postorder []int) bool {
	n := len(postorder)
	if n <= 1 {
		return true
	}
	ind := n - 2
	for ind >= 0 && postorder[ind] > postorder[n-1] {
		ind--
	}
	// check left subtree
	for i := ind; i >= 0; i-- {
		if postorder[i] > postorder[n-1] {
			return false
		}
	}

	return verifyPostorder(postorder[:ind+1]) && verifyPostorder(postorder[ind+1:n-1])
}
