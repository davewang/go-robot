package okexmarket

import (
	"testing"
	"fmt"
	"time"
)



func TestGetUsdtBtcTickerData(t *testing.T) {
	data,err := GetUsdtBtcTickerData()
	if err != nil {
		fmt.Printf(err.Error())
		t.FailNow()
	}
	fmt.Printf("data:%v\n",data)
}

func TestGetUsdtBtcDepthData(t *testing.T) {
	data,err := GetUsdtBtcDepthData(10)
	if err != nil {
		fmt.Printf(err.Error())
		t.FailNow()
	}
	fmt.Printf("data:%v\n",data)
}

func TestGetUsdtBtcTradeData(t *testing.T) {
	data,err := GetUsdtBtcTradeData(time.Now().Unix())
	if err != nil {
		fmt.Printf(err.Error())
		t.FailNow()
	}
	fmt.Printf("data:%v\n",data)
}

func TestGetUserInfo(t *testing.T) {


	api_key := "api_key"
	api_secret := "api_secret"
	c := Context{ApiKey:api_key,ApiSecret:api_secret}

	data,err := c.GetUserInfo()
	if err != nil {
		fmt.Printf(err.Error())
		t.FailNow()
	}
	fmt.Printf("data:%v\n",data)


}