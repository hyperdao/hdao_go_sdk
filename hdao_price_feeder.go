package hdao_sdk4go

import (
	"strconv"
	"strings"
)

var gasPrice = 0.0001
var gasLimit = 1000000

type PriceFeeder struct {
	Account    string
	Contract   string
	Wallet_api *HXWalletApi
}

func NewPriceFeeder(account string, contract string, wallet_api *HXWalletApi) *PriceFeeder {
	r := PriceFeeder{
		Account:    account,
		Contract:   contract,
		Wallet_api: wallet_api,
	}
	return &r
}

// apis
func (priceFeeder *PriceFeeder) Init_config(baseAsset string, quoteAsset string, init_price float64, maxChangeRatio float64) (string, error) {
	apiargs := []string{baseAsset, quoteAsset, strconv.FormatFloat(init_price, 'f', -1, 64), strconv.FormatFloat(maxChangeRatio, 'f', -1, 64)}
	args := []interface{}{priceFeeder.Account, gasPrice, gasLimit, priceFeeder.Contract, "init_config", strings.Join(apiargs, ",")}
	return priceFeeder.Wallet_api.Rpc_request("invoke_contract", args)
}

func (priceFeeder *PriceFeeder) Add_feeder(addr string) (string, error) {
	args := []interface{}{priceFeeder.Account, gasPrice, gasLimit, priceFeeder.Contract, "add_feeder", addr}
	return priceFeeder.Wallet_api.Rpc_request("invoke_contract", args)
}

func (priceFeeder *PriceFeeder) Remove_feeder(addr string) (string, error) {
	args := []interface{}{priceFeeder.Account, gasPrice, gasLimit, priceFeeder.Contract, "remove_feeder", addr}
	return priceFeeder.Wallet_api.Rpc_request("invoke_contract", args)
}

func (priceFeeder *PriceFeeder) Change_owner(addr string) (string, error) {
	args := []interface{}{priceFeeder.Account, gasPrice, gasLimit, priceFeeder.Contract, "change_owner", addr}
	return priceFeeder.Wallet_api.Rpc_request("invoke_contract", args)
}

func (priceFeeder *PriceFeeder) Feed_price(price float64) (string, error) {
	args := []interface{}{priceFeeder.Account, gasPrice, gasLimit, priceFeeder.Contract, "feed_price", strconv.FormatFloat(price, 'f', 8, 64)}
	return priceFeeder.Wallet_api.Rpc_request("invoke_contract", args)
}

//offline api
func (priceFeeder *PriceFeeder) Get_price() (string, error) {
	args := []interface{}{priceFeeder.Account, priceFeeder.Contract, "getPrice", ""}
	return priceFeeder.Wallet_api.Rpc_request("invoke_contract_offline", args)
}

func (priceFeeder *PriceFeeder) Get_owner() (string, error) {
	args := []interface{}{priceFeeder.Account, priceFeeder.Contract, "owner", ""}
	return priceFeeder.Wallet_api.Rpc_request("invoke_contract_offline", args)
}

func (priceFeeder *PriceFeeder) Get_state() (string, error) {
	args := []interface{}{priceFeeder.Account, priceFeeder.Contract, "state", ""}
	return priceFeeder.Wallet_api.Rpc_request("invoke_contract_offline", args)
}

func (priceFeeder *PriceFeeder) Get_feeders() (string, error) {
	args := []interface{}{priceFeeder.Account, priceFeeder.Contract, "feeders", ""}
	return priceFeeder.Wallet_api.Rpc_request("invoke_contract_offline", args)
}

func (priceFeeder *PriceFeeder) Get_feedPrices() (string, error) {
	args := []interface{}{priceFeeder.Account, priceFeeder.Contract, "feedPrices", ""}
	return priceFeeder.Wallet_api.Rpc_request("invoke_contract_offline", args)
}

func (priceFeeder *PriceFeeder) Get_base_asset() (string, error) {
	args := []interface{}{priceFeeder.Account, priceFeeder.Contract, "baseAsset", ""}
	return priceFeeder.Wallet_api.Rpc_request("invoke_contract_offline", args)
}

func (priceFeeder *PriceFeeder) Get_quote_asset() (string, error) {
	args := []interface{}{priceFeeder.Account, priceFeeder.Contract, "quotaAsset", ""}
	return priceFeeder.Wallet_api.Rpc_request("invoke_contract_offline", args)
}
