package main

func main() {

}

type unionFind struct {
	// 记录连通分量
	count int
	// 存储若干个树
	parent []int
	// 记录树的重量
	size []int
}

func new(n int) *unionFind {
	uf := unionFind{
		count:  n,
		parent: make([]int, n),
		size:   make([]int, n),
	}

	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}

	return &uf
}

// 将p和q连通
func (uf *unionFind) union(p, q int) {
	rootP := uf.find(p)
	rootQ := uf.find(q)
	if rootP == rootQ {
		return
	}

	// 小树接到大树下面，较平衡
	if uf.size[rootP] > uf.size[rootQ] {
		uf.parent[rootQ] = rootP
		uf.size[rootP] += uf.size[rootQ]
	} else {
		uf.parent[rootP] = rootQ
		uf.size[rootQ] += uf.size[rootP]
	}
	uf.count--
}

// 返回节点x的根节点
func (uf *unionFind) find(x int) int {
	for uf.parent[x] != x {
		// 进行路径压缩
		uf.parent[x] = uf.parent[uf.parent[x]]
		x = uf.parent[x]
	}
	return x
}

// 判断p和q是否互相连通
func (uf *unionFind) connected(p, q int) bool {
	rootP := uf.find(p)
	rootQ := uf.find(q)
	return rootP == rootQ
}
