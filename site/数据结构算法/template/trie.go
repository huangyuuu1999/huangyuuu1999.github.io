package main

import "fmt"

type trieNode struct {
	end  int
	pass int
	son  [26]*trieNode
}

type trie struct{ root *trieNode }

func newTrie() *trie {
	return &trie{&trieNode{}} // 三个字段、自动赋值0值
}

func (trie) ord(c rune) rune { return c - 'a' }
func (trie) chr(v byte) byte { return v + 'a' }

func (t *trie) put(s string) *trieNode { // 插入字符串到字典树
	o := t.root
	for _, b := range s { // b type rune
		b = t.ord(b)
		if o.son[b] == nil {
			o.son[b] = &trieNode{} // 创建新节点
		}
		o = o.son[b]
		o.pass++
	}
	o.end++
	return o
}

func (t *trie) countPath(o *trieNode) (ans int, first_index int) {
	for i := 0; i < 26; i++ {
		if o.son[i] != nil {
			if ans == 0 {
				first_index = i
			}
			ans += 1
		}
	}
	return
}

// 查询是否有字符串s
func (t *trie) search(s string) bool {
	o := t.root
	for _, b := range s {
		if o.son[t.ord(b)] == nil {
			return false
		}
		o = o.son[t.ord(b)]
	}
	return o.end > 0
}

// 返回整棵树所有字符串的 最长公共前缀
func (t *trie) longestSharedPrefix() (o *trieNode, s string) {
	o = t.root
	for {
		if count, target := t.countPath(o); count == 1 {
			fmt.Printf("target: %v\n", target)
			s += string(t.chr(byte(target)))
			fmt.Printf("s: %v\n", s)
			o = o.son[target]
			if o.end > 0 {
				break
			}
		} else {
			break
		}
	}
	return o, s
}

func main() {
	tree := newTrie()
	tree.put("flower")
	tree.put("flow")
	tree.put("flowk")
	_, s := tree.longestSharedPrefix()
	fmt.Println(s)
	ans := tree.search("flower")
	fmt.Printf("ans: %v\n", ans)
	ans = tree.search("flow")
	fmt.Printf("ans: %v\n", ans)
	ans = tree.search("fliht")
	fmt.Printf("ans: %v\n", ans)
}
