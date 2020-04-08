package client

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
)

var Client *ethclient.Client

func Init() {
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		fmt.Println("连接 [eth] 主网失败")
		panic(err)
	}
	Client = client
}
