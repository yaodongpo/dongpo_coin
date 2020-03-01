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
type Chain struct {
	chain []Block
}

//生成区块链的时候初始化创世区块
func NewChain() *Chain {
	block0 := NewBlock("I am the fuck 创世区块", "") //创世区块没有前区块
	chain := []Block{*block0}
	return &Chain{chain}
}

//根据下标获取区块
func (chain Chain) GetIndex(index int8) Block {
	return chain.chain[index]
}

//新增区块到链上
func (chain *Chain) AddBlock(block Block) error {
	chain.chain = append(chain.chain, block)
	return nil
}

//测试main函数
func main() {
	chain := NewChain()
	fmt.Println("chain:", chain)
	fmt.Println("chain[0]:", chain.GetIndex(0))
	block1 := NewBlock("测试新增区块到链上！", chain.GetIndex(0).Hash)
	chain.AddBlock(*block1)
	fmt.Println("chain 增加区块后:", chain)

}
