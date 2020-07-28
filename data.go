package main

import (
	"github.com/ktsymbal/simple-blockchain-service/blockchain"
	"sync"
)

type Data struct {
	BC *blockchain.Blockchain
	Buffer []string
	Mux *sync.RWMutex
}

func NewData(bufferCapacity int) Data {
	if bufferCapacity < 1 { bufferCapacity = 1 }
	data := Data{blockchain.NewBlockchain(), make([]string, 0, bufferCapacity), &sync.RWMutex{}}
	return data
}

func (data *Data) AddData(newData string) {
	data.Mux.Lock()
	defer data.Mux.Unlock()
	data.Buffer = append(data.Buffer, newData)
	if len(data.Buffer) == cap(data.Buffer) {
		data.BC.AddNewBlock(data.Buffer...)
		data.Buffer = make([]string, 0, cap(data.Buffer))
	}
}

func (data Data) GetNLastBlocks(n int) []blockchain.Block {
	data.Mux.RLock()
	defer data.Mux.RUnlock()
	return data.BC.GetNLastBlocks(n)
}