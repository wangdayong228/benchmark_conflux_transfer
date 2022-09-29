package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	minimaltransfer "github.com/wangdayong228/minimal_transfer/contracts"
)

func main() {
	client := sdk.MustNewClient("http://net8888cfx.confluxrpc.com", sdk.ClientOption{
		KeystorePath: "keystore",
	})

	if len(client.AccountManager.List()) == 0 {
		_, err := client.AccountManager.ImportKey("0f7d769ee463fe97d40dc5a527f3140dcd83fabb7846a15d00142cd1821e8979", "")
		if err != nil {
			panic(err)
		}
	}

	addr, err := client.AccountManager.GetDefault()
	if err != nil {
		panic(err)
	}

	err = client.AccountManager.UnlockDefault("")
	if err != nil {
		panic(err)
	}

	_, hash, _, err := minimaltransfer.DeploySimpleTransfer(&types.ContractMethodSendOption{From: addr}, client)
	if err != nil {
		panic(err)
	}

	r, err := client.WaitForTransationReceipt(*hash, time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Printf("deploy hash %v, contract address %v\n", hash, r.ContractCreated)

	instance, err := minimaltransfer.NewSimpleTransfer(*r.ContractCreated, client)
	if err != nil {
		panic(err)
	}

	users := []common.Address{
		common.HexToAddress("0x292f04a44506c2fd49bac032e1ca148c35a478c8"),
		common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
		common.HexToAddress("0x55fe002aeff02f77364de339a1292923a15844b8"),
		common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
		common.HexToAddress("0x9a49227a58a945cc9ba13f16debdce141b8f07b7"),
		common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7"),
	}

	for {
		from := users[rand.Int()%len(users)]
		to := users[rand.Int()%len(users)]

		delta := rand.Int() % 700
		_, hash, err = instance.SimpleTransferTransactor.Run(nil, from, to, big.NewInt(int64(2800+delta)))
		if err != nil {
			panic(err)
		}
		fmt.Printf("c hash %v\n", hash)
	}
}

// func initAccountManager() {
// 	am := cfx_tx_engine.NewPrivatekeyAccountManager(nil, 8888)

// 	_, err := client.AccountManager.ImportKey("0f7d769ee463fe97d40dc5a527f3140dcd83fabb7846a15d00142cd1821e8979", "")

// 	addr, err := client.AccountManager.GetDefault()
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = client.AccountManager.UnlockDefault("")
// 	if err != nil {
// 		panic(err)
// 	}
// }
