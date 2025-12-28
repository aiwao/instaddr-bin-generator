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

			accounts, err := common.GetAccountsFromDatabase(db)
			if err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

			responseJSON := common.ResponseJSON{
				AccountAmount: len(accounts),
				Accounts:      accounts,
			}
			b, err := json.Marshal(&responseJSON)
			if err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write(b)
		})
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
}
