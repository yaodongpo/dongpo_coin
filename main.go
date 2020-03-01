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

//链，区块的单链
type Chain []Block

//生成区块链的时候初始化创世区块
func NewChain() *Chain {
	block0 := NewBlock("I am the fuck 创世区块", "")
	return &Chain{*block0}
}

//测试main函数
func main() {
	chain := NewChain()
	fmt.Println("chain:", chain)
}
