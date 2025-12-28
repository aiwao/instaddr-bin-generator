package common

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

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

func GetAccountsFromDatabase(db *sql.DB, payload RequestJSON) ([]Account, error) {
	rows, err := db.Query(
		"SELECT * FROM accounts WHERE amount >= ? ORDER BY created_at DESC LIMIT ?",
		payload.MinAddressAmount,
		payload.AccountAmount,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accounts := []Account{}
	for rows.Next() {
		var account Account
		if err := rows.Scan(&account.ID, &account.Password, &account.AddressAmount, &account.CreatedAt); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}
