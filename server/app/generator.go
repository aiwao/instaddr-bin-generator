package app

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strconv"
	"time"

	instaddr "github.com/aiwao/instaddr_api"
)

var domains []string
var createAccountDelay = 1000
var createAddressDelay = 1000
var onErrorDelay = 5000
var addressAmount = 50
var mustLegitToAmount = false

func StartGenerator(db *sql.DB) {
	accDelayParse, err := strconv.Atoi(os.Getenv("CREATE_ACCOUNT_DELAY"))
	if err == nil {
		createAccountDelay = accDelayParse
	}
	addrDelayParse, err := strconv.Atoi(os.Getenv("CREATE_ADDRESS_DELAY"))
	if err == nil {
		createAddressDelay = addrDelayParse
	}
	errDelayParse, err := strconv.Atoi(os.Getenv("ON_ERROR_DELAY"))
	if err == nil {
		onErrorDelay = errDelayParse
	}
	addrAmountParse, err := strconv.Atoi(os.Getenv("ADDRESS_AMOUNT"))
	if err == nil {
		addressAmount = addrAmountParse
	}
	mustLegitParse, err := strconv.Atoi(os.Getenv("MUST_LEGIT_TO_AMOUNT"))
	if err == nil {
		mustLegitToAmount = mustLegitParse == 1
	}

	for {
		acc, err := instaddr.NewAccount(instaddr.Options{})
		if err != nil {
			log.Println(err)
			log.Println("Retry in 5 seconds")
			time.Sleep(time.Duration(onErrorDelay) * time.Millisecond)
			continue
		}

		domain := "mail4.uk"
		if domains == nil || len(domains) == 0 {
			domains, err = acc.GetMailDomains(instaddr.Options{})
			if err != nil {
				log.Println(err)
			}
		}

		tried := 0
		created := 0
		resultStr := ""
		for {
			resultStr = fmt.Sprintf("[Tried :%d, Created :%d]", tried, created)
			if tried == addressAmount && !mustLegitToAmount {
				break
			}
			if created == addressAmount {
				break
			}
			if domains != nil && len(domains) > 0 {
				domain = domains[rand.IntN(len(domains))]
			}
			mailAcc, err := acc.CreateAddressWithDomainAndName(instaddr.OptionsWithName{
				Name: random(),
			}, domain)
			tried++
			if err != nil {
				log.Printf("%s %v\n", resultStr, err)
				time.Sleep(time.Duration(onErrorDelay) * time.Millisecond)
				continue
			}
			log.Printf("%s %s\n", resultStr, mailAcc.Address)
			time.Sleep(time.Duration(createAddressDelay) * time.Millisecond)
			created++
		}
		log.Printf(resultStr)
		info, err := acc.GetAuthInfo(instaddr.Options{})
		if err != nil {
			log.Println(err)
		} else {
			_, err := db.Exec(
				"INSERT INTO accounts(id, password, amount) VALUES (?, ?, ?)",
				info.AccountID, info.Password, created,
			)
			if err != nil {
				log.Println(err)
			}
		}
		time.Sleep(time.Duration(createAccountDelay) * time.Millisecond)
	}
}

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

func random() string {
	b := make([]byte, 39)
	for i := range b {
		b[i] = charset[rand.IntN(len(charset))]
	}
	return string(b)
}
