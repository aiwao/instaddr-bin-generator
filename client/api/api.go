package api

import (
	"bytes"
	"common"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type ClientRequestJSON struct {
	DBPath    string
	ServerURL string
	common.RequestJSON
}

func RequestDatabase(payload ClientRequestJSON) (common.ResponseJSON, error) {
	if payload.DBPath != "" {
		db, err := sql.Open("sqlite3", payload.DBPath)
		if err != nil {
			return common.ResponseJSON{}, err
		}
		defer db.Close()
		accounts, err := common.GetAccountsFromDatabase(db, payload.RequestJSON)
		if err != nil {
			return common.ResponseJSON{}, err
		}

		return common.ResponseJSON{
			AccountAmount: len(accounts),
			Accounts:      accounts,
		}, nil
	}

	jsonBytes, err := json.Marshal(&payload.RequestJSON)
	if err != nil {
		return common.ResponseJSON{}, err
	}

	res, err := http.Post(payload.ServerURL, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return common.ResponseJSON{}, err
	}
	defer res.Body.Close()
	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return common.ResponseJSON{}, err
	}

	if res.StatusCode != http.StatusOK {
		return common.ResponseJSON{}, errors.New(fmt.Sprintf("%d %s", res.StatusCode, string(resBytes)))
	}

	var responseJSON common.ResponseJSON
	if err := json.Unmarshal(resBytes, &responseJSON); err != nil {
		return common.ResponseJSON{}, err
	}

	return responseJSON, nil
}
