package openinsider

import "time"

type Disclosure struct {
	X           string
	FilingDate  time.Time
	TradeDate   time.Time
	Ticker      string
	CompanyName string
	Industry    string
	Ins         uint8
	TradeType   string
	Price       float64
	Qty         uint32
	DeltaOwn    float64
	Value       float64
	Title       string
}

type Options struct {
	Websocket bool
	Ttl       time.Duration
}

type Filter struct {
}
