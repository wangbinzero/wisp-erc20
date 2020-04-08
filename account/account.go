package account

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"wisp-erc20/client"
	common2 "wisp-erc20/common"
)

// 根据账户地址查询最新余额
func GetBalanceByAddress(address string) *big.Float {

	account := common.HexToAddress(address)
	balance, err := client.Client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		fmt.Println("查询账户余额失败, 账户地址为: ", account)
	}

	return common2.WeiToEth(balance)

}

// 根据区账户地址和块编号查询账户余额
func GetBalanceByBlock(address string, block int64) *big.Float {
	account := common.HexToAddress(address)
	blockNumber := big.NewInt(block)
	balance, err := client.Client.BalanceAt(context.Background(), account, blockNumber)

	if err != nil {
		fmt.Printf("查询账户余额失败,账户地址为: %s 区块编号为: %d/n", address, block)
	}
	return common2.WeiToEth(balance)
}

// 根据地址查询待处理账户余额
func GetPendingBalance(address string) *big.Float {
	account := common.HexToAddress(address)
	pendingBalance, err := client.Client.PendingBalanceAt(context.Background(), account)

	if err != nil {
		fmt.Println("查询待处理账户余额失败,账户地址为: ", address)
	}
	return common2.WeiToEth(pendingBalance)
}
