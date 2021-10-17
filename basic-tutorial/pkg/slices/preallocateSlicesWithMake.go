package main

// 通过内置的make函数，可以对slice进行预分配策略
// 尽量避免额外的内存分配和数组复制操作

func main() {
	dwarfs := make([]string, 0, 10)
	dump("dwarfs", dwarfs)
	// 追加5个元素
	dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")
	dump("dwarfs", dwarfs)
}
