package main

import (
	"github.com/wangdayong228/benchmark_conflux_transfer/cfx_tx_engine"
)

// type PrivatekeyAccount struct {
// 	Privatekey string
// 	Address    cfxaddress.Address
// }

func genPrivatekeyAccounts(chainId uint32, num int) []string {
	am := cfx_tx_engine.NewPrivatekeyAccountManager(nil, chainId)

	keys := []string{}

	for i := 0; i < num; i++ {
		addr, err := am.Create("")
		if err != nil {
			panic(err)
		}
		privKey, err := am.Export(addr, "")
		if err != nil {
			panic(err)
		}

		keys = append(keys, privKey)
	}

	return keys
}
