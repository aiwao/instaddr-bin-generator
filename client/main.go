package main

import (
	"client/api"
	"common"
	"flag"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var dbPath string
	var serverURL string
	flag.StringVar(&dbPath, "d", "", "local database file path")
	flag.StringVar(&serverURL, "s", "http://localhost:8080", "custom server address (default: localhost:8080)")
	var accountAmount int
	flag.IntVar(&accountAmount, "acc", 1, "maximum account amount to get")
	var minAddressAmount int
	flag.IntVar(&minAddressAmount, "addr", 0, "minimum amount of addresses in account")
	flag.Parse()

	payload := api.ClientRequestJSON{
		DBPath:    dbPath,
		ServerURL: serverURL,
		RequestJSON: common.RequestJSON{
			AccountAmount:    accountAmount,
			MinAddressAmount: minAddressAmount,
		},
	}
	res, err := api.RequestDatabase(payload)
	if err != nil {
		fmt.Printf("%s%v%s\n", common.Red, err, common.Reset)
		return
	}

	fmt.Printf("%sGot%s %s%d%s%s Accounts%s\n", common.Blue, common.Reset, common.Green, res.AccountAmount, common.Reset, common.Blue, common.Reset)
	totalAddresses := 0
	for _, acc := range res.Accounts {
		fmt.Printf("%s%s%s:%s%s%s %s(%d Addresses)%s\n", common.Blue, acc.ID, common.Reset, common.Blue, acc.Password, common.Reset, common.Green, acc.AddressAmount, common.Reset)
		totalAddresses += acc.AddressAmount
	}
	fmt.Printf("%sTotal addresses:%s %s%d%s\n", common.Blue, common.Reset, common.Green, totalAddresses, common.Reset)
}
