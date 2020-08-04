package main

import (
	"../hdao"
	"encoding/json"
	"github.com/shopspring/decimal"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type FeederInfo struct {
	account  string
	websites []string
}

type ContractFeedingInfo struct {
	symbol                        string
	wallet_api_url                string
	price_feeder_contract_address string
	feeders                       []string
	feedernames                   []string
	inteverl                      int
}

type PriceGrab struct {
	symbolpair          string
	exchangeWebSiteName string
	url                 string
	accessPriceKeys     []string
	muti_factor         int
}

type APriceFeeder struct {
	PriceGrabs           []PriceGrab
	walletPriceFeederApi hdao.PriceFeeder
	account_addr         string
	account_name         string
}

func (pGrab *PriceGrab) grab_price() (string, error) {
	var result = ""
	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", pGrab.url, nil)
	resp, err := client.Do(reqest)
	defer resp.Body.Close()
	if resp.StatusCode != 200 || err != nil {
		return result, err
	}
	body, _ := ioutil.ReadAll(resp.Body)

	var m map[string]interface{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		return result, err
	}
	var price = ""
	for k, v := range pGrab.accessPriceKeys {
		var temp map[string]interface{}
		if k == 0 {
			temp = m[v].(map[string]interface{})
		} else {
			price = temp[v].(string)
		}
	}
	if pGrab.muti_factor != 1 {
		n_price, _ := decimal.NewFromString(price)
		n_price = n_price.Round(10000000)
		result = n_price.String()
	}
	return result, nil
}

func newApriceFeeder(symbolPair string, priceFeeder_contract_address string, accountName string, wallet_api_url string, exchangeWebSites []website) APriceFeeder {
	var priceGrabs []PriceGrab
	var result APriceFeeder
	for _, v := range exchangeWebSites {
		priceGrabs = append(priceGrabs, PriceGrab{symbolpair: symbolPair, exchangeWebSiteName: v.name, url: v.url, accessPriceKeys: v.accessPriceKeys, muti_factor: v.muti_factor})
	}
	result.PriceGrabs = priceGrabs
	walletApi := hdao.HXWalletApi{Name: "priceFeeder_service", Rpc_url: wallet_api_url}
	result.walletPriceFeederApi = hdao.PriceFeeder{Account: accountName, Contract: priceFeeder_contract_address, Wallet_api: walletApi}
	args := []interface{}{accountName}
	ret, _ := walletApi.Rpc_request("get_account_addr", args)
	result.account_addr = ret
	result.account_name = accountName
	return result
}

func (feeder *APriceFeeder) feederPrice() bool {
	maxChangeRatio := 0.099999
	var pricestr = ""
	for _,v:= range feeder.PriceGrabs {
		pricestr,_ = v.grab_price()
		if pricestr != ""{
			break
		}
	}
	if pricestr != "" {
		return false
	}
	r,err:= feeder.walletPriceFeederApi.Get_feedPrices()
	if err != nil{
		return false
	}
	feedPrices := make(map[string]string)
	err = json.Unmarshal([]byte(r), &feedPrices)
	if err != nil {
		return false
	}
	origPriceStr := feedPrices[feeder.account_addr]
	price,err:= strconv.ParseFloat(pricestr,64)
	origPrice,err := strconv.ParseFloat(origPriceStr,64)

	if price> origPrice {
		for {
			if  (price <= origPrice*(1+maxChangeRatio)){
				break
			}
			newPrice := origPrice*(1+maxChangeRatio)
			r ,err= feeder.walletPriceFeederApi.Feed_price(newPrice)
			if err != nil {
				return false
			}
			origPrice = newPrice
		}
	}else{
		for {
			if price >= origPrice*(1- maxChangeRatio) {
				newPrice := origPrice * (1 - maxChangeRatio)
				r ,err= feeder.walletPriceFeederApi.Feed_price(newPrice)
				if err != nil {
					return false
				}
				origPrice = newPrice
			}
		}

	}
	r ,_= feeder.walletPriceFeederApi.Feed_price(price)
	return true
}

type  ContractPriceFeedingRobot struct {
	contractAddr string
	contractFeedingInfo ContractFeedingInfo
	aPriceFeeders []APriceFeeder
	running bool
	startTime time.Time
	stopTime time.Time
	successFeedCount int
	failFeedCount    int
	symbolPair string
	symbolPairExchangeWebSitesInfo string
	webnames string
	wallet_api_url string

}

func (robot * ContractPriceFeedingRobot) run()  {
	robot.running = true
	var interval  = robot.contractFeedingInfo.inteverl
	isContinue := true
	robot.startTime = time.Now()
	feedersCount := len(robot.aPriceFeeders)
	rounds := 0
	for {
		if (isContinue && robot.running) == false{
			return
		}

		for _,v := range robot.aPriceFeeders {
			r := v.feederPrice()
			if r== false{
				robot.failFeedCount += 1
			}else {
				robot.successFeedCount += 1
			}
		}
		rounds += 1
		time.Sleep( time.Duration(interval) * time.Second)

		if (robot.failFeedCount >5 && robot.failFeedCount /(robot.failFeedCount + robot.successFeedCount) >= 1/feedersCount) {
			isContinue = false
		}
	}
	if robot.running == true {
		robot.running = false
	}
	robot.stopTime = time.Now()
}

func (robot *ContractPriceFeedingRobot)start()  {
	 robot.run()
}

func (robot *ContractPriceFeedingRobot)stop()  {
	robot.running = false
}


type PriceFeedingFactory struct {
	robots []APriceFeeder
	json_config map[string]interface{}
	robot_config_filepath string
}
func (factory* PriceFeedingFactory) loadConfigFile() error {
	viper.SetConfigFile(factory.robot_config_filepath)
	err := viper.ReadInConfig()
	if err !=nil {
		return err
	}
	factory.json_config["feedingContractsInfo"] = viper.Get("feedingContractsInfo")
	factory.json_config["exchangeWebSitesInfo"] = viper.Get("exchangeWebSitesInfo")

	return nil
}


