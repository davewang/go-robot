package okexmarketcap

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"net/url"
)

var (
	baseURL               = "https://www.okex.com/api/v1"
	proxyAddr = "http://192.168.31.185:8123"
)
func getTransportFieldURL(proxy_addr *string) (transport *http.Transport) {
	url_i := url.URL{}
	url_proxy, _ := url_i.Parse(*proxy_addr)
	transport = &http.Transport{Proxy: http.ProxyURL(url_proxy)}
	return transport
}
func getTransportFromEnvironment() (transport *http.Transport) {
	transport = &http.Transport{Proxy : http.ProxyFromEnvironment}
	return transport

}

// GetGlobalMarketData get information about the global market data of the cryptocurrencies
func GetUsdtBtcTickerData() (TickerData, error) {
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