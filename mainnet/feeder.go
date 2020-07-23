package main

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"net/http"
)

type  FeederInfo struct {
	account string
	websites []string
}

type ContractFeedingInfo struct {
	symbol string
	wallet_api_url string
	price_feeder_contract_address string
	feeders []string
	feedernames []string
	inteverl int
}

type PriceGrab struct {
	symbolpair string
	exchangeWebSiteName string
	url string
	accessPriceKeys []string
	muti_factor int
}

func (pGrab * PriceGrab) grab_price()  (string ,error){
	var result = ""
	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", pGrab.url,nil)
	resp, err := client.Do(reqest)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return result,err
	}
	body, _ := ioutil.ReadAll(resp.Body)

	var m map[string]interface{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		return result,err
	}
	var price=""
	for k,v :=range pGrab.accessPriceKeys {
		var temp map[string]interface{}
		if k == 0 {
			temp = m[v].(map[string]interface{})
		}else {
			price = temp[v].(string)
		}
	}
	if pGrab.muti_factor != 1 {
		n_price,_:= decimal.NewFromString(price)
		n_price =n_price.Round(10000000)
		result = n_price.String()
	}
	return result,nil
}