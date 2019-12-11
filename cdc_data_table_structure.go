// hdao_sdk4go project hdao_sdk4go.go

package hdao_sdk4go

//import "math/big"

//////////////////////////////////////////////////////////////

type CdcChainTable struct {
	Chain_id                    string `gorose:"chain_id"`
	Cdc_contract_address        string `gorose:"cdc_contract_address"`
	Contract_register_block_num int    `gorose:"contract_register_block_num"`
	End_block_num               int    `gorose:"end_block_num"`
	End_block_id                string `gorose:"end_block_id"`
}

func (u *CdcChainTable) TableName() string {
	return "cdc_chain"
}

////////////
type CdcTable struct {
	CdcId string `gorose:"cdcId"`
	State int    `gorose:"state"`
	//StabilityFee      int64  `gorose:"stabilityFee"`
	CollateralAmount  int64  `gorose:"collateralAmount"`
	StableTokenAmount int64  `gorose:"stableTokenAmount"`
	Owner             string `gorose:"owner"`
	Liquidator        string `gorose:"liquidator"`
	Block_number      int    `gorose:"block_number"`
	Last_event_id     string `gorose:"last_event_id"`
}

func (u *CdcTable) TableName() string {
	return "cdcs"
}

////
type CdcEventTable struct {
	Event_id int    `gorose:"event_id"`
	Tx_id    string `gorose:"tx_id"`
	op       string `gorose:"op"`
}

func (u *CdcEventTable) TableName() string {
	return "CdcEvent"
}

////
type StableTokenSupplyHistoryTable struct {
	Block_number int   `gorose:"block_number"`
	Supply       int64 `gorose:"supply"`
}

func (u *StableTokenSupplyHistoryTable) TableName() string {
	return "StableTokenSupplyHistory"
}

type EventOpenCdcTable struct {
	Event_id string `gorose:"event_id"`
	CdcId    string `gorose:"cdcId"`
	Owner    string `gorose:"owner"`
	////
	SecSinceEpoch     int   `gorose:"secSinceEpoch"`
	CollateralAmount  int64 `gorose:"collateralAmount"`
	StableTokenAmount int64 `gorose:"stableTokenAmount"`
	Block_number      int   `gorose:"block_number"`
}

func (u *EventOpenCdcTable) TableName() string {
	return "OpenCdc"
}

type EventLiquidateTable struct {
	Event_id                string `gorose:"event_id"`
	CdcId                   string `gorose:"cdcId"`
	Owner                   string `gorose:"owner"`
	SecSinceEpoch           int    `gorose:"secSinceEpoch"`
	CollateralAmount        int64  `gorose:"collateralAmount"`
	StableTokenAmount       int64  `gorose:"stableTokenAmount"`
	CurPrice                string `gorose:"curPrice"`
	IsBadDebt               bool   `gorose:"isBadDebt"`
	Liquidator              string `gorose:"liquidator"`
	AuctionPrice            string `gorose:"auctionPrice"`
	ReturnAmount            int64  `gorose:"returnAmount"`
	PenaltyAmount           int64  `gorose:"penaltyAmount"`
	IsNeedLiquidation       bool   `gorose:"isNeedLiquidation"`
	StabilityFee            int64  `gorose:"stabilityFee"`
	RepayStableTokenAmount  int64  `gorose:"repayStableTokenAmount"`
	AuctionCollateralAmount int64  `gorose:"auctionCollateralAmount"`
	Block_number            int    `gorose:"block_number"`
}

func (u *EventLiquidateTable) TableName() string {
	return "Liquidate"
}

//
type EventAddCollateralTable struct {
	Event_id     string `gorose:"event_id"`
	CdcId        string `gorose:"cdcId"`
	AddAmount    int64  `gorose:"addAmount"`
	Block_number int    `gorose:"block_number"`
}

func (u *EventAddCollateralTable) TableName() string {
	return "AddCollateral"
}

type EventCloseCdcTable struct {
	Event_id          string `gorose:"event_id"`
	CdcId             string `gorose:"cdcId"`
	Owner             string `gorose:"owner"`
	SecSinceEpoch     int    `gorose:"secSinceEpoch"`
	StabilityFee      int64  `gorose:"stabilityFee"`
	CollateralAmount  int64  `gorose:"collateralAmount"`
	StableTokenAmount int64  `gorose:"stableTokenAmount"`
	Block_number      int    `gorose:"block_number"`
}

func (u *EventCloseCdcTable) TableName() string {
	return "CloseCdc"
}

type EventExpandLoanTable struct {
	Event_id         string `gorose:"event_id"`
	CdcId            string `gorose:"cdcId"`
	From_address     string `gorose:"from_address"`
	RepayFee         int64  `gorose:"repayFee"`
	ExpandLoanAmount int64  `gorose:"expandLoanAmount"`
	RealGotAmount    int64  `gorose:"realGotAmount"`
	Block_number     int    `gorose:"block_number"`
}

func (u *EventExpandLoanTable) TableName() string {
	return "ExpandLoan"
}

type EventWidrawCollateralTable struct {
	Event_id               string `gorose:"event_id"`
	CdcId                  string `gorose:"cdcId"`
	From_address           string `gorose:"from_address"`
	WidrawCollateralAmount int64  `gorose:"widrawCollateralAmount"`
	Block_number           int    `gorose:"block_number"`
}

func (u *EventWidrawCollateralTable) TableName() string {
	return "WidrawCollateral"
}

type EventPayBackTable struct {
	Event_id          string `gorose:"event_id"`
	CdcId             string `gorose:"cdcId"`
	From_address      string `gorose:"from_address"`
	Fee               int64  `gorose:"fee"`
	RepayPrincipal    int64  `gorose:"repayPrincipal"`
	PayBackAmount     int64  `gorose:"payBackAmount"`
	RealPayBackAmount int64  `gorose:"realPayBackAmount"`
	Block_number      int    `gorose:"block_number"`
}

func (u *EventPayBackTable) TableName() string {
	return "PayBack"
}

type EventTransferCdcTable struct {
	Event_id     string `gorose:"event_id"`
	CdcId        string `gorose:"cdcId"`
	From_address string `gorose:"from_address"`
	To_address   string `gorose:"to_address"`
	Block_number int    `gorose:"block_number"`
}

func (u *EventTransferCdcTable) TableName() string {
	return "TransferCdc"
}
