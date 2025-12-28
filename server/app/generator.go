package app

import (
	"common"
	"database/sql"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"net/url"
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
var proxy *url.URL

func StartGenerator(db *sql.DB) {
	accDelayParsed, err := strconv.Atoi(os.Getenv("CREATE_ACCOUNT_DELAY"))
	if err == nil {
		createAccountDelay = accDelayParsed
		log.Printf("CREATE_ACCOUNT_DELAY: %d\n", accDelayParsed)
	}
	addrDelayParsed, err := strconv.Atoi(os.Getenv("CREATE_ADDRESS_DELAY"))
	if err == nil {
		createAddressDelay = addrDelayParsed
		log.Printf("CCREATE_ADDRESS_DELAY: %d\n", addrDelayParsed)
	}
	errDelayParsed, err := strconv.Atoi(os.Getenv("ON_ERROR_DELAY"))
	if err == nil {
		onErrorDelay = errDelayParsed
		log.Printf("ON_ERROR_DELAY: %d\n", errDelayParsed)
	}
	addrAmountParsed, err := strconv.Atoi(os.Getenv("ADDRESS_AMOUNT"))
	if err == nil {
		addressAmount = addrAmountParsed
		log.Printf("ADDRESS_AMOUNT: %d\n", addrAmountParsed)
	}
	mustLegitParsed, err := strconv.Atoi(os.Getenv("MUST_LEGIT_TO_AMOUNT"))
	if err == nil {
		mustLegitToAmount = mustLegitParsed == 1
		log.Printf("MUST_LEGIT_TO_AMOUNT: %v\n", mustLegitParsed == 1)
	}
	proxyEnv := os.Getenv("PROXY")
	proxyURLParsed, err := url.Parse(proxyEnv)
	if err == nil {
		proxy = proxyURLParsed
		log.Printf("PROXY: %s\n", proxyEnv)
	}

	for {
		client := &http.Client{}
		if proxy != nil {
			client.Transport = &http.Transport{
				Proxy: http.ProxyURL(proxy),
			}
		}
		option := instaddr.Options{
			Client: client,
		}

		acc, err := instaddr.NewAccount(option)
		if err != nil {
			log.Printf("%s%v%s\n", common.Red, err, common.Reset)
			time.Sleep(time.Duration(onErrorDelay) * time.Millisecond)
			continue
		}

		domain := "mail4.uk"
		if domains == nil || len(domains) == 0 {
			domains, err = acc.GetMailDomains(option)
			if err != nil {
				log.Printf("%s%v%s\n", common.Red, err, common.Reset)
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
				Name:    random(),
				Options: option,
			}, domain)
			tried++
			if err != nil {
				log.Printf("%s%s %v%s\n", common.Red, resultStr, err, common.Reset)
				time.Sleep(time.Duration(onErrorDelay) * time.Millisecond)
				continue
			}
			log.Printf("%s%s%s %s%s%s\n", common.Green, resultStr, common.Reset, common.Blue, mailAcc.Address, common.Reset)
			time.Sleep(time.Duration(createAddressDelay) * time.Millisecond)
			created++
		}
		log.Printf("%s%s%s\n", common.Blue, resultStr, common.Reset)
		info, err := acc.GetAuthInfo(option)
		if err != nil {
			log.Printf("%s%v%s\n", common.Red, err, common.Reset)
		} else {
			_, err := db.Exec(
				"INSERT INTO accounts(id, password, amount) VALUES (?, ?, ?)",
				info.AccountID, info.Password, created,
			)
			if err != nil {
				log.Printf("%s%v%s\n", common.Red, err, common.Reset)
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
