package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"
	"sync/atomic"
	"time"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/wangdayong228/minimal_transfer/cfx_tx_engine"
	minimaltransfer "github.com/wangdayong228/minimal_transfer/contracts"
)

func main() {

	am := initAccountManager()
	cAddr := cfxaddress.MustNewFromBase32("net8888:aam1eawbm9pzp0dnwv96tts5shnbdfv9nucbh4c593")

	normalSum, cSum := int32(0), int32(0)

	go func() {
		client := sdk.MustNewClient("http://net8888cfx.confluxrpc.com", sdk.ClientOption{
			RetryCount: 3,
		})
		client.SetAccountManager(am)
		users := client.AccountManager.List()

		instance := deploy(client, cAddr)
		for {
			from := users[rand.Int()%len(users)]
			to := users[rand.Int()%len(users)]

			delta := rand.Int() % 700

			_, hash, err := instance.SimpleTransferTransactor.Run(&types.ContractMethodSendOption{
				From:     &cAddr,
				GasPrice: types.NewBigInt(2e9),
			}, from.MustGetCommonAddress(), to.MustGetCommonAddress(), big.NewInt(int64(5000+delta)))

			if err != nil {
				fmt.Printf("cwei err %v\n", err)
				time.Sleep(time.Second)
				continue
			}
			time.Sleep(time.Second / 5)
			fmt.Printf("cwei hash %v\n", hash)

			atomic.AddInt32(&cSum, 1)
			if cSum%10 == 0 {
				fmt.Printf("[%v] cwei sent %v\n", time.Now(), cSum)
			}
		}
	}()

	urls := []string{
		"http://101.132.158.162:12537",
		// "http://39.100.97.209:12537",
		"http://47.92.105.52:12537",
		"http://47.92.7.84:12537",
	}

	for ui := 0; ui < len(urls); ui++ {
		subClient := sdk.MustNewClient(urls[ui], sdk.ClientOption{
			RetryCount: 3,
		})
		subClient.SetAccountManager(am)
		users := subClient.AccountManager.List()[:100]

		userGroupLen := len(users) / len(urls)
		for i := 0; i < userGroupLen; i++ {

			go func(index int, _client *sdk.Client) {
				if users[index].String() == cAddr.String() {
					return
				}

				nodeurl := _client.GetNodeURL()

				for {
					status, err := _client.GetStatus()
					if err != nil {
						time.Sleep(time.Second)
						fmt.Printf("get status err:%v\n", err)
						continue
					}
					for j := 0; j < 1000; j++ {
						tx := types.UnsignedTransaction{}
						from := users[index] //users[rand.Int()%len(users)]
						tx.From = &from
						tx.To = &from
						tx.Gas = types.NewBigInt(21000)
						tx.GasPrice = types.NewBigInt(1e9)
						tx.EpochHeight = &status.EpochNumber
						tx.StorageLimit = types.NewUint64(0)
						tx.ChainID = types.NewUint(uint(uint64(status.ChainID)))

						_, err := _client.SendTransaction(tx)
						if err != nil {
							fmt.Printf("%v normal err: %v\n", nodeurl, err)
							time.Sleep(time.Second)
							continue
						}
						// fmt.Printf("normal hash %v\n", hash)
						atomic.AddInt32(&normalSum, 1)
						if normalSum%1000 == 0 {
							fmt.Printf("[%v] sent %v\n", time.Now(), normalSum)
						}
					}
				}

			}(i+ui*userGroupLen, subClient)
		}
	}
	select {}
}

func initAccountManager() sdk.AccountManagerOperator {
	am := cfx_tx_engine.NewPrivatekeyAccountManager(nil, 8888)
	am.ImportKey("0f7d769ee463fe97d40dc5a527f3140dcd83fabb7846a15d00142cd1821e8979", "")

	content, err := ioutil.ReadFile("/Users/dayong/myspace/mywork/minimal_transfer/private_keys.json")
	if err != nil {
		panic(err)
	}

	accounts := []PrivatekeyAccount{}
	err = json.Unmarshal(content, &accounts)
	if err != nil {
		panic(err)
	}

	for _, a := range accounts {
		if _, err = am.ImportKey(a.Privatekey, ""); err != nil {
			panic(err)
		}
	}

	return am
}

func deploy(client *sdk.Client, cAddr cfxaddress.Address) *minimaltransfer.SimpleTransfer {
	_, hash, _, err := minimaltransfer.DeploySimpleTransfer(&types.ContractMethodSendOption{From: &cAddr}, client)
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
	return instance
}
