//根据一棵树的前序遍历与中序遍历构造二叉树。 
//
// 注意: 
//你可以假设树中没有重复的元素。 
//
// 例如，给出 
//
// 前序遍历 preorder = [3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7] 
//
// 返回如下的二叉树： 
//
//     3
//   / \
//  9  20
//    /  \
//   15   7 
// Related Topics 树 深度优先搜索 数组 
// 👍 718 👎 0


package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := TreeNode{
		Val: preorder[0],
	}

	rootPos := 0

	for i := 0; i < len(inorder); i ++ {
		if inorder[i] == preorder[0] {
			rootPos	= i
			break
		}
	}
	root.Left = buildTree(preorder[1 :1 + len(inorder[:rootPos])], inorder[:rootPos])
	root.Right = buildTree(preorder[1 + len(inorder[:rootPos]):] , inorder[rootPos+1:])
	return &root
}
//leetcode submit region end(Prohibit modification and deletion)
