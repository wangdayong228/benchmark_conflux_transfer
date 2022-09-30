package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/Conflux-Chain/go-conflux-sdk/types/unit"
	"github.com/stretchr/testify/assert"
)

func TestGenAccounts(t *testing.T) {
	keys := genPrivatekeyAccounts(1, 1000)
	fmt.Println(keys)
	j, _ := json.Marshal(keys)
	err := ioutil.WriteFile("/Users/dayong/myspace/mywork/minimal_transfer/private_keys.json", j, fs.ModePerm)
	assert.NoError(t, err)
}

func TestDispatchTreasure(t *testing.T) {

	// url := "http://net8888cfx.confluxrpc.com"
	// cAddr := cfxaddress.MustNewFromBase32("net8888:aam1eawbm9pzp0dnwv96tts5shnbdfv9nucbh4c593")

	url := "http://test.confluxrpc.com"
	cAddr := cfxaddress.MustNewFromBase32("cfxtest:aam1eawbm9pzp0dnwv96tts5shnbdfv9nuwu7zgzz8")

	client := sdk.MustNewClient(url, sdk.ClientOption{
		RetryCount: 3,
	})
	client.SetAccountManager(initAccountManager())
	am := initAccountManager()
	d, _ := unit.NewDripFromString("10 CFX")

	users := am.List()
	for i := 0; i < len(users); i++ {
		tx := types.UnsignedTransaction{}
		tx.From = &cAddr
		tx.To = &users[i]

		tx.Value = types.NewBigIntByRaw(d.BigInt())
		hash, err := client.SendTransaction(tx)
		assert.NoError(t, err)
		fmt.Printf("send to %v hash %v\n", tx.To, hash)
	}
}
