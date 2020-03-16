// hdao_sdk4go project hdao_sdk4go.go

package hdao_sdk4go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gohouse/gorose"
	_ "github.com/mattn/go-sqlite3"
)

var err error
var engin *gorose.Engin

func init() {
	content, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	var conf map[string]string
	if err = json.Unmarshal(content, &conf); err != nil {
		fmt.Println(err)
		return
	}
	cdcs_db_driver := conf["cdcs_db_driver"]
	cdcs_db_path := conf["cdcs_db_path"]
	engin, err = gorose.Open(&gorose.Config{Driver: cdcs_db_driver, Dsn: cdcs_db_path})

}

func DB() gorose.IOrm {
	return engin.NewOrm()
}

func Query_cdc_by_address(ownerAddress string) (CdcTable, error) {
	db := DB()
	var cdc CdcTable
	//.Fields("chain_id,cdc_contract_address,contract_register_block_num,end_block_num,end_block_id")
	err := db.Table(&cdc).Where("owner", ownerAddress).Select()
	//fmt.Println(db.LastSql())
	return cdc, err
}

//state=-1 表示查询不限state  ownerAddress="" 表示查询不限owner
func Query_cdcs(start int, limit int, state int, ownerAddress string) (*[]CdcTable, error) {
	db := DB()
	var cdcs []CdcTable

	r := db.Table(&cdcs)

	if state >= 0 {
		r = db.Table(&cdcs).Where("state", state)
	}

	if ownerAddress != "" {
		r = r.Where("owner", ownerAddress)
	}
	err := r.Limit(limit).Offset(start).Select()
	if err != nil {
	    fmt.Println(err)
	}
	return &cdcs, err
}

func Query_cdc_by_id(cdcId string) (CdcTable, error) {
	db := DB()
	var cdc CdcTable
	err := db.Table(&cdc).Where("cdcId", cdcId).Select()
	if err != nil {
	    fmt.Println(err)
	}
	return cdc, err
}

func Query_supply(startblocknum int, endblocknum int) (*[]StableTokenSupplyHistoryTable, error) {
	db := DB()
	var supplys []StableTokenSupplyHistoryTable
	err := db.Table(&supplys).Where("block_number", ">=", startblocknum).Where("block_number", "<=", endblocknum).Select()
	if err != nil {
	    fmt.Println(err)
	}
	return &supplys, err
}
