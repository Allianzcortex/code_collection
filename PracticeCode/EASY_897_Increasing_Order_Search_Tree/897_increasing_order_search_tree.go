import (
	"fmt"
	"strconv"
	"strings"
)

/**
Runtime: 48 ms, faster than 31.03% of Go online submissions for Increasing Order Search Tree.

Given a tree, rearrange the tree in in-order so that the leftmost node in the tree is now the root of the tree, and every node has no left child and only 1 right child.

Example 1:
Input: [5,3,6,2,4,null,8,1,null,null,null,7,9]

       5
      / \
    3    6
   / \    \
  2   4    8
 /        / \
1        7   9

Output: [1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]

 1
  \
   2
    \
     3
      \
       4
        \
         5
          \
           6
            \
             7
              \
               8
                \
                 9
Note:

The number of nodes in the given tree will be between 1 and 100.
Each node will have a unique integer value from 0 to 1000.


**/
func traverseInOrder(root *TreeNode, res *string) {
	if root == nil {
		return
	}
	traverseInOrder(root.Left, res)
	*res += strconv.Itoa((*root).Val)
	*res += ","
	traverseInOrder(root.Right, res)

}

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	res := ""
	traverseInOrder(root, &res)
	ss := strings.Split(res, ",")
	// 如果不这么做的话 split 之后还不会又多出一个 0？？我不知道是什么情况
	ss = ss[:len(ss)-1]
	fmt.Println(ss)
	length := len(ss)
	// var set [length]TreeNode
	set := make([]TreeNode, length)

	s := string(ss[length-1])
	fmt.Println(strconv.Atoi(s))
	set[length-1].Val, _ = strconv.Atoi(s)
	for i := 0; i < length-1; i++ {
		set[i].Val, _ = strconv.Atoi(string(ss[i]))
		set[i].Right = &set[i+1]
	}

	return &set[0]

}