package network

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/uptempotech/golang_blockchain/core"
)

// SendAddr ...
func SendAddr(address string) {
	log.Println("Entering SendAddr")
	nodes := Addr{KnownNodes}
	nodes.AddrList = append(nodes.AddrList, nodeAddress)
	payload := GobEncode(nodes)
	request := append(CmdToBytes("addr"), payload...)

	log.Printf("SendAddr: %s", address)
	SendData(address, request)
}

// SendBlock ...
func SendBlock(addr string, b *core.Block) {
	log.Println("Entering SendBlock")
	data := Block{nodeAddress, b.Serialize()}
	payload := GobEncode(data)
	request := append(CmdToBytes("block"), payload...)

	SendData(addr, request)
}

// SendData ...
func SendData(addr string, data []byte) {
	log.Println("Entering SendData")
	conn, err := net.Dial(protocol, addr)
	if err != nil {
		fmt.Printf("%s is not available\n", addr)
		var updatedNodes []string

		for _, node := range KnownNodes {
			if node != addr {
				updatedNodes = append(updatedNodes, node)
			}
		}

		KnownNodes = updatedNodes

		return
	}

	defer conn.Close()

	_, err = io.Copy(conn, bytes.NewReader(data))
	if err != nil {
		log.Panic(err)
	}
}

// SendInv ...
func SendInv(address, kind string, items [][]byte) {
	log.Println("Entering SendInv")
	inventory := Inv{nodeAddress, kind, items}
	payload := GobEncode(inventory)
	request := append(CmdToBytes("inv"), payload...)

	SendData(address, request)
}

// SendGetBlocks ...
func SendGetBlocks(address string) {
	log.Println("Entering SendGetBlocks")
	payload := GobEncode(GetBlocks{nodeAddress})
	request := append(CmdToBytes("getblocks"), payload...)

	SendData(address, request)
}

// SendGetData ...
func SendGetData(address, kind string, id []byte) {
	log.Println("Entering SendGetData")
	payload := GobEncode(GetData{nodeAddress, kind, id})
	request := append(CmdToBytes("getdata"), payload...)

	SendData(address, request)
}

// SendTx ...
func SendTx(addr string, tnx *core.Transaction) {
	log.Println("Entering SendTX")
	data := Tx{nodeAddress, tnx.Serialize()}
	payload := GobEncode(data)
	request := append(CmdToBytes("tx"), payload...)

	SendData(addr, request)
}

// SendVersion ...
func SendVersion(addr string, chain *core.BlockChain) {
	log.Println("Entering SendVersion")
	bestHeight := chain.GetBestHeight()
	payload := GobEncode(Version{version, bestHeight, nodeAddress})
	request := append(CmdToBytes("version"), payload...)

	SendData(addr, request)
}
