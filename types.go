package okexmarketcap

// Interface interface
type Interface interface {
	GetUsdtBtcTickerData() (TickerDate, start int64, end int64)

}
// TickerData struct
type TickerData struct {
	Date  string `json:"date"`
	Ticker Ticker `json:"ticker"`
}
type Ticker struct{
	High string  `json:"high"`
	Vol string `json:"vol"`
	Last string `json:"last"`
	Low string `json:"low"`
	Buy string `json:"buy"`
	Sell string `json:"sell"`
}