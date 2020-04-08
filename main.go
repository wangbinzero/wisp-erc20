package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	client2 "wisp-erc20/client"
)

func main() {

	client2.Init()

	address := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	fmt.Println(address.Hex())
	fmt.Println(address.Hash().Hex())
	fmt.Println("以太主网连接成功")

	fmt.Println("查询余额")
	account := common.HexToAddress("0x812F6Fe50b7189CB2f6031846cf9Ef0f395a1C20")
	balance, err := client2.Client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("账户余额为: ", balance)
}
