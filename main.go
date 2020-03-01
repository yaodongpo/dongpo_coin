package main

import (
	"fmt"
)

//区块结构体
type Block struct {
	Data, Hash, Prehash string
}

//区块构造函数
func NewBlock(data, prehash string) *Block {
	hash := ComputeHash(data)
	return &Block{data, hash, prehash}
}

//计算区块的hash
func ComputeHash(data string) string {
	return GetSHA256HashCode([]byte(data))
}

//测试main函数
func main() {
	block0 := NewBlock("test Newblock", "")
	block1 := NewBlock("test block1", "683e34a61c3867bc4657a00079de786df6aa6521d7433ebbdaf34635437a20f6")
	fmt.Println("block0:", block0)
	fmt.Println("block1:", block1)
}
