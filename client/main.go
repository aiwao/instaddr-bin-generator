package main

import (
	"bytes"
	"client/utility"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"common"
)

func main() {
	accountAmount := utility.ScanInt("Maximum account amount to get")
	minAddressAmount := utility.ScanInt("Minimum amount of addresses of account")
	payload := common.RequestJSON{
		AccountAmount:    accountAmount,
		MinAddressAmount: minAddressAmount,
	}
	jsonBytes, err := json.Marshal(&payload)
	if err != nil {
		fmt.Printf("%s%v%s\n", common.Red, err, common.Reset)
		return
	}

	res, err := http.Post("http://localhost:8080", "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Printf("%s%v%s\n", common.Red, err, common.Reset)
		return
	}
	defer res.Body.Close()
	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("%s%v%s\n", common.Red, err, common.Reset)
		return
	}

	if res.StatusCode != http.StatusOK {
		fmt.Printf("%s%d %s%s\n", common.Red, res.StatusCode, string(resBytes), common.Reset)
		return
	}

	var responseJSON common.ResponseJSON
	if err := json.Unmarshal(resBytes, &responseJSON); err != nil {
		fmt.Printf("%s%v%s\n", common.Red, err, common.Reset)
		return
	}

	fmt.Printf("Got %d Accounts\n", responseJSON.AccountAmount)
	for _, acc := range responseJSON.Accounts {
		fmt.Printf("%s%s%s:%s%s%s %s(%d Addresses)%s\n", common.Blue, acc.ID, common.Reset, common.Blue, acc.Password, common.Reset, common.Green, acc.AddressAmount, common.Reset)
	}
}
