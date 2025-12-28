package common

import "time"

type RequestJSON struct {
	AccountAmount    int `json:"account_amount"`
	MinAddressAmount int `json:"min_address_amount"`
}

type Account struct {
	ID            string    `json:"id"`
	Password      string    `json:"password"`
	AddressAmount int       `json:"address_amount"`
	CreatedAt     time.Time `json:"created_at"`
}

type ResponseJSON struct {
	AccountAmount int       `json:"account_amount"`
	Accounts      []Account `json:"accounts"`
}

const Reset = "\033[0m"
const Green = "\033[92m"
const Red = "\033[31m"
const Blue = "\033[94m"
