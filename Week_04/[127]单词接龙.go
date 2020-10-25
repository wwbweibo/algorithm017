//给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：
// 
//
// 
// 每次转换只能改变一个字母。 
// 转换过程中的中间单词必须是字典中的单词。 
// 
//
// 说明: 
//
// 
// 如果不存在这样的转换序列，返回 0。 
// 所有单词具有相同的长度。 
// 所有单词只由小写字母组成。 
// 字典中不存在重复的单词。 
// 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。 
// 
//
// 示例 1: 
//
// 输入:
//beginWord = "hit",
//endWord = "cog",
//wordList = ["hot","dot","dog","lot","log","cog"]
//
//输出: 5
//
//解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
//     返回它的长度 5。
// 
//
// 示例 2: 
//
// 输入:
//beginWord = "hit"
//endWord = "cog"
//wordList = ["hot","dot","dog","lot","log"]
//
//输出: 0
//
//解释: endWord "cog" 不在字典中，所以无法进行转换。 
// Related Topics 广度优先搜索 
// 👍 478 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == "" || endWord == "" || len(wordList) <= 0 {
		return 0
	}

	wordLength := len(beginWord)
	combines := make(map[string][]string)

	for _, word := range wordList {
		for i := 0; i < wordLength; i ++ {
			newWord := word[:i] + "*"+word[i + 1:]

			if wordArray, ok := combines[newWord]; ok {
				wordArray = append(wordArray, word)
				combines[newWord] = wordArray
			}else {
				wordArray = make([]string, 1)
				wordArray[0] = word
				combines[newWord] = wordArray
			}
		}
	}

	queue := make([]wordPair, 0)
	queue = append(queue, wordPair{beginWord, 1})
	visited := make(map[string]bool)
	visited[beginWord] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for i := 0; i < wordLength;i ++ {
			newWord := node.word[:i] + "*" + node.word[i+1:]
			for _, adjWord := range combines[newWord] {
				if adjWord == endWord {
					return node.idx + 1
				}
				if !visited[adjWord] {
					visited[adjWord] = true
					queue = append(queue, wordPair{adjWord, node.idx + 1})
				}
			}
		}

	}

	return 0
}

type wordPair struct {
	word string
	idx int
}
//leetcode submit region end(Prohibit modification and deletion)
