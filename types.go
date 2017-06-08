package krakenapi

const (
	DASHEUR = "DASHEUR"
	DASHUSD = "DASHUSD"
	DASHXBT = "DASHXBT"
	GNOETH = "GNOETH"
	GNOEUR = "GNOEUR"
	GNOUSD = "GNOUSD"
	GNOXBT = "GNOXBT"
	USDTZUSD = "USDTZUSD"
	XETCXETH = "XETCXETH"
	XETCXXBT = "XETCXXBT"
	XETCXEUR = "XETCXEUR"
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
	USDTZUSD AssetPairInfo
	XETCXETH AssetPairInfo
	XETCXXBT AssetPairInfo
	XETCXEUR AssetPairInfo
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
	DASHEUR PairTickerInfo
	DASHUSD PairTickerInfo
	DASHXBT PairTickerInfo
	GNOETH PairTickerInfo
	GNOEUR PairTickerInfo
	GNOUSD PairTickerInfo
	GNOXBT PairTickerInfo
	USDTZUSD PairTickerInfo
	XETCXETH PairTickerInfo
	XETCXXBT PairTickerInfo
	XETCXEUR PairTickerInfo
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
