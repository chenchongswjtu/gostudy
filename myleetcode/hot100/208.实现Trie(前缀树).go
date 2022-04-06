package main

// 208.实现前缀树

type Trie struct {
	Next  [26]*Trie
	IsEnd bool
}

func Constructor1() Trie {
	trie := Trie{}
	trie.Next = [26]*Trie{} //初始化，next数组都为nil
	return trie
}

func (this *Trie) Insert(word string) {
	i := 0
	for ; i < len(word); i++ { // 先生查找已有的，找到最后
		if this.Next[word[i]-'a'] == nil {
			break
		}
		this = this.Next[word[i]-'a'] // 向下走
	}

	for ; i < len(word); i++ { // 没有找完，插入数据
		this.Next[word[i]-'a'] = &Trie{}
		this = this.Next[word[i]-'a'] // 向下走
	}
	this.IsEnd = true // 设置最后的单词结束
}

func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return this.IsEnd // 是否结尾
	}
	if this == nil || this.Next[word[0]-'a'] == nil {
		return false
	}
	return this.Next[word[0]-'a'].Search(word[1:])
}

func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	if this == nil || this.Next[prefix[0]-'a'] == nil {
		return false
	}
	return this.Next[prefix[0]-'a'].StartsWith(prefix[1:])
}
