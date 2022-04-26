package main

// 最大公约数
func gcd(a, b int) int {
	if a < b {
		a, b = b, a // 交换，保证a>=b
	}

	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

// 最小公倍数
func lcm(a, b int) int {
	return a / gcd(a, b) * b //先计算商，防止溢出
}
