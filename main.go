package main

import (
	"fmt"
	"wisp-erc20/account"
	client2 "wisp-erc20/client"
)

func main() {

	client2.Init()

	balance := account.GetBalanceByAddress("0x812F6Fe50b7189CB2f6031846cf9Ef0f395a1C20")
	fmt.Println("根据地址查询余额为: ", balance)

	balance1 := account.GetBalanceByBlock("0x812F6Fe50b7189CB2f6031846cf9Ef0f395a1C20", 0)
	fmt.Println("根据地址以及区块号查询余额为: ", balance1)
}
