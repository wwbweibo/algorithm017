//给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。 
//
// 找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。 
//
// 示例: 
//
// X X X X
//X O O X
//X X O X
//X O X X
// 
//
// 运行你的函数后，矩阵变为： 
//
// X X X X
//X X X X
//X X X X
//X O X X
// 
//
// 解释: 
//
// 被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被
//填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。 
// Related Topics 深度优先搜索 广度优先搜索 并查集 
// 👍 422 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func solve(board [][]byte) {
	if len(board) <= 2 {
		return
	}

	m, n := len(board), len(board[0])
	dummy, ufind := m*n, NewUnionFind(m*n+1)

	position := func(i, j int) int {
		return i*n + j
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				if i == 0 || j == 0 || i == m-1 || j == n-1 {
					ufind.Union(position(i, j), dummy)
				}
				if i > 0 && board[i-1][j] == 'O' {
					ufind.Union(position(i, j), position(i-1, j))
				}
				if j > 0 && board[i][j-1] == 'O' {
					ufind.Union(position(i, j), position(i, j-1))
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && !ufind.IsConnected(position(i, j), dummy) {
				board[i][j] = 'X'
			}
		}
	}
}

type UnionFind struct {
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{
		parent: parent,
	}
}

func (u *UnionFind) Union(i, j int) {
	pi := u.Find(i)
	pj := u.Find(j)
	if pi != pj {
		u.parent[pi] = pj
	}
}

func (u *UnionFind) Find(i int) int {
	root := i
	for u.parent[root] != root {
		root = u.parent[root]
	}
	for u.parent[i] != root {
		i, u.parent[i] = u.parent[i], root
	}
	return root
}

func (u *UnionFind) IsConnected(i, j int) bool {
	return u.Find(i) == u.Find(j)
}

//leetcode submit region end(Prohibit modification and deletion)
