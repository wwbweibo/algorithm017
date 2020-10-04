func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	preorder(root, &result)
	return result
}

func preorder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	*result = append(*result, node.Val)
	preorder(node.Left, result)
	preorder(node.Right, result)
}