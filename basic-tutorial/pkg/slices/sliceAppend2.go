package main

import "fmt"

// 获取slice的长度和容量
func dump1(label string, slice []string) {
	fmt.Println("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}
func main() {
	dwarfs1 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfs2 := append(dwarfs1, "Orcus")                      //goland中扩容两倍到10
	dwarfs3 := append(dwarfs2, "Salacia", "Quaoar", "Sedna") // 容量已经是10,此处无需扩容

	dump1("dwarfs1", dwarfs1)
	dump1("dwarfs2", dwarfs2)
	dump1("dwarfs3", dwarfs3)
}
