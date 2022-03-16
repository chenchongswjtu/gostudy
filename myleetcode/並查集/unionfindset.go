package main

type UnionFindSet struct {
	Parent []int // 人员及其Header数组
	N      int   // 一共有多少人
}

func NewUnionFindSet(n int) *UnionFindSet {
	parent := make([]int, n)
	// 让每一个人的父亲指向自己
	for i := range parent {
		parent[i] = i
	}

	return &UnionFindSet{Parent: parent, N: n}
}

// Find查找根节点
func (u *UnionFindSet) Find(x int) int {
	//if u.Parent[x] == x {
	//	return x
	//} else {
	//	// 如果他不是根节点，接着往上面找根节点，并把根节点赋给当前元素的父节点，构造二层的平铺树
	//	// 缩短查找距离
	//	u.Parent[x] = u.Find(u.Parent[x])
	//	return u.Parent[x]
	//}

	for u.Parent[x] != x {
		// 进行路径压缩
		u.Parent[x] = u.Parent[u.Parent[x]]
		x = u.Parent[x]
	}
	return x
}

// Union合并两个节点到同一个联通域
func (u *UnionFindSet) Union(x, y int) {
	fx := u.Find(x)
	fy := u.Find(y)

	// 已经联通
	if fx == fy {
		return
	}

	u.Parent[fx] = fy
}
