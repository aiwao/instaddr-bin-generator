package app

import (
	"common"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func StartAPI(db *sql.DB) {
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
				return
			}

			var parsedBody common.RequestJSON
			err := json.NewDecoder(r.Body).Decode(&parsedBody)
			if err != nil {
				http.Error(w, "invalid body", http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			rows, err := db.Query(
				"SELECT * FROM accounts WHERE amount >= ? ORDER BY created_at DESC LIMIT ?",
				parsedBody.MinAddressAmount,
				parsedBody.AccountAmount,
			)
			if err != nil {
				log.Printf("A: %v", err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}
			defer rows.Close()

			accounts := []common.Account{}
			for rows.Next() {
				var account common.Account
				if err := rows.Scan(&account.ID, &account.Password, &account.AddressAmount, &account.CreatedAt); err != nil {
					log.Printf("B: %v", err)
					http.Error(w, "internal server error", http.StatusInternalServerError)
					return
				}
				accounts = append(accounts, account)
			}

			responseJSON := common.ResponseJSON{
				AccountAmount: len(accounts),
				Accounts:      accounts,
			}
			b, err := json.Marshal(&responseJSON)
			if err != nil {
				log.Printf("C: %v", err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write(b)
		})
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
}
