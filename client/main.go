package main

import (
	"client/api"
	"common"
	"log"
	"os"
	"strconv"
)

var local = false
var accountAmount = 100
var minAddressAmount = 10

func main() {
	local = os.Getenv("LOCAL") == "1"
	log.Printf("Local: %v\n", local)
	var serverURL string
	if !local {
		serverURLEnv := os.Getenv("SERVER_URL")
		if serverURLEnv == "" {
			serverURL = "http://localhost:8080"
		} else {
			serverURL = serverURLEnv
			log.Println("Server URL: " + serverURLEnv)
		}
	}
	accAmParsed, err := strconv.Atoi(os.Getenv("AMOUNT_ACCOUNT"))
	if err == nil {
		accountAmount = accAmParsed
		log.Printf("Account amount: %d\n", accAmParsed)
	}
	minAddrParsed, err := strconv.Atoi(os.Getenv("MIN_AMOUNT_ADDRESS"))
	if err == nil {
		minAddressAmount = minAddrParsed
		log.Printf("Minimum account amount: %d\n", minAddressAmount)
	}

	payload := api.ClientRequestJSON{
		Local:     local,
		ServerURL: serverURL,
		RequestJSON: common.RequestJSON{
			AccountAmount:    accountAmount,
			MinAddressAmount: minAddressAmount,
		},
	}
	res, err := api.RequestDatabase(payload)
	if err != nil {
		log.Printf("%s%v%s\n", common.Red, err, common.Reset)
		return
	}

	log.Printf("%sGot%s %s%d%s%s Accounts%s\n", common.Blue, common.Reset, common.Green, res.AccountAmount, common.Reset, common.Blue, common.Reset)
	totalAddresses := 0
	for _, acc := range res.Accounts {
		log.Printf("%s%s%s:%s%s%s %s(%d Addresses)%s\n", common.Blue, acc.ID, common.Reset, common.Blue, acc.Password, common.Reset, common.Green, acc.AddressAmount, common.Reset)
		totalAddresses += acc.AddressAmount
	}
	log.Printf("%sTotal addresses:%s %s%d%s\n", common.Blue, common.Reset, common.Green, totalAddresses, common.Reset)
}
