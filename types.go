package krakenapi

const (
	DASHEUR = "DASHEUR"
	DASHUSD = "DASHUSD"
	DASHXBT = "DASHXBT"
	GNOETH = "GNOETH"
	GNOEUR = "GNOEUR"
	GNOUSD = "GNOUSD"
	GNOXBT = "GNOXBT"
	XETHXXBT = "XETHXXBT"
	XETHZCAD = "XETHZCAD"
	XETHZEUR = "XETHZEUR"
	XETHZGBP = "XETHZGBP"
	XETHZJPY = "XETHZJPY"
	XETHZUSD = "XETHZUSD"
	XLTCZCAD = "XLTCZCAD"
	XLTCZEUR = "XLTCZEUR"
	XLTCZUSD = "XLTCZUSD"
	XXBTXLTC = "XXBTXLTC"
	XXBTXNMC = "XXBTXNMC"
	XXBTXXDG = "XXBTXXDG"
	XXBTXXLM = "XXBTXXLM"
	XXBTXXRP = "XXBTXXRP"
	XXBTZCAD = "XXBTZCAD"
	XXBTZEUR = "XXBTZEUR"
	XXBTZGBP = "XXBTZGBP"
	XXBTZJPY = "XXBTZJPY"
	XXBTZUSD = "XXBTZUSD"
	XXLMZUSD = "XXLMZUSD"
	XXLMZEUR = "XXLMZEUR"
	XXLMXXBT = "XXLMXXBT"
	XXMRZUSD = "XXMRZUSD"
	XXMRZEUR = "XXMRZEUR"
	XXMRXXBT = "XXMRXXBT"
	XZECZUSD = "XZECZUSD"
	XZECZEUR = "XZECZEUR"
	XZECXXBT = "XZECXXBT"
)

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
	DASHEUR AssetPairInfo
	DASHUSD AssetPairInfo
	DASHXBT AssetPairInfo
	GNOETH AssetPairInfo
	GNOEUR AssetPairInfo
	GNOUSD AssetPairInfo
	GNOXBT AssetPairInfo
	XETHXXBT AssetPairInfo
	XETHZCAD AssetPairInfo
	XETHZEUR AssetPairInfo
	XETHZGBP AssetPairInfo
	XETHZJPY AssetPairInfo
	XETHZUSD AssetPairInfo
	XLTCZCAD AssetPairInfo
	XLTCZEUR AssetPairInfo
	XLTCZUSD AssetPairInfo
	XXBTXLTC AssetPairInfo
	XXBTXNMC AssetPairInfo
	XXBTXXDG AssetPairInfo
	XXBTXXLM AssetPairInfo
	XXBTXXRP AssetPairInfo
	XXBTZCAD AssetPairInfo
	XXBTZEUR AssetPairInfo
	XXBTZGBP AssetPairInfo
	XXBTZJPY AssetPairInfo
	XXBTZUSD AssetPairInfo
	XXLMZUSD AssetPairInfo
	XXLMZEUR AssetPairInfo
	XXLMZXBT AssetPairInfo
	XXMRZUSD AssetPairInfo
	XXMRZEUR AssetPairInfo
	XXMRZXBT AssetPairInfo
	XZECZUSD AssetPairInfo
	XZECZEUR AssetPairInfo
	XZECZXBT AssetPairInfo
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
	GNO AssetInfo
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
	AssetClass string
	// Scaling decimal places for record keeping
	Decimals int
	// Scaling decimal places for output display
	DisplayDecimals int `json:"display_decimals"`
}

// TickerResponse includes the requested ticker pairs
type TickerResponse struct {
	DASHEUR PairTickerInfo
	DASHUSD PairTickerInfo
	DASHXBT PairTickerInfo
	GNOETH PairTickerInfo
	GNOEUR PairTickerInfo
	GNOUSD PairTickerInfo
	GNOXBT PairTickerInfo
	XETHXXBT PairTickerInfo
	XETHZCAD PairTickerInfo
	XETHZEUR PairTickerInfo
	XETHZGBP PairTickerInfo
	XETHZJPY PairTickerInfo
	XETHZUSD PairTickerInfo
	XLTCZCAD PairTickerInfo
	XLTCZEUR PairTickerInfo
	XLTCZUSD PairTickerInfo
	XXBTXLTC PairTickerInfo
	XXBTXNMC PairTickerInfo
	XXBTXXDG PairTickerInfo
	XXBTXXLM PairTickerInfo
	XXBTXXRP PairTickerInfo
	XXBTZCAD PairTickerInfo
	XXBTZEUR PairTickerInfo
	XXBTZGBP PairTickerInfo
	XXBTZJPY PairTickerInfo
	XXBTZUSD PairTickerInfo
	XXLMZUSD PairTickerInfo
	XXLMZEUR PairTickerInfo
	XXLMXXBT PairTickerInfo
	XXMRZUSD PairTickerInfo
	XXMRZEUR PairTickerInfo
	XXMRXXBT PairTickerInfo
	XZECZUSD PairTickerInfo
	XZECZEUR PairTickerInfo
	XZECXXBT PairTickerInfo
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
