package main

import (
	"time"
)

// 区块链的头部信息
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	// 用于pow的计算
	Nonce         int
}

// 创建新的区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// 创建"创世区块"
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}