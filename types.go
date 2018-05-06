package okexmarket


// core interface
type Context struct {
	ApiKey string
	ApiSecret string

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

type DepthData struct{
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
}


type TradeData struct{
	Date int64  `json:"date"`
	DateMS int64 `json:"date_ms"`
	Price float64 `json:"price"`
	Amount float64 `json:"amount"`
	TId int64 `json:"tid"`
	Type string `json:"type"`
}

