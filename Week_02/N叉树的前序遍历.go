func preorder(root *Node) []int {
	var result = []int{}
	var fun func(node *Node)
	fun = func(node *Node) {
		if node != nil {
			result = append(result, node.Val)
			for _, x := range node.Children {
				fun(x)
			}
		}
	}
	fun(root)
	return result
}