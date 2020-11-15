//实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。 
//
// 示例: 
//
// Trie trie = new Trie();
//
//trie.insert("apple");
//trie.search("apple");   // 返回 true
//trie.search("app");     // 返回 false
//trie.startsWith("app"); // 返回 true
//trie.insert("app");   
//trie.search("app");     // 返回 true 
//
// 说明: 
//
// 
// 你可以假设所有的输入都是由小写字母 a-z 构成的。 
// 保证所有输入均为非空字符串。 
// 
// Related Topics 设计 字典树 
// 👍 461 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
type Trie struct {
	Val      rune
	Leaf     bool
	Children map[rune]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	child := make(map[rune]*Trie)
	return Trie{Val: 0, Children: child}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	ptr := this
	for _, c := range word {
		_, exists := ptr.Children[c]
		if !exists {
			child := make(map[rune]*Trie)
			ptr.Children[c] = &Trie{Val: c, Children: child}
		}
		ptr = ptr.Children[c]
	}
	ptr.Leaf = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.search(word)
	return node != nil && node.Leaf
}

/** Returns if the word is in the trie. */
func (this *Trie) search(word string) *Trie {
	ptr := this
	for _, c := range word {
		ct, ok := ptr.Children[c]
		if !ok {
			return nil
		}
		ptr = ct
	}
	return ptr
}

func (this *Trie) Delete(word string) {
	var parents []*Trie

	ptr := this
	for _, c := range word {
		parents = append(parents, ptr)
		ct, ok := ptr.Children[c]
		if !ok {
			return
		}
		ptr = ct
	}

	if !ptr.Leaf {
		return
	}

	if len(ptr.Children) != 0 {
		ptr.Leaf = false
		return
	}

	for len(parents) > 0 {
		p := parents[len(parents)-1]
		parents = parents[:len(parents)-1]

		delete(p.Children, ptr.Val)
		if len(p.Children) != 0 || p.Leaf {
			break
		}
		ptr = p
	}
}

func (this *Trie) Walk() {
	var walk func(string, *Trie)
	walk = func(pfx string, node *Trie) {
		if node == nil {
			return
		}

		if node.Val != 0 {
			pfx += string(node.Val)
		}

		if node.Leaf {
			fmt.Println(string(pfx))
		}

		for _, v := range node.Children {
			walk(pfx, v)
		}
	}
	walk("", this)
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.search(prefix)
	return node != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
//leetcode submit region end(Prohibit modification and deletion)
