package main

import (
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/wangdayong228/minimal_transfer/cfx_tx_engine"
)

type PrivatekeyAccount struct {
	Privatekey string
	Address    cfxaddress.Address
}

func genPrivatekeyAccounts(num int) []PrivatekeyAccount {
	am := cfx_tx_engine.NewPrivatekeyAccountManager(nil, 8888)

	keys := []PrivatekeyAccount{}

	for i := 0; i < num; i++ {
		addr, err := am.Create("")
		if err != nil {
			panic(err)
		}
		privKey, err := am.Export(addr, "")
		if err != nil {
			panic(err)
		}

		keys = append(keys, PrivatekeyAccount{privKey, addr})
	}

	return keys
}
