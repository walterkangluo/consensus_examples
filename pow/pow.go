package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//pow 挖矿算法
//定义难度系数
//const difiiculty = 4
var difiiculty int

type Block struct {
	Index      int // 区块高度
	TimeStamp  int64
	Data       string //交易记录
	Hash       string
	Prehash    string
	Nonce      int
	Difficulty int //难度系数
}

//创建区块链
var BlockChain []Block

//创世区块
func GenesisBlock() *Block {
	var geneBlock = Block{0, time.Now().Unix(), "", "", "", 0, difiiculty}
	geneBlock.Hash = hex.EncodeToString(BlockHash(geneBlock))
	return &geneBlock
}

func BlockHash(block Block) []byte {
	re := strconv.Itoa(block.Index) + strconv.Itoa(int(block.TimeStamp)) + block.Data + block.Prehash +
		strconv.Itoa(block.Nonce) + strconv.Itoa(block.Difficulty)
	h := sha256.New()
	h.Write([]byte(re))
	hashed := h.Sum(nil)
	return hashed
}

func isBlockValid(block Block) bool {
	prefix := strings.Repeat("0", block.Difficulty)
	return strings.HasPrefix(block.Hash, prefix)
}

//创建新区块 pow挖矿
func CreateNewBlock(lastBlock *Block, data string) *Block {
	var newBlock Block
	newBlock.Index = lastBlock.Index + 1
	newBlock.TimeStamp = time.Now().Unix()
	newBlock.Data = data
	newBlock.Prehash = lastBlock.Hash
	newBlock.Difficulty = difiiculty
	newBlock.Nonce = 0
	//开挖-当前区块的hash值的前面的0的个数与难度系数值相同
	for {
		//计算hash
		cuhash := hex.EncodeToString(BlockHash(newBlock))
		//fmt.Println("挖矿中: ", cuhash)
		newBlock.Hash = cuhash
		if isBlockValid(newBlock) {
			//校验区块
			if VerflyBlock(newBlock, *lastBlock) {
				fmt.Println("挖矿成功: ", cuhash)
				return &newBlock
			}
		}

		newBlock.Nonce++
	}
}

func VerflyBlock(newblock Block, lastBlock Block) bool {
	if lastBlock.Index+1 != newblock.Index {
		return false
	}
	if newblock.Prehash != lastBlock.Hash {
		return false
	}
	return true
}

func main() {
	var genBlock = GenesisBlock()
	var newBlock *Block
	newBlock = genBlock
	for i := 0; i < 10; i++ {
		newBlock = CreateNewBlock(newBlock, fmt.Sprintf("new block %d", i))
		difiiculty = i + 2
		fmt.Print("New block info: \n")
		fmt.Printf("height [%d], hash [%s], data [%s], nonce [%d], difficulty [%d].\n",
			newBlock.Index, newBlock.Hash, newBlock.Data, newBlock.Nonce, newBlock.Difficulty)
	}
}
