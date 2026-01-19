package generator

import (
	"database/sql"
	"fmt"
	"instaddr_bin/utility"
	"log"
	"math/rand/v2"
	"net/http"
	"net/url"
	"time"

	"github.com/aiwao/envar"
	instaddr "github.com/aiwao/instaddr_api"
)

var domains []string
var createAccountDelay = 1000
var createAddressDelay = 1000
var onErrorDelay = 5000
var addressAmount = 50
var mustLegitToAmount = true
var proxy *url.URL

func Start(db *sql.DB) {
	envar.GetIntv("CREATE_ACCOUNT_DELAY", &createAccountDelay)
	envar.GetIntv("CREATE_ADDRESS_DELAY", &createAddressDelay)
	envar.GetIntv("ON_ERROR_DELAY", &onErrorDelay)
	envar.GetIntv("ADDRESS_AMOUNT", &addressAmount)
	envar.GetBoolv("MUST_LEGIT_TO_AMOUNT", &mustLegitToAmount)
	proxy, _ = envar.GetURL("PROXY")

	log.Printf("CREATE_ACCOUNT_DELAY: %d\n", createAccountDelay)
	log.Printf("CREATE_ADDRESS_DELAY: %d\n", createAddressDelay)
	log.Printf("ON_ERROR_DELAY: %d\n", onErrorDelay)
	log.Printf("MUST_LEGIT_TO_AMOUNT: %v\n", mustLegitToAmount)
	log.Printf("PROXY: %s\n", proxy.String())

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
			log.Printf("%s%v%s\n", utility.Red, err, utility.Reset)
			time.Sleep(time.Duration(onErrorDelay) * time.Millisecond)
			continue
		}

		domain := "mail4.uk"
		if domains == nil || len(domains) == 0 {
			domains, err = acc.GetMailDomains(option)
			if err != nil {
				log.Printf("%s%v%s\n", utility.Red, err, utility.Reset)
			}
		}

		tried := 1
		created := 1
		resultStr := ""
		for {
			resultStr = fmt.Sprintf("[Tried :%d, Created :%d]", tried, created)
			if (tried >= addressAmount && !mustLegitToAmount) || created >= addressAmount {
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
				log.Printf("%s%s %v%s\n", utility.Red, resultStr, err, utility.Reset)
				time.Sleep(time.Duration(onErrorDelay) * time.Millisecond)
				continue
			}
			log.Printf("%s%s%s %s%s%s\n", utility.Green, resultStr, utility.Reset, utility.Blue, mailAcc.Address, utility.Reset)
			time.Sleep(time.Duration(createAddressDelay) * time.Millisecond)
			created++
		}
		log.Printf("%s%s%s\n", utility.Blue, resultStr, utility.Reset)
		info, err := acc.GetAuthInfo(option)
		if err != nil {
			log.Printf("%s%v%s\n", utility.Red, err, utility.Reset)
		} else {
			_, err := db.Exec(
				"INSERT INTO accounts(id, password, amount) VALUES ($1, $2, $3)",
				info.AccountID, info.Password, created,
			)
			if err != nil {
				log.Printf("%s%v%s\n", utility.Red, err, utility.Reset)
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
