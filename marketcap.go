package okexmarket

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"net/url"
	"strconv"
	"io"
	"crypto/md5"
	"strings"

)
var (
	baseURL    = "https://www.okex.com/api/v1"
	//proxyAddr = "http://192.168.31.185:8123"
	proxyAddr = "http://127.0.0.1:1080"

)
func getTransportFieldURL(proxy_addr *string) (transport *http.Transport) {
	url_i := url.URL{}
	url_proxy, _ := url_i.Parse(*proxy_addr)
	transport = &http.Transport{Proxy: http.ProxyURL(url_proxy)}
	return transport
}
//func getTransportFromEnvironment() (transport *http.Transport) {
//	transport = &http.Transport{Proxy : http.ProxyFromEnvironment}
//	return transport
//
//}

// GetGlobalMarketData get information about the global market data of the cryptocurrencies
func  GetUsdtBtcTickerData() (TickerData, error) {
	url := fmt.Sprintf(baseURL + "/ticker.do?symbol=btc_usdt")
	resp, err := makeReq(url)
	var data TickerData

	err = json.Unmarshal(resp, &data)
	fmt.Printf("%v \n",string(resp))
	if err != nil {
		return TickerData{}, err
	}

	return data, nil
}

// GetUsdtBtcDepthData get information about the global market data of the cryptocurrencies
func GetUsdtBtcDepthData(size int64) (DepthData, error) {
	url := fmt.Sprintf(baseURL + "/depth.do?symbol=btc_usdt&size="+strconv.FormatInt(size,10) )
	fmt.Printf("%v \n",url)
	resp, err := makeReq(url)
	var data DepthData

	err = json.Unmarshal(resp, &data)
	fmt.Printf("%v \n",string(resp))
	if err != nil {
		return DepthData{}, err
	}

	return data, nil
}


// GetUsdtBtcTradeData get information about the global market data of the cryptocurrencies
func GetUsdtBtcTradeData(since int64) ([]TradeData, error) {
	url := fmt.Sprintf(baseURL + "/trades.do?symbol=btc_usdt&size="+strconv.FormatInt(since,10) )
	fmt.Printf("%v \n",url)
	resp, err := makeReq(url)
	var data []TradeData

	err = json.Unmarshal(resp, &data)
	fmt.Printf("%v \n",string(resp))
	if err != nil {
		return []TradeData{}, err
	}


	return data, nil
}
// GetUserInfo get information about the global market data of the cryptocurrencies
func (c Context) GetUserInfo() (map[string]interface{}, error) {
	url_ := fmt.Sprintf(baseURL + "/userinfo.do" )
	fmt.Printf("%v \n",url_)

	v := url.Values{}
	v.Add("api_key",c.ApiKey)
	v.Add("sign",md5_32(v.Encode()+"&secret_key="+c.ApiSecret))
	resp, err := makePostFormReq(url_,v)
	var data interface{}
	fmt.Printf("%v \n",string(resp))
	err = json.Unmarshal(resp, &data)
	msgMap := data.(map[string]interface{})
	if err != nil {
		return nil, err
	}

	return msgMap, nil
}

// makeReq HTTP request helper
func makePostFormReq(url string,body url.Values) ([]byte, error) {
	fmt.Printf("body %v \n",body.Encode())
	req, err := http.NewRequest("POST",url,strings.NewReader(body.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(body.Encode())))
	if err != nil {
		return nil, err
	}
	resp, err := doReq(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
func md5_32(data string) string {
	fmt.Printf("md5_32 data %v \n",data)
	h := md5.New()
	io.WriteString(h, data)
	sign := h.Sum(nil)
	return strings.ToUpper( fmt.Sprintf("%x", sign))
}

// makeReq HTTP request helper
func makeReq(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := doReq(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}


// doReq HTTP client
func doReq(req *http.Request) ([]byte, error) {
	transport := getTransportFieldURL(&proxyAddr)

	client := &http.Client{Transport : transport}
	//client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}