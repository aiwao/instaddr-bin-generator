package common

type RequestJSON struct {
	AccountAmount    int `json:"account_amount"`
	MinAddressAmount int `json:"min_address_amount"`
}

type Account struct {
	ID            string `json:"id"`
	Password      string `json:"password"`
	AddressAmount int    `json:"address_amount"`
}

type ResponseJSON struct {
	AccountAmount int       `json:"account_amount"`
	Accounts      []Account `json:"accounts"`
}
