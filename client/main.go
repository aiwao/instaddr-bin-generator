package main

import (
	"bin-gen-client/utility"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"common"
)

func main() {
	accountAmount := utility.ScanInt("Account amount to get")
	minAddressAmount := utility.ScanInt("Minimum amount of addresses of account")
	payload := common.RequestJSON{
		AccountAmount:    accountAmount,
		MinAddressAmount: minAddressAmount,
	}
	jsonBytes, err := json.Marshal(&payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := http.Post("http://localhost:8080", "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	if res.StatusCode != http.StatusOK {
		fmt.Printf("%d %s¥n", res.StatusCode, string(resBytes))
		return
	}

	var responseJSON common.ResponseJSON
	if err := json.Unmarshal(resBytes, &responseJSON); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Got %d Accounts", responseJSON.AccountAmount)
	for _, acc := range responseJSON.Accounts {
		fmt.Printf("[%d Addresses] %s:%s¥n", acc.AddressAmount, acc.ID, acc.Password)
	}
}
