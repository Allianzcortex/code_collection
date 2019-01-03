/**

A binary tree is univalued if every node in the tree has the same value.

Return true if and only if the given tree is univalued.

* Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 }

**/
func traverseWithValue(root *TreeNode, value int) bool {
	if root == nil {
		return true
	}
	if (*root).Val != value {
		return false
	}
	target := (*root).Val
	return traverseWithValue((*root).Left, target) && traverseWithValue((*root).Right, target)

}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return traverseWithValue(root, (*root).Val)

}