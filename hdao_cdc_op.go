package hdao_sdk4go

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

//var gasPrice = 0.0001
//var gasLimit = 1000000

type CDCOperation struct {
	Account    string
	Contract   string
	Wallet_api *HXWalletApi
	Asset      string
	Precision  int
}

func NewCDCOperation(account string, contract string, wallet_api *HXWalletApi) (*CDCOperation, error) {
	r := CDCOperation{
		Account:    account,
		Contract:   contract,
		Wallet_api: wallet_api,
	}
	result, err := r.Get_Info()
	if err != nil {
		return &r, err
	}
	var m map[string]interface{}
	json.Unmarshal([]byte(result), &m)
	var ok = false
	r.Asset, ok = m["collateralAsset"].(string)
	if !ok {
		return &r, errors.New("Get_Info wrong")
	}

	if r.Asset == "HX" {
		r.Precision = 100000
	} else {
		r.Precision = 100000000
	}
	return &r, nil
}

var stableCoinPrecision = 100000000

//apis

func (cdcOperation *CDCOperation) Init_config(collateralAsset string, collateralizationRatio float64, annualStabilityFee float64, liquidationRatio float64, liquidationPenalty float64, liquidationDiscount float64, priceFeederAddr string, stableCoinAddr string, proxyAddr string) (string, error) {
	apiargs := []string{collateralAsset, strconv.FormatFloat(collateralizationRatio, 'f', -1, 64), strconv.FormatFloat(annualStabilityFee, 'f', -1, 64), strconv.FormatFloat(liquidationRatio, 'f', -1, 64), strconv.FormatFloat(liquidationPenalty, 'f', -1, 64), strconv.FormatFloat(liquidationDiscount, 'f', -1, 64), priceFeederAddr, stableCoinAddr, proxyAddr}
	args := []interface{}{cdcOperation.Account, gasPrice, gasLimit, cdcOperation.Contract, "init_config", strings.Join(apiargs, ",")}
	return cdcOperation.Wallet_api.Rpc_request("invoke_contract", args)
}

func (cdcOperation *CDCOperation) Open_cdc(collateralAmount float64, stableCoinAmount float64) (string, error) {
	if collateralAmount < 0.00000001 {
		return "", errors.New("collateralAmount must > 0")
	}

	apiargstr := "openCdc"
	if stableCoinAmount >= 0.00000001 {
		apiargstr = apiargstr + "," + strconv.FormatFloat(stableCoinAmount*float64(stableCoinPrecision), 'f', 0, 64)
	}
	return cdcOperation.Wallet_api.Rpc_request("transfer_to_contract", []interface{}{cdcOperation.Account, cdcOperation.Contract, collateralAmount, cdcOperation.Asset, apiargstr, gasPrice, gasLimit, true})
}

func (cdcOperation *CDCOperation) Add_collateral(cdc_id string, collateralAmount float64) (string, error) {
	if collateralAmount < 0.00000001 {
		return "", errors.New("collateralAmount must >= 0.00000001")
	}
	apiargstr := "addCollateral," + cdc_id
	return cdcOperation.Wallet_api.Rpc_request("transfer_to_contract", []interface{}{cdcOperation.Account, cdcOperation.Contract, collateralAmount, cdcOperation.Asset, apiargstr, gasPrice, gasLimit, true})
}

func (cdcOperation *CDCOperation) Generate_stable_coin(cdc_id string, stableCoinAmount float64) (string, error) {
	if stableCoinAmount < 0.00000001 {
		return "", errors.New("stableCoinAmount must >= 0.00000001")
	}
	apiargs := []string{cdc_id, strconv.FormatFloat(stableCoinAmount*float64(stableCoinPrecision), 'f', 0, 64)}
	args := []interface{}{cdcOperation.Account, gasPrice, gasLimit, cdcOperation.Contract, "expandLoan", strings.Join(apiargs, ",")}
	return cdcOperation.Wallet_api.Rpc_request("invoke_contract", args)
}

func (cdcOperation *CDCOperation) Withdraw_collateral(cdc_id string, amount float64) (string, error) {
	if amount < 0.00000001 {
		return "", errors.New("amount must >= 0.00000001")
	}
	apiargs := []string{cdc_id, strconv.FormatFloat(amount*float64(cdcOperation.Precision), 'f', 0, 64)}
	args := []interface{}{cdcOperation.Account, gasPrice, gasLimit, cdcOperation.Contract, "widrawCollateral", strings.Join(apiargs, ",")}
	return cdcOperation.Wallet_api.Rpc_request("invoke_contract", args)
}

func (cdcOperation *CDCOperation) Transfer_cdc(cdc_id string, addr string) (string, error) {

	apiargs := []string{cdc_id, addr}
	args := []interface{}{cdcOperation.Account, gasPrice, gasLimit, cdcOperation.Contract, "transferCdc", strings.Join(apiargs, ",")}
	return cdcOperation.Wallet_api.Rpc_request("invoke_contract", args)
}

func (cdcOperation *CDCOperation) Pay_back(cdc_id string, stableCoinAmount float64) (string, error) {
	if stableCoinAmount < 0.00000001 {
		return "", errors.New("stableCoinAmount must >= 0.00000001")
	}
	apiargs := []string{cdc_id, strconv.FormatFloat(stableCoinAmount*float64(stableCoinPrecision), 'f', 0, 64)}
	args := []interface{}{cdcOperation.Account, gasPrice, gasLimit, cdcOperation.Contract, "payBack", strings.Join(apiargs, ",")}
	return cdcOperation.Wallet_api.Rpc_request("invoke_contract", args)
}

func (cdcOperation *CDCOperation) Liquidate(cdc_id string, stableCoinAmount float64, collateralAmount float64) (string, error) {
	if (stableCoinAmount < 0.00000001) || (collateralAmount < 0.00000001) {
		return "", errors.New("amount must >= 0.00000001")
	}
	apiargs := []string{cdc_id, strconv.FormatFloat(stableCoinAmount*float64(stableCoinPrecision), 'f', 0, 64), strconv.FormatFloat(collateralAmount*float64(cdcOperation.Precision), 'f', 0, 64)}
	args := []interface{}{cdcOperation.Account, gasPrice, gasLimit, cdcOperation.Contract, "liquidate", strings.Join(apiargs, ",")}
	return cdcOperation.Wallet_api.Rpc_request("invoke_contract", args)
}

func (cdcOperation *CDCOperation) Close_cdc(cdc_id string) (string, error) {
	args := []interface{}{cdcOperation.Account, gasPrice, gasLimit, cdcOperation.Contract, "closeCdc", cdc_id}
	return cdcOperation.Wallet_api.Rpc_request("invoke_contract", args)
}

func (cdcOperation *CDCOperation) Set_annual_stability_fee(annual_stability_fee float64) (string, error) {
	args := []interface{}{cdcOperation.Account, gasPrice, gasLimit, cdcOperation.Contract, "setAnnualStabilityFee", strconv.FormatFloat(annual_stability_fee, 'f', -1, 64)}
	return cdcOperation.Wallet_api.Rpc_request("invoke_contract", args)
}

func (cdcOperation *CDCOperation) Set_liquidation_ratio(ratio float64) (string, error) {
	args := []interface{}{cdcOperation.Account, gasPrice, gasLimit, cdcOperation.Contract, "setLiquidationRatio", strconv.FormatFloat(ratio, 'f', -1, 64)}
	return cdcOperation.Wallet_api.Rpc_request("invoke_contract", args)
}

func (cdcOperation *CDCOperation) Set_liquidation_penalty(penalty float64) (string, error) {
	args := []interface{}{cdcOperation.Account, gasPrice, gasLimit, cdcOperation.Contract, "setLiquidationPenalty", strconv.FormatFloat(penalty, 'f', -1, 64)}
	return cdcOperation.Wallet_api.Rpc_request("invoke_contract", args)
}

func (cdcOperation *CDCOperation) Set_liquidation_discount(discount float64) (string, error) {
	args := []interface{}{cdcOperation.Account, gasPrice, gasLimit, cdcOperation.Contract, "setLiquidationDiscount", strconv.FormatFloat(discount, 'f', -1, 64)}
	return cdcOperation.Wallet_api.Rpc_request("invoke_contract", args)
}

func (cdcOperation *CDCOperation) Set_price_feeder_addr(addr string) (string, error) {
	args := []interface{}{cdcOperation.Account, gasPrice, gasLimit, cdcOperation.Contract, "setPriceFeederAddr", addr}
	return cdcOperation.Wallet_api.Rpc_request("invoke_contract", args)
}

func (cdcOperation *CDCOperation) Change_admin(addr string) (string, error) {
	args := []interface{}{cdcOperation.Account, gasPrice, gasLimit, cdcOperation.Contract, "changeAdmin", addr}
	return cdcOperation.Wallet_api.Rpc_request("invoke_contract", args)
}

func (cdcOperation *CDCOperation) Global_liquidate() (string, error) {
	args := []interface{}{cdcOperation.Account, gasPrice, gasLimit, cdcOperation.Contract, "globalLiquidate", ""}
	return cdcOperation.Wallet_api.Rpc_request("invoke_contract", args)
}

func (cdcOperation *CDCOperation) Take_back_collateral_by_cdc(cdc_id string) (string, error) {
	args := []interface{}{cdcOperation.Account, gasPrice, gasLimit, cdcOperation.Contract, "takeBackCollateralByCdc", cdc_id}
	return cdcOperation.Wallet_api.Rpc_request("invoke_contract", args)
}

func (cdcOperation *CDCOperation) Take_back_collateral_by_token(cdc_id string) (string, error) {
	args := []interface{}{cdcOperation.Account, gasPrice, gasLimit, cdcOperation.Contract, "takeBackCollateralByToken", cdc_id}
	return cdcOperation.Wallet_api.Rpc_request("invoke_contract", args)
}

func (cdcOperation *CDCOperation) Close_contract() (string, error) {
	args := []interface{}{cdcOperation.Account, gasPrice, gasLimit, cdcOperation.Contract, "closeContract", ""}
	return cdcOperation.Wallet_api.Rpc_request("invoke_contract", args)
}

//offline apis
func (cdcOperation *CDCOperation) Get_Info() (string, error) {
	args := []interface{}{cdcOperation.Account, cdcOperation.Contract, "getInfo", ""}
	return cdcOperation.Wallet_api.Rpc_request("invoke_contract_offline", args)
}

func (cdcOperation *CDCOperation) Get_cdc(cdc_id string) (string, error) {
	args := []interface{}{cdcOperation.Account, cdcOperation.Contract, "getCdc", cdc_id}
	return cdcOperation.Wallet_api.Rpc_request("invoke_contract_offline", args)
}

func (cdcOperation *CDCOperation) Get_liquidable_info(cdc_id string) (string, error) {
	args := []interface{}{cdcOperation.Account, cdcOperation.Contract, "getLiquidableInfo", cdc_id}
	return cdcOperation.Wallet_api.Rpc_request("invoke_contract_offline", args)
}
