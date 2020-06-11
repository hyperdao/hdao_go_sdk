package hdao_sdk4go

import (
	"testing"

	"fmt"
)

func TestQuery(t *testing.T) {
	r, err := Query_cdc_by_address("HXNYM7NT7nbNZPdHjzXf2bkDR53riKxV9kgh")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("cdc:")
	fmt.Println(r)

	r2, err := Query_cdcs(0, 10, -1, "")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("cdcs:")
	fmt.Println(r2)
}

func TestPriceFeeder(t *testing.T) {

	walletApi := NewHXWalletApi("senator0", "", "", "http://192.168.1.121:30088/")
	priceFeeder := NewPriceFeeder("april", "HXCGba6bUaGeBtUQRGpHUePHVXzF1ygMAxR1", walletApi)

	r, err := priceFeeder.Get_price()
	fmt.Println("prrice:")
	fmt.Println(r)

	priceFeeder.Init_config("a", "b", 1.1234567890, 0.123)
	priceFeeder.Feed_price(1.1234567890)

	r, err = priceFeeder.Get_base_asset()
	if err != nil {
		t.Error(err.Error())
	}
	r, err = priceFeeder.Get_feeders()
	if err != nil {
		t.Error(err.Error())
	}

	r, err = priceFeeder.Get_feedPrices()
	if err != nil {
		t.Error(err.Error())
	}

	r, err = priceFeeder.Get_state()
	if err != nil {
		t.Error(err.Error())
	}

	r, err = priceFeeder.Get_quote_asset()
	if err != nil {
		t.Error(err.Error())
	}

}

func TestCdcOperation(t *testing.T) {
	walletApi := NewHXWalletApi("senator0", "", "", "http://192.168.1.121:30088/")
	cdcOperation, err := NewCDCOperation("april", "HXCSSGDHqaJDLto13BSZpAbrZoJf4RrGCtks", walletApi)
	if err != nil {
		t.Error(err.Error())
	}
	/*
		result, err := cdcOperation.Open_cdc(0.000123456789, 0.001234567890123)
		if err != nil {
			t.Error(err.Error())
		}
		fmt.Println(result)
	*/
	result, err := cdcOperation.Get_cdc("8e3fa713d9ce4386a9847c872e6fd80cf424b1fc")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(result)
}

func TestQuerySupplys(t *testing.T) {
	r, err := Query_supply(969189, 999189)
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println("supplies:")
	fmt.Println(r)
}
