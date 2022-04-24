package main

type Node struct {
	val    int
	left   *Node
	right  *Node
	parent *Node
}

//由于节点中包含父节点的指针，所以二叉树的根节点就没必要输入了。
//这道题其实不是公共祖先的问题，而是单链表相交的问题，你把parent指针想象成单链表的next指针，题目就变成了：
//给你输入两个单链表的头结点p和q，这两个单链表必然会相交，请你返回相交点。

func lowestCommonAncestor1(p, q *Node) *Node {
	a := p
	b := q
	for a != b {
		// a走一步，如果走到根节点,转到q节点
		if a == nil {
			a = q
		} else {
			a = a.parent
		}

		// b走一步，如果走到根节点，转到p节点
		if b == nil {
			b = p
		} else {
			b = b.parent
		}
	}

	return a
}
