package krakenapi

import (
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"time"
)

// trade pairs constants
const (
	ADACAD   = "ADACAD"
	ADAETH   = "ADAETH"
	ADAEUR   = "ADAEUR"
	ADAUSD   = "ADAUSD"
	ADAXBT   = "ADAXBT"
	AAVEUSD  = "AAVEUSD"
	BCHEUR   = "BCHEUR"
	BCHUSD   = "BCHUSD"
	BCHXBT   = "BCHXBT"
	DASHEUR  = "DASHEUR"
	DASHUSD  = "DASHUSD"
	DASHXBT  = "DASHXBT"
	EOSETH   = "EOSETH"
	EOSEUR   = "EOSEUR"
	EOSUSD   = "EOSUSD"
	EOSXBT   = "EOSXBT"
	GNOETH   = "GNOETH"
	GNOEUR   = "GNOEUR"
	GNOUSD   = "GNOUSD"
	GNOXBT   = "GNOXBT"
	QTUMCAD  = "QTUMCAD"
	QTUMETH  = "QTUMETH"
	QTUMEUR  = "QTUMEUR"
	QTUMUSD  = "QTUMUSD"
	QTUMXBT  = "QTUMXBT"
	REPV2USD = "REPV2USD"
	TRXUSD   = "TRXUSD"
	UNIUSD   = "UNIUSD"
	USDTZUSD = "USDTZUSD"
	XBTUSDT  = "XBTUSDT"
	XETCXETH = "XETCXETH"
	XETCXXBT = "XETCXXBT"
	XETCZEUR = "XETCZEUR"
	XETCZUSD = "XETCZUSD"
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
	XTZCAD   = "XTZCAD"
	XTZETH   = "XTZETH"
	XTZEUR   = "XTZEUR"
	XTZUSD   = "XTZUSD"
	XTZXBT   = "XTZXBT"
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

// actions constants
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
	ADACAD   AssetPairInfo
	AAVEUSD  AssetPairInfo
	ADAETH   AssetPairInfo
	ADAEUR   AssetPairInfo
	ADAUSD   AssetPairInfo
	ADAXBT   AssetPairInfo
	BCHEUR   AssetPairInfo
	BCHUSD   AssetPairInfo
	BCHXBT   AssetPairInfo
	DASHEUR  AssetPairInfo
	DASHUSD  AssetPairInfo
	DASHXBT  AssetPairInfo
	EOSETH   AssetPairInfo
	EOSEUR   AssetPairInfo
	EOSUSD   AssetPairInfo
	EOSXBT   AssetPairInfo
	GNOETH   AssetPairInfo
	GNOEUR   AssetPairInfo
	GNOUSD   AssetPairInfo
	GNOXBT   AssetPairInfo
	QTUMCAD  AssetPairInfo
	QTUMETH  AssetPairInfo
	QTUMEUR  AssetPairInfo
	QTUMUSD  AssetPairInfo
	QTUMXBT  AssetPairInfo
	REPV2USD AssetPairInfo
	TRXUSD   AssetPairInfo
	UNIUSD   AssetPairInfo
	USDTZUSD AssetPairInfo
	XETCXETH AssetPairInfo
	XETCXXBT AssetPairInfo
	XETCZEUR AssetPairInfo
	XETCZUSD AssetPairInfo
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
	XTZCAD   AssetPairInfo
	XTZETH   AssetPairInfo
	XTZEUR   AssetPairInfo
	XTZUSD   AssetPairInfo
	XTZXBT   AssetPairInfo
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
	LeverageBuy []float64 `json:"leverage_buy"`
	// Array of leverage amounts available when selling
	LeverageSell []float64 `json:"leverage_sell"`
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
	ADA   AssetInfo
	AAVE  AssetInfo
	BCH   AssetInfo
	DASH  AssetInfo
	EOS   AssetInfo
	GNO   AssetInfo
	KFEE  AssetInfo
	QTUM  AssetInfo
	REPV2 AssetInfo
	TRX   AssetInfo
	UNI   AssetInfo
	USDT  AssetInfo
	XDAO  AssetInfo
	XETC  AssetInfo
	XETH  AssetInfo
	XICN  AssetInfo
	XLTC  AssetInfo
	XMLN  AssetInfo
	XNMC  AssetInfo
	XREP  AssetInfo
	XXBT  AssetInfo
	XXDG  AssetInfo
	XXLM  AssetInfo
	XXMR  AssetInfo
	XXRP  AssetInfo
	XTZ   AssetInfo
	XXVN  AssetInfo
	XZEC  AssetInfo
	ZCAD  AssetInfo
	ZEUR  AssetInfo
	ZGBP  AssetInfo
	ZJPY  AssetInfo
	ZKRW  AssetInfo
	ZUSD  AssetInfo
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

// BalanceResponse represents the account's balances (list of currencies)
type BalanceResponse struct {
	ADA   float64 `json:"ADA,string"`
	AAVE  float64 `json:"AAVE,string"`
	BCH   float64 `json:"BCH,string"`
	DASH  float64 `json:"DASH,string"`
	EOS   float64 `json:"EOS,string"`
	GNO   float64 `json:"GNO,string"`
	QTUM  float64 `json:"QTUM,string"`
	KFEE  float64 `json:"KFEE,string"`
	REPV2 float64 `json:"REPV2,string"`
	TRX   float64 `json:"TRX,string"`
	UNI   float64 `json:"UNI,string"`
	USDT  float64 `json:"USDT,string"`
	XDAO  float64 `json:"XDAO,string"`
	XETC  float64 `json:"XETC,string"`
	XETH  float64 `json:"XETH,string"`
	XICN  float64 `json:"XICN,string"`
	XLTC  float64 `json:"XLTC,string"`
	XMLN  float64 `json:"XMLN,string"`
	XNMC  float64 `json:"XNMC,string"`
	XREP  float64 `json:"XREP,string"`
	XXBT  float64 `json:"XXBT,string"`
	XXDG  float64 `json:"XXDG,string"`
	XXLM  float64 `json:"XXLM,string"`
	XXMR  float64 `json:"XXMR,string"`
	XXRP  float64 `json:"XXRP,string"`
	XTZ   float64 `json:"XTZ,string"`
	XXVN  float64 `json:"XXVN,string"`
	XZEC  float64 `json:"XZEC,string"`
	ZCAD  float64 `json:"ZCAD,string"`
	ZEUR  float64 `json:"ZEUR,string"`
	ZGBP  float64 `json:"ZGBP,string"`
	ZJPY  float64 `json:"ZJPY,string"`
	ZKRW  float64 `json:"ZKRW,string"`
	ZUSD  float64 `json:"ZUSD,string"`
}

// TradeBalanceResponse struct used as the response for the TradeBalance method
type TradeBalanceResponse struct {
	EquivalentBalance         float64 `json:"eb,string"`
	TradeBalance              float64 `json:"tb,string"`
	MarginOP                  float64 `json:"m,string"`
	UnrealizedNetProfitLossOP float64 `json:"n,string"`
	CostBasisOP               float64 `json:"c,string"`
	CurrentValuationOP        float64 `json:"v,string"`
	Equity                    float64 `json:"e,string"`
	FreeMargin                float64 `json:"mf,string"`
	MarginLevel               float64 `json:"ml,string"`
}

// Fees includes fees information for different currencies
type Fees struct {
	ADACAD   FeeInfo
	ADAETH   FeeInfo
	ADAEUR   FeeInfo
	ADAUSD   FeeInfo
	ADAXBT   FeeInfo
	AAVEUSD  FeeInfo
	BCHEUR   FeeInfo
	BCHUSD   FeeInfo
	BCHXBT   FeeInfo
	DASHEUR  FeeInfo
	DASHUSD  FeeInfo
	DASHXBT  FeeInfo
	EOSETH   FeeInfo
	EOSEUR   FeeInfo
	EOSUSD   FeeInfo
	EOSXBT   FeeInfo
	GNOETH   FeeInfo
	GNOEUR   FeeInfo
	GNOUSD   FeeInfo
	GNOXBT   FeeInfo
	QTUMCAD  FeeInfo
	QTUMETH  FeeInfo
	QTUMEUR  FeeInfo
	QTUMUSD  FeeInfo
	QTUMXBT  FeeInfo
	REPV2    FeeInfo
	TRX      FeeInfo
	UNI      FeeInfo
	USDTZUSD FeeInfo
	XETCXETH FeeInfo
	XETCXXBT FeeInfo
	XETCZEUR FeeInfo
	XETCZUSD FeeInfo
	XETHXXBT FeeInfo
	XETHZCAD FeeInfo
	XETHZEUR FeeInfo
	XETHZGBP FeeInfo
	XETHZJPY FeeInfo
	XETHZUSD FeeInfo
	XICNXETH FeeInfo
	XICNXXBT FeeInfo
	XLTCXXBT FeeInfo
	XLTCZEUR FeeInfo
	XLTCZUSD FeeInfo
	XMLNXETH FeeInfo
	XMLNXXBT FeeInfo
	XREPXETH FeeInfo
	XREPXXBT FeeInfo
	XREPZEUR FeeInfo
	XREPZUSD FeeInfo
	XTZCAD   FeeInfo
	XTZETH   FeeInfo
	XTZEUR   FeeInfo
	XTZUSD   FeeInfo
	XTZXBT   FeeInfo
	XXBTZCAD FeeInfo
	XXBTZEUR FeeInfo
	XXBTZGBP FeeInfo
	XXBTZJPY FeeInfo
	XXBTZUSD FeeInfo
	XXDGXXBT FeeInfo
	XXLMXXBT FeeInfo
	XXLMZEUR FeeInfo
	XXLMZUSD FeeInfo
	XXMRXXBT FeeInfo
	XXMRZEUR FeeInfo
	XXMRZUSD FeeInfo
	XXRPXXBT FeeInfo
	XXRPZCAD FeeInfo
	XXRPZEUR FeeInfo
	XXRPZJPY FeeInfo
	XXRPZUSD FeeInfo
	XZECXXBT FeeInfo
	XZECZEUR FeeInfo
	XZECZUSD FeeInfo
}

// FeeInfo represents a fee information
type FeeInfo struct {
	Fee        float64 `json:"fee,string"`
	MinFee     float64 `json:"minfee,string"`
	MaxFee     float64 `json:"maxfee,string"`
	NextFee    float64 `json:"nextfee,string"`
	NextVolume float64 `json:"nextvolume,string"`
	TierVolume float64 `json:"tiervolume,string"`
}

// TradeVolumeResponse represents the response for trade volume
type TradeVolumeResponse struct {
	Volume    float64 `json:"volume,string"`
	Currency  string  `json:"currency"`
	Fees      Fees    `json:"fees"`
	FeesMaker Fees    `json:"fees_maker"`
}

// TickerResponse includes the requested ticker pairs
type TickerResponse struct {
	ADACAD   PairTickerInfo
	ADAETH   PairTickerInfo
	ADAEUR   PairTickerInfo
	ADAUSD   PairTickerInfo
	ADAXBT   PairTickerInfo
	AAVEUSD  PairTickerInfo
	BCHEUR   PairTickerInfo
	BCHUSD   PairTickerInfo
	BCHXBT   PairTickerInfo
	DASHEUR  PairTickerInfo
	DASHUSD  PairTickerInfo
	DASHXBT  PairTickerInfo
	EOSETH   PairTickerInfo
	EOSEUR   PairTickerInfo
	EOSUSD   PairTickerInfo
	EOSXBT   PairTickerInfo
	GNOETH   PairTickerInfo
	GNOEUR   PairTickerInfo
	GNOUSD   PairTickerInfo
	GNOXBT   PairTickerInfo
	QTUMCAD  PairTickerInfo
	QTUMETH  PairTickerInfo
	QTUMEUR  PairTickerInfo
	QTUMUSD  PairTickerInfo
	QTUMXBT  PairTickerInfo
	REPV2USD PairTickerInfo
	TRXUSD   PairTickerInfo
	UNIUSD   PairTickerInfo
	USDTZUSD PairTickerInfo
	XBTUSDT  PairTickerInfo
	XETCXETH PairTickerInfo
	XETCXXBT PairTickerInfo
	XETCZEUR PairTickerInfo
	XETCZUSD PairTickerInfo
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
	XTZCAD   PairTickerInfo
	XTZETH   PairTickerInfo
	XTZEUR   PairTickerInfo
	XTZUSD   PairTickerInfo
	XTZXBT   PairTickerInfo
	XZECXXBT PairTickerInfo
	XZECZEUR PairTickerInfo
	XZECZUSD PairTickerInfo
}

// DepositAddressesResponse is the response type of a DepositAddresses query to the Kraken API.
type DepositAddressesResponse []struct {
	Address  string `json:"address"`
	Expiretm string `json:"expiretm"`
	New      bool   `json:"new,omitempty"`
}

// WithdrawResponse is the response type of a Withdraw query to the Kraken API.
type WithdrawResponse struct {
	RefID string `json:"refid"`
}

// WithdrawInfoResponse is the response type showing withdrawal information for a selected withdrawal method.
type WithdrawInfoResponse struct {
	Method string    `json:"method"`
	Limit  big.Float `json:"limit"`
	Amount big.Float `json:"amount"`
	Fee    big.Float `json:"fee"`
}

// GetPairTickerInfo is a helper method that returns given `pair`'s `PairTickerInfo`
func (v *TickerResponse) GetPairTickerInfo(pair string) PairTickerInfo {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(pair)

	return f.Interface().(PairTickerInfo)
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
	OpeningPrice float64 `json:"o,string"`
}

// TradesResponse represents a list of the last trades
type TradesResponse struct {
	Last   int64
	Trades []TradeInfo
}

// TradesHistoryResponse represents a list of executed trade
type TradesHistoryResponse struct {
	Trades map[string]TradeHistoryInfo `json:"trades"`
	Count  int                         `json:"count"`
}

// TradeHistoryInfo represents a transaction
type TradeHistoryInfo struct {
	TransactionID string  `json:"ordertxid"`
	PostxID       string  `json:"postxid"`
	AssetPair     string  `json:"pair"`
	Time          float64 `json:"time"`
	Type          string  `json:"type"`
	OrderType     string  `json:"ordertype"`
	Price         float64 `json:"price,string"`
	Cost          float64 `json:"cost,string"`
	Fee           float64 `json:"fee,string"`
	Volume        float64 `json:"vol,string"`
	Margin        float64 `json:"margin,string"`
	Misc          string  `json:"misc"`
}

// TradeInfo represents a trades information
type TradeInfo struct {
	Price         string
	PriceFloat    float64
	Volume        string
	VolumeFloat   float64
	Time          int64
	Buy           bool
	Sell          bool
	Market        bool
	Limit         bool
	Miscellaneous string
}

// LedgersResponse represents an associative array of ledgers infos
type LedgersResponse struct {
	Ledger map[string]LedgerInfo `json:"ledger"`
}

// LedgerInfo Represents the ledger informations
type LedgerInfo struct {
	RefID   string    `json:"refid"`
	Time    float64   `json:"time"`
	Type    string    `json:"type"`
	Aclass  string    `json:"aclass"`
	Asset   string    `json:"asset"`
	Amount  big.Float `json:"amount"`
	Fee     big.Float `json:"fee"`
	Balance big.Float `json:"balance"`
}

// OrderTypes for AddOrder
const (
	OTMarket              = "market"
	OTLimit               = "limit"                  // (price = limit price)
	OTStopLoss            = "stop-loss"              // (price = stop loss price)
	OTTakeProfi           = "take-profit"            // (price = take profit price)
	OTStopLossProfit      = "stop-loss-profit"       // (price = stop loss price, price2 = take profit price)
	OTStopLossProfitLimit = "stop-loss-profit-limit" // (price = stop loss price, price2 = take profit price)
	OTStopLossLimit       = "stop-loss-limit"        // (price = stop loss trigger price, price2 = triggered limit price)
	OTTakeProfitLimit     = "take-profit-limit"      // (price = take profit trigger price, price2 = triggered limit price)
	OTTrailingStop        = "trailing-stop"          // (price = trailing stop offset)
	OTTrailingStopLimit   = "trailing-stop-limit"    // (price = trailing stop offset, price2 = triggered limit offset)
	OTStopLossAndLimit    = "stop-loss-and-limit"    // (price = stop loss price, price2 = limit price)
	OTSettlePosition      = "settle-position"
)

// OrderDescription represents an orders description
type OrderDescription struct {
	AssetPair      string `json:"pair"`
	Close          string `json:"close"`
	Leverage       string `json:"leverage"`
	Order          string `json:"order"`
	OrderType      string `json:"ordertype"`
	PrimaryPrice   string `json:"price"`
	SecondaryPrice string `json:"price2"`
	Type           string `json:"type"`
}

// Order represents a single order
type Order struct {
	TransactionID  string           `json:"-"`
	ReferenceID    string           `json:"refid"`
	UserRef        int              `json:"userref"`
	Status         string           `json:"status"`
	OpenTime       float64          `json:"opentm"`
	StartTime      float64          `json:"starttm"`
	ExpireTime     float64          `json:"expiretm"`
	Description    OrderDescription `json:"descr"`
	Volume         string           `json:"vol"`
	VolumeExecuted float64          `json:"vol_exec,string"`
	Cost           float64          `json:"cost,string"`
	Fee            float64          `json:"fee,string"`
	Price          float64          `json:"price,string"`
	StopPrice      float64          `json:"stopprice.string"`
	LimitPrice     float64          `json:"limitprice,string"`
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

// OrderBookItem is a piece of information about an order.
type OrderBookItem struct {
	Price  float64
	Amount float64
	Ts     int64
}

// UnmarshalJSON takes a json array from kraken and converts it into an OrderBookItem.
func (o *OrderBookItem) UnmarshalJSON(data []byte) error {
	tmpStruct := struct {
		price  string
		amount string
		ts     int64
	}{}
	tmpArray := []interface{}{&tmpStruct.price, &tmpStruct.amount, &tmpStruct.ts}
	err := json.Unmarshal(data, &tmpArray)
	if err != nil {
		return err
	}

	o.Price, err = strconv.ParseFloat(tmpStruct.price, 64)
	if err != nil {
		return err
	}
	o.Amount, err = strconv.ParseFloat(tmpStruct.amount, 64)
	if err != nil {
		return err
	}
	o.Ts = tmpStruct.ts
	return nil
}

// DepthResponse is a response from kraken to Depth request.
type DepthResponse map[string]OrderBook

// OrderBook contains top asks and bids.
type OrderBook struct {
	Asks []OrderBookItem
	Bids []OrderBookItem
}

// OpenOrdersResponse response when opening an order
type OpenOrdersResponse struct {
	Open  map[string]Order `json:"open"`
	Count int              `json:"count"`
}

// AddOrderResponse response when adding an order
type AddOrderResponse struct {
	Description    OrderDescription `json:"descr"`
	TransactionIds []string         `json:"txid"`
}

// CancelOrderResponse response when cancelling and order
type CancelOrderResponse struct {
	Count   int  `json:"count"`
	Pending bool `json:"pending"`
}

// QueryOrdersResponse response when checking all orders
type QueryOrdersResponse map[string]Order

// NewOHLC constructor for OHLC
func NewOHLC(input []interface{}) (*OHLC, error) {
	if len(input) != 8 {
		return nil, fmt.Errorf("the length is not 8 but %d", len(input))
	}

	tmp := new(OHLC)
	tmp.Time = time.Unix(int64(input[0].(float64)), 0)
	tmp.Open, _ = strconv.ParseFloat(input[1].(string), 64)
	tmp.High, _ = strconv.ParseFloat(input[2].(string), 64)
	tmp.Low, _ = strconv.ParseFloat(input[3].(string), 64)
	tmp.Close, _ = strconv.ParseFloat(input[4].(string), 64)
	tmp.Vwap, _ = strconv.ParseFloat(input[5].(string), 64)
	tmp.Volume, _ = strconv.ParseFloat(input[6].(string), 64)
	tmp.Count = int(input[7].(float64))

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

// OHLCResponse represents the OHLC's response
type OHLCResponse struct {
	Pair string  `json:"pair"`
	OHLC []*OHLC `json:"OHLC"`
	Last float64 `json:"last"`
}
