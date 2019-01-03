import "strconv"

/**
Runtime: 0 ms, faster than 100.00% of Go online submissions for Leaf-Similar Trees


*/

func getLeaf(root *TreeNode, res *string) {
	if root == nil {
		return
	}
	if (*root).Left == nil && (*root).Right == nil {
		*res += strconv.Itoa((*root).Val)
		return
	}
	getLeaf(root.Left, res)
	getLeaf(root.Right, res)

}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	res1, res2 := "", ""
	getLeaf(root1, &res1)
	getLeaf(root2, &res2)

	if res1 == res2 {
		return true
	}
	return false

}