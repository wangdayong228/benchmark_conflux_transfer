package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenAccounts(t *testing.T) {
	keys := genPrivatekeyAccounts(100)
	fmt.Println(keys)
	j, _ := json.Marshal(keys)
	err := ioutil.WriteFile("/Users/dayong/myspace/mywork/minimal_transfer/private_keys.json", j, fs.ModePerm)
	assert.NoError(t, err)
}
