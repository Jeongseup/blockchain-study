package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

/*
func (b *Block) SetHash() {
	// 뭔지는 모르겠지만 타임스탬프 작성?
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp})

}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}
*/
func main() {
	now := time.Now().UnixMicro()
	timestamp := []byte(strconv.FormatInt(now, 10))
	fmt.Println(timestamp)

	// data := []string{"data"}
	// prevBlockHash := "0x01"
	// headers := bytes.Join([][]byte{[]byte(prevBlockHash), data, timestamp}, []byte{})
	hash := sha256.Sum256(timestamp)
	// b.Hash = hash[:]
	fmt.Println(hash)

	// https://m.blog.naver.com/pjt3591oo/221354809830
}
