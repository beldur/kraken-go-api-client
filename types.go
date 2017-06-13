package krakenapi

import (
	"fmt"
	"time"
)

// Those constants are used to define the Kraken pairs
const (
	DASHEUR  = "DASHEUR"
	DASHUSD  = "DASHUSD"
	DASHXBT  = "DASHXBT"
	GNOETH   = "GNOETH"
	GNOEUR   = "GNOEUR"
	GNOUSD   = "GNOUSD"
	GNOXBT   = "GNOXBT"
	USDTZUSD = "USDTZUSD"
	XETCXETH = "XETCXETH"
	XETCXXBT = "XETCXXBT"
	XETCZEUR = "XETCZEUR"
	XETCXUSD = "XETCXUSD"
	XETHXXBT = "XETHXXBT"
	XETHZCAD = "XETHZCAD"
	XETHZEUR = "XETHZEUR"
	XETHZGBP = "XETHZGBP"
	XETHZJPY = "XETHZJPY"
	XETHZUSD = "XETHZUSD"
	XICNXETH = "XICNXETH"
	XICNXXBT = "XICNXXBT"
	XLTCXXBT = "XLTCXXBT"
	XLTCZEUR = "XLTCZEUR"
	XLTCZUSD = "XLTCZUSD"
	XMLNXETH = "XMLNXETH"
	XMLNXXBT = "XMLNXXBT"
	XREPXETH = "XREPXETH"
	XREPXXBT = "XREPXXBT"
	XREPZEUR = "XREPZEUR"
	XREPZUSD = "XREPZUSD"
	XXBTZCAD = "XXBTZCAD"
	XXBTZEUR = "XXBTZEUR"
	XXBTZGBP = "XXBTZGBP"
	XXBTZJPY = "XXBTZJPY"
	XXBTZUSD = "XXBTZUSD"
	XXDGXXBT = "XXDGXXBT"
	XXLMXXBT = "XXLMXXBT"
	XXLMZEUR = "XXLMZEUR"
	XXLMZUSD = "XXLMZUSD"
	XXMRXXBT = "XXMRXXBT"
	XXMRZEUR = "XXMRZEUR"
	XXMRZUSD = "XXMRZUSD"
	XXRPXXBT = "XXRPXXBT"
	XXRPZCAD = "XXRPZCAD"
	XXRPZEUR = "XXRPZEUR"
	XXRPZJPY = "XXRPZJPY"
	XXRPZUSD = "XXRPZUSD"
	XZECXXBT = "XZECXXBT"
	XZECZEUR = "XZECZEUR"
	XZECZUSD = "XZECZUSD"
)

// It's for the type of price
const (
	BUY    = "b"
	SELL   = "s"
	MARKET = "m"
	LIMIT  = "l"
)

// KrakenResponse wraps the Kraken API JSON response
type KrakenResponse struct {
	Error  []string    `json:"error"`
	Result interface{} `json:"result"`
}

// TimeResponse represents the server's time
type TimeResponse struct {
	// Unix timestamp
	Unixtime int64
	// RFC 1123 time format
	Rfc1123 string
}

// AssetPairsResponse includes asset pair informations
type AssetPairsResponse struct {
	DASHEUR  AssetPairInfo
	DASHUSD  AssetPairInfo
	DASHXBT  AssetPairInfo
	GNOETH   AssetPairInfo
	GNOEUR   AssetPairInfo
	GNOUSD   AssetPairInfo
	GNOXBT   AssetPairInfo
	USDTZUSD AssetPairInfo
	XETCXETH AssetPairInfo
	XETCXXBT AssetPairInfo
	XETCZEUR AssetPairInfo
	XETCXUSD AssetPairInfo
	XETHXXBT AssetPairInfo
	XETHZCAD AssetPairInfo
	XETHZEUR AssetPairInfo
	XETHZGBP AssetPairInfo
	XETHZJPY AssetPairInfo
	XETHZUSD AssetPairInfo
	XICNXETH AssetPairInfo
	XICNXXBT AssetPairInfo
	XLTCXXBT AssetPairInfo
	XLTCZEUR AssetPairInfo
	XLTCZUSD AssetPairInfo
	XMLNXETH AssetPairInfo
	XMLNXXBT AssetPairInfo
	XREPXETH AssetPairInfo
	XREPXXBT AssetPairInfo
	XREPZEUR AssetPairInfo
	XREPZUSD AssetPairInfo
	XXBTZCAD AssetPairInfo
	XXBTZEUR AssetPairInfo
	XXBTZGBP AssetPairInfo
	XXBTZJPY AssetPairInfo
	XXBTZUSD AssetPairInfo
	XXDGXXBT AssetPairInfo
	XXLMXXBT AssetPairInfo
	XXLMZEUR AssetPairInfo
	XXLMZUSD AssetPairInfo
	XXMRXXBT AssetPairInfo
	XXMRZEUR AssetPairInfo
	XXMRZUSD AssetPairInfo
	XXRPXXBT AssetPairInfo
	XXRPZCAD AssetPairInfo
	XXRPZEUR AssetPairInfo
	XXRPZJPY AssetPairInfo
	XXRPZUSD AssetPairInfo
	XZECXXBT AssetPairInfo
	XZECZEUR AssetPairInfo
	XZECZUSD AssetPairInfo
}

// AssetPairInfo represents asset pair information
type AssetPairInfo struct {
	// Alternate pair name
	Altname string `json:"altname"`
	// Asset class of base component
	AssetClassBase string `json:"aclass_base"`
	// Asset id of base component
	Base string `json:"base"`
	// Asset class of quote component
	AssetClassQuote string `json:"aclass_quote"`
	// Asset id of quote component
	Quote string `json:"quote"`
	// Volume lot size
	Lot string `json:"lot"`
	// Scaling decimal places for pair
	PairDecimals int `json:"pair_decimals"`
	// Scaling decimal places for volume
	LotDecimals int `json:"lot_decimals"`
	// Amount to multiply lot volume by to get currency volume
	LotMultiplier int `json:"lot_multiplier"`
	// Array of leverage amounts available when buying
	LeverageBuy []float32 `json:"leverage_buy"`
	// Array of leverage amounts available when selling
	LeverageSell []float32 `json:"leverage_sell"`
	// Fee schedule array in [volume, percent fee] tuples
	Fees [][]float64 `json:"fees"`
	// // Maker fee schedule array in [volume, percent fee] tuples (if on maker/taker)
	FeesMaker [][]float64 `json:"fees_maker"`
	// // Volume discount currency
	FeeVolumeCurrency string `json:"fee_volume_currency"`
	// Margin call level
	MarginCall int `json:"margin_call"`
	// Stop-out/Liquidation margin level
	MarginStop int `json:"margin_stop"`
}

// AssetsResponse includes asset informations
type AssetsResponse struct {
	DASH AssetInfo
	GNO  AssetInfo
	KFEE AssetInfo
	USDT AssetInfo
	XDAO AssetInfo
	XETC AssetInfo
	XETH AssetInfo
	XICN AssetInfo
	XLTC AssetInfo
	XMLN AssetInfo
	XNMC AssetInfo
	XREP AssetInfo
	XXBT AssetInfo
	XXDG AssetInfo
	XXLM AssetInfo
	XXMR AssetInfo
	XXRP AssetInfo
	XXVN AssetInfo
	XZEC AssetInfo
	ZCAD AssetInfo
	ZEUR AssetInfo
	ZGBP AssetInfo
	ZJPY AssetInfo
	ZKRW AssetInfo
	ZUSD AssetInfo
}

// AssetInfo represents an asset information
type AssetInfo struct {
	// Alternate name
	Altname string
	// Asset class
	AssetClass string `json:"aclass"`
	// Scaling decimal places for record keeping
	Decimals int
	// Scaling decimal places for output display
	DisplayDecimals int `json:"display_decimals"`
}

// BalanceResponse permit to build a Balance object directly from JSON
type BalanceResponse struct {
	DASH float32 `json:"DASH,string"`
	GNO  float32 `json:"GNO,string"`
	KFEE float32 `json:"KFEE,string"`
	USDT float32 `json:"USDT,string"`
	XDAO float32 `json:"XDAO,string"`
	XETC float32 `json:"XETC,string"`
	XETH float32 `json:"XETH,string"`
	XICN float32 `json:"XICN,string"`
	XLTC float32 `json:"XLTC,string"`
	XMLN float32 `json:"XMLN,string"`
	XNMC float32 `json:"XNMC,string"`
	XREP float32 `json:"XREP,string"`
	XXBT float32 `json:"XXBT,string"`
	XXDG float32 `json:"XXDG,string"`
	XXLM float32 `json:"XXLM,string"`
	XXMR float32 `json:"XXMR,string"`
	XXRP float32 `json:"XXRP,string"`
	XXVN float32 `json:"XXVN,string"`
	XZEC float32 `json:"XZEC,string"`
	ZCAD float32 `json:"ZCAD,string"`
	ZEUR float32 `json:"ZEUR,string"`
	ZGBP float32 `json:"ZGBP,string"`
	ZJPY float32 `json:"ZJPY,string"`
	ZKRW float32 `json:"ZKRW,string"`
	ZUSD float32 `json:"ZUSD,string"`
}

// TickerResponse includes the requested ticker pairs
type TickerResponse struct {
	DASHEUR  PairTickerInfo
	DASHUSD  PairTickerInfo
	DASHXBT  PairTickerInfo
	GNOETH   PairTickerInfo
	GNOEUR   PairTickerInfo
	GNOUSD   PairTickerInfo
	GNOXBT   PairTickerInfo
	USDTZUSD PairTickerInfo
	XETCXETH PairTickerInfo
	XETCXXBT PairTickerInfo
	XETCZEUR PairTickerInfo
	XETCXUSD PairTickerInfo
	XETHXXBT PairTickerInfo
	XETHZCAD PairTickerInfo
	XETHZEUR PairTickerInfo
	XETHZGBP PairTickerInfo
	XETHZJPY PairTickerInfo
	XETHZUSD PairTickerInfo
	XICNXETH PairTickerInfo
	XICNXXBT PairTickerInfo
	XLTCXXBT PairTickerInfo
	XLTCZEUR PairTickerInfo
	XLTCZUSD PairTickerInfo
	XMLNXETH PairTickerInfo
	XMLNXXBT PairTickerInfo
	XREPXETH PairTickerInfo
	XREPXXBT PairTickerInfo
	XREPZEUR PairTickerInfo
	XREPZUSD PairTickerInfo
	XXBTZCAD PairTickerInfo
	XXBTZEUR PairTickerInfo
	XXBTZGBP PairTickerInfo
	XXBTZJPY PairTickerInfo
	XXBTZUSD PairTickerInfo
	XXDGXXBT PairTickerInfo
	XXLMXXBT PairTickerInfo
	XXLMZEUR PairTickerInfo
	XXLMZUSD PairTickerInfo
	XXMRXXBT PairTickerInfo
	XXMRZEUR PairTickerInfo
	XXMRZUSD PairTickerInfo
	XXRPXXBT PairTickerInfo
	XXRPZCAD PairTickerInfo
	XXRPZEUR PairTickerInfo
	XXRPZJPY PairTickerInfo
	XXRPZUSD PairTickerInfo
	XZECXXBT PairTickerInfo
	XZECZEUR PairTickerInfo
	XZECZUSD PairTickerInfo
}

// PairTickerInfo represents ticker information for a pair
type PairTickerInfo struct {
	// Ask array(<price>, <whole lot volume>, <lot volume>)
	Ask []string `json:"a"`
	// Bid array(<price>, <whole lot volume>, <lot volume>)
	Bid []string `json:"b"`
	// Last trade closed array(<price>, <lot volume>)
	Close []string `json:"c"`
	// Volume array(<today>, <last 24 hours>)
	Volume []string `json:"v"`
	// Volume weighted average price array(<today>, <last 24 hours>)
	VolumeAveragePrice []string `json:"p"`
	// Number of trades array(<today>, <last 24 hours>)
	Trades []int `json:"t"`
	// Low array(<today>, <last 24 hours>)
	Low []string `json:"l"`
	// High array(<today>, <last 24 hours>)
	High []string `json:"h"`
	// Today's opening price
	OpeningPrice float32 `json:"o,string"`
}

// TradesResponse represents a list of the last trades
type TradesResponse struct {
	Last   int64
	Trades []TradeInfo
}

// TradeInfo represents a trades information
type TradeInfo struct {
	Price         string
	PriceFloat    float64
	PriceInt      int64
	Volume        float64
	Time          float64
	Buy           bool
	Sell          bool
	Market        bool
	Limit         bool
	Miscellaneous string
}

// OrderDescription represents an orders description
type OrderDescription struct {
	AssetPair      string  `json:"pair"`
	Close          string  `json:"close"`
	Leverage       string  `json:"leverage"`
	Order          string  `json:"order"`
	OrderType      string  `json:"ordertype"`
	PrimaryPrice   float64 `json:"price,string"`
	SecondaryPrice float64 `json:"price2,string"`
	Type           string  `json:"type"`
}

// Order represents a single order
type Order struct {
	ReferenceID    string           `json:"refid"`
	UserRef        string           `json:"userref"`
	Status         string           `json:"status"`
	OpenTime       float64          `json:"opentm"`
	StartTime      float64          `json:"starttm"`
	ExpireTime     float64          `json:"expiretm"`
	Description    OrderDescription `json:"descr"`
	Volume         string           `json:"vol1"`
	VolumeExecuted float64          `json:"vol_exec,string"`
	Cost           float64          `json:"cost,string"`
	Fee            string           `json:"fee"`
	Price          float64          `json:"price,string"`
	StopPrice      float64          `json:"stopprice"`
	LimitPrice     float64          `json:"limitprice"`
	Misc           string           `json:"misc"`
	OrderFlags     string           `json:"oflags"`
	CloseTime      float64          `json:"closetm"`
	Reason         string           `json:"reason"`
}

// ClosedOrdersResponse represents a list of closed orders, indexed by id
type ClosedOrdersResponse struct {
	Closed map[string]Order `json:"closed"`
	Count  int              `json:"count"`
}

func NewOHLC(input []interface{}) (*OHLC, error) {
	if len(input) != 8 {
		return nil, fmt.Errorf("the length is not 8 but %d", len(input))
	}

	tmp := new(OHLC)

	tmp.Time = time.Unix(input[0].(int64), 0)
	tmp.Open = input[1].(float64)
	tmp.High = input[2].(float64)
	tmp.Low = input[3].(float64)
	tmp.Close = input[4].(float64)
	tmp.Vwap = input[5].(float64)
	tmp.Volume = input[6].(float64)
	tmp.Count = input[7].(int)

	return tmp, nil
}

// OHLC represents the "Open-high-low-close chart"
type OHLC struct {
	Time   time.Time `json:"time"`
	Open   float64   `json:"open"`
	High   float64   `json:"high"`
	Low    float64   `json:"low"`
	Close  float64   `json:"close"`
	Vwap   float64   `json:"vwap"`
	Volume float64   `json:"volume"`
	Count  int       `json:"count"`
}

func newOHLCResponse(pair string) OHLCResponse {
	switch pair {
	case DASHEUR:
		return new(OHLCResponseDASHEUR)
	case DASHUSD:
		return new(OHLCResponseDASHUSD)
	case DASHXBT:
		return new(OHLCResponseDASHXBT)
	case GNOETH:
		return new(OHLCResponseGNOETH)
	case GNOEUR:
		return new(OHLCResponseGNOEUR)
	case GNOUSD:
		return new(OHLCResponseGNOUSD)
	case GNOXBT:
		return new(OHLCResponseGNOXBT)
	case USDTZUSD:
		return new(OHLCResponseUSDTZUSD)
	case XETCXETH:
		return new(OHLCResponseXETCXETH)
	case XETCXXBT:
		return new(OHLCResponseXETCXXBT)
	case XETCZEUR:
		return new(OHLCResponseXETCZEUR)
	case XETCXUSD:
		return new(OHLCResponseXETCXUSD)
	case XETHXXBT:
		return new(OHLCResponseXETHXXBT)
	case XETHZCAD:
		return new(OHLCResponseXETHZCAD)
	case XETHZEUR:
		return new(OHLCResponseXETHZEUR)
	case XETHZGBP:
		return new(OHLCResponseXETHZGBP)
	case XETHZJPY:
		return new(OHLCResponseXETHZJPY)
	case XETHZUSD:
		return new(OHLCResponseXETHZUSD)
	case XICNXETH:
		return new(OHLCResponseXICNXETH)
	case XICNXXBT:
		return new(OHLCResponseXICNXXBT)
	case XLTCXXBT:
		return new(OHLCResponseXLTCXXBT)
	case XLTCZEUR:
		return new(OHLCResponseXLTCZEUR)
	case XLTCZUSD:
		return new(OHLCResponseXLTCZUSD)
	case XMLNXETH:
		return new(OHLCResponseXMLNXETH)
	case XMLNXXBT:
		return new(OHLCResponseXMLNXXBT)
	case XREPXETH:
		return new(OHLCResponseXREPXETH)
	case XREPXXBT:
		return new(OHLCResponseXREPXXBT)
	case XREPZEUR:
		return new(OHLCResponseXREPZEUR)
	case XREPZUSD:
		return new(OHLCResponseXREPZUSD)
	case XXBTZCAD:
		return new(OHLCResponseXXBTZCAD)
	case XXBTZEUR:
		return new(OHLCResponseXXBTZEUR)
	case XXBTZGBP:
		return new(OHLCResponseXXBTZGBP)
	case XXBTZJPY:
		return new(OHLCResponseXXBTZJPY)
	case XXBTZUSD:
		return new(OHLCResponseXXBTZUSD)
	case XXDGXXBT:
		return new(OHLCResponseXXDGXXBT)
	case XXLMXXBT:
		return new(OHLCResponseXXLMXXBT)
	case XXLMZEUR:
		return new(OHLCResponseXXLMZEUR)
	case XXLMZUSD:
		return new(OHLCResponseXXLMZUSD)
	case XXMRXXBT:
		return new(OHLCResponseXXMRXXBT)
	case XXMRZEUR:
		return new(OHLCResponseXXMRZEUR)
	case XXMRZUSD:
		return new(OHLCResponseXXMRZUSD)
	case XXRPXXBT:
		return new(OHLCResponseXXRPXXBT)
	case XXRPZCAD:
		return new(OHLCResponseXXRPZCAD)
	case XXRPZEUR:
		return new(OHLCResponseXXRPZEUR)
	case XXRPZJPY:
		return new(OHLCResponseXXRPZJPY)
	case XXRPZUSD:
		return new(OHLCResponseXXRPZUSD)
	case XZECXXBT:
		return new(OHLCResponseXZECXXBT)
	case XZECZEUR:
		return new(OHLCResponseXZECZEUR)
	case XZECZUSD:
		return new(OHLCResponseXZECZUSD)
	}
	return nil
}

type OHLCResponse interface {
	GetPair() string
	GetOHLC() []*OHLC
	GetLast() float64
}

type OHLCResponseBase struct {
	Pair      string
	OHLCArray [][]interface{}
	OHLC      []*OHLC
	Last      float64
}

func (o *OHLCResponseBase) GetPair() string {
	return o.Pair
}
func (o *OHLCResponseBase) GetOHLC() []*OHLC {
	if o.OHLC == nil {
		err := o.setOHLC()
		if err != nil {
			return nil
		}
	}
	return o.OHLC
}
func (o *OHLCResponseBase) GetLast() float64 {
	return o.Last
}
func (o *OHLCResponseBase) setOHLC() error {
	if o.OHLCArray == nil {
		return fmt.Errorf("the OHLCResponseBase.OHLCArray is nil and can't populate OHLCResponseBase.OHLC field")
	}

	for _, tmpOHLCArray := range o.OHLCArray {
		tmpOHLC, buildErr := NewOHLC(tmpOHLCArray)
		if buildErr != nil {
			continue
		}

		o.OHLC = append(o.OHLC, tmpOHLC)
	}

	return nil
}

type OHLCResponseDASHEUR struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"DASHEUR"`
}
type OHLCResponseDASHUSD struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"DASHUSD"`
}
type OHLCResponseDASHXBT struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"DASHXBT"`
}
type OHLCResponseGNOETH struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"GNOETH"`
}
type OHLCResponseGNOEUR struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"GNOEUR"`
}
type OHLCResponseGNOUSD struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"GNOUSD"`
}
type OHLCResponseGNOXBT struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"GNOXBT"`
}
type OHLCResponseUSDTZUSD struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"USDTZUSD"`
}
type OHLCResponseXETCXETH struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XETCXETH"`
}
type OHLCResponseXETCXXBT struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XETCXXBT"`
}
type OHLCResponseXETCZEUR struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XETCZEUR"`
}
type OHLCResponseXETCXUSD struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XETCXUSD"`
}
type OHLCResponseXETHXXBT struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XETHXXBT"`
}
type OHLCResponseXETHZCAD struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XETHZCAD"`
}
type OHLCResponseXETHZEUR struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XETHZEUR"`
}
type OHLCResponseXETHZGBP struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XETHZGBP"`
}
type OHLCResponseXETHZJPY struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XETHZJPY"`
}
type OHLCResponseXETHZUSD struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XETHZUSD"`
}
type OHLCResponseXICNXETH struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XICNXETH"`
}
type OHLCResponseXICNXXBT struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XICNXXBT"`
}
type OHLCResponseXLTCXXBT struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XLTCXXBT"`
}
type OHLCResponseXLTCZEUR struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XLTCZEUR"`
}
type OHLCResponseXLTCZUSD struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XLTCZUSD"`
}
type OHLCResponseXMLNXETH struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XMLNXETH"`
}
type OHLCResponseXMLNXXBT struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XMLNXXBT"`
}
type OHLCResponseXREPXETH struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XREPXETH"`
}
type OHLCResponseXREPXXBT struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XREPXXBT"`
}
type OHLCResponseXREPZEUR struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XREPZEUR"`
}
type OHLCResponseXREPZUSD struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XREPZUSD"`
}
type OHLCResponseXXBTZCAD struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XXBTZCAD"`
}
type OHLCResponseXXBTZEUR struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XXBTZEUR"`
}
type OHLCResponseXXBTZGBP struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XXBTZGBP"`
}
type OHLCResponseXXBTZJPY struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XXBTZJPY"`
}
type OHLCResponseXXBTZUSD struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XXBTZUSD"`
}
type OHLCResponseXXDGXXBT struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XXDGXXBT"`
}
type OHLCResponseXXLMXXBT struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XXLMXXBT"`
}
type OHLCResponseXXLMZEUR struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XXLMZEUR"`
}
type OHLCResponseXXLMZUSD struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XXLMZUSD"`
}
type OHLCResponseXXMRXXBT struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XXMRXXBT"`
}
type OHLCResponseXXMRZEUR struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XXMRZEUR"`
}
type OHLCResponseXXMRZUSD struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XXMRZUSD"`
}
type OHLCResponseXXRPXXBT struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XXRPXXBT"`
}
type OHLCResponseXXRPZCAD struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XXRPZCAD"`
}
type OHLCResponseXXRPZEUR struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XXRPZEUR"`
}
type OHLCResponseXXRPZJPY struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XXRPZJPY"`
}
type OHLCResponseXXRPZUSD struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XXRPZUSD"`
}
type OHLCResponseXZECXXBT struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XZECXXBT"`
}
type OHLCResponseXZECZEUR struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XZECZEUR"`
}
type OHLCResponseXZECZUSD struct {
	*OHLCResponseBase
	OHLCArray []interface{} `json:"XZECZUSD"`
}
