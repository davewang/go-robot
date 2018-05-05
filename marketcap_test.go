package okexmarketcap

import (
	"testing"
	"fmt"
)

func TestUsdtBtcTickerData(t *testing.T) {
	data,err := GetUsdtBtcTickerData()
	if err != nil {
		fmt.Printf(err.Error())
		t.FailNow()
	}
	fmt.Printf("data:%v",data)
}
