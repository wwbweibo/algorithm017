func levelOrder(root *Node) [][]int {
	nextLevel := make([]*Node, 0)
	result := make([][]int, 0)
	if root == nil {
		return [][]int{}
	}
	currLevel := []*Node{root}
	for len(currLevel) > 0 {
		levelResult := make([]int, 0)
		for _, node := range currLevel {
			levelResult = append(levelResult, node.Val)
			for _, nNode := range node.Children {
				nextLevel = append(nextLevel, nNode)
			}
		}
		result = append(result, levelResult)
		currLevel = nextLevel
		nextLevel = make([]*Node, 0)
	}

	return result
}