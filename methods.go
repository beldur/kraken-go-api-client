package krakenapi

import (
	"errors"
	"strconv"
)

// Params for public method Assets
var AssetsParams = map[string]bool{
	"info":   false, // info to retrieve (optional): info = all info (default)
	"aclass": false, // asset class (optional): currency (default)
	"asset":  false, // comma delimited list of assets to get info on (optional.  default = all for given asset class)
}

// Params for public method AssetPairs
var AssetPairsParams = map[string]bool{
	"info": false, // info to retrieve (optional): info = all info (default) | leverage = leverage info | fees = fees schedule | margin = margin info
	"pair": false, // comma delimited list of asset pairs to get info on (optional.  default = all)
}

// Params for public method Ticker
var TickerParams = map[string]bool{
	"pair": true, // comma delimited list of asset pairs to get info on
}

// Params for public method OHLC
var OHLCParams = map[string]bool{
	"pair":     true,  // asset pair to get OHLC data for
	"interval": false, // time frame interval in minutes (optional): 1 (default), 5, 15, 30, 60, 240, 1440, 10080, 21600
	"since":    false, // return committed OHLC data since given id (optional.  exclusive)
}

// Params for public method Depth
var DepthParams = map[string]bool{
	"pair":  true,  // asset pair to get market depth for
	"count": false, // maximum number of asks/bids (optional)
}

// Params for public method Trades
var TradesParams = map[string]bool{
	"pair":  true,  // asset pair to get market depth for
	"since": false, // return trade data since given id (optional.  exclusive)
}

// Params for public method Spread
var SpreadParams = map[string]bool{
	"pair":  true,  // asset pair to get spread data for
	"since": false, // return spread data since given id (optional.  inclusive)
}

// Params for public method TradeBalance
var TradeBalanceParams = map[string]bool{
	"aclass": false, // asset class (optional): currency (default)
	"asset":  false, // base asset used to determine balance (default = ZUSD)
}

// Params for private method OpenOrders
var OpenOrdersParams = map[string]bool{
	"trades":  false, // whether or not to include trades in output (optional.  default = false)
	"userref": false, // restrict results to given user reference id (optional)
}

// Params for private method ClosedOrders
var ClosedOrdersParams = map[string]bool{
	"trades":    false, // whether or not to include trades in output (optional.  default = false)
	"userref":   false, // restrict results to given user reference id (optional)
	"start":     false, // starting unix timestamp or order tx id of results (optional.  exclusive)
	"end":       false, // ending unix timestamp or order tx id of results (optional.  inclusive)
	"ofs":       false, // result offset
	"closetime": false, // which time to use (optional): open | close | both (default)
}

// Params for private method QueryOrders
var QueryOrdersParams = map[string]bool{
	"trades":  false, // whether or not to include trades in output (optional.  default = false)
	"userref": false, // restrict results to given user reference id (optional)
	"txid":    false, // comma delimited list of transaction ids to query info about (20 maximum)
}

// Params for private method TradesHistory
var TradesHistoryParams = map[string]bool{}

// Params for private method QueryTrades
var QueryTradesParams = map[string]bool{}

// Params for private method OpenPositions
var OpenPositionsParams = map[string]bool{}

// Params for private method Ledgers
var LedgersParams = map[string]bool{}

// Params for private method QueryLedgers
var QueryLedgersParams = map[string]bool{}

// Params for private method TradeVolume
var TradeVolumeParams = map[string]bool{}

// Params for private method AddOrder
var AddOrderParams = map[string]bool{
	"pair": false, // asset pair
	"type": false, // type of order (buy/sell)
	// .............. TODO: add all the other!
}

// Params for private method CancelOrder
var CancelOrderParams = map[string]bool{
	"txid": false, // transaction id
}

// Time returns the server's time
// https://www.kraken.com/help/api#get-server-time
func (api *KrakenApi) Time() (*TimeResponse, error) {
	resp, err := api.queryPublic("Time", nil, &TimeResponse{})
	if err != nil {
		return nil, err
	}

	return resp.(*TimeResponse), nil
}

// Assets returns the servers available assets
// https://www.kraken.com/help/api#get-asset-info
func (api *KrakenApi) Assets(args map[string]string) (*AssetsResponse, error) {
	values, err := prepareValues(AssetsParams, args)
	if err != nil {
		return nil, err
	}

	resp, err := api.queryPublic("Assets", values, &AssetsResponse{})
	if err != nil {
		return nil, err
	}

	return resp.(*AssetsResponse), nil
}

// AssetPairs returns the servers available asset pairs
// https://www.kraken.com/help/api#get-tradable-pairs
func (api *KrakenApi) AssetPairs(args map[string]string) (*AssetPairsResponse, error) {
	values, err := prepareValues(AssetPairsParams, args)
	if err != nil {
		return nil, err
	}

	resp, err := api.queryPublic("AssetPairs", values, &AssetPairsResponse{})
	if err != nil {
		return nil, err
	}

	return resp.(*AssetPairsResponse), nil
}

// Ticker returns the ticker for given comma separated pairs
// https://www.kraken.com/help/api#get-ticker-info
func (api *KrakenApi) Ticker(args map[string]string) (*TickerResponse, error) {
	values, err := prepareValues(TickerParams, args)
	if err != nil {
		return nil, err
	}

	resp, err := api.queryPublic("Ticker", values, &TickerResponse{})
	if err != nil {
		return nil, err
	}

	return resp.(*TickerResponse), nil
}

// OHLC
// https://www.kraken.com/help/api#get-ohlc-data

// Depth returns the order book for given pair and orders count.
// https://www.kraken.com/help/api#get-order-book
func (api *KrakenApi) Depth(args map[string]string) (*OrderBook, error) {
	values, err := prepareValues(DepthParams, args)
	if err != nil {
		return nil, err
	}

	dr := DepthResponse{}
	_, err = api.queryPublic("Depth", values, &dr)
	if err != nil {
		return nil, err
	}

	if book, found := dr[args["pair"]]; found {
		return &book, nil
	}

	return nil, errors.New("invalid response")
}

// Trades returns the recent trades for given pair
// https://www.kraken.com/help/api#get-recent-trades
func (api *KrakenApi) Trades(args map[string]string) (*TradesResponse, error) {
	values, err := prepareValues(TradesParams, args)
	if err != nil {
		return nil, err
	}

	resp, err := api.queryPublic("Trades", values, nil)
	if err != nil {
		return nil, err
	}

	v := resp.(map[string]interface{})

	last, err := strconv.ParseInt(v["last"].(string), 10, 64)
	if err != nil {
		return nil, err
	}

	result := &TradesResponse{
		Last:   last,
		Trades: make([]TradeInfo, 0),
	}

	trades := v[args["pair"]].([]interface{})
	for _, v := range trades {
		trade := v.([]interface{})

		priceString := trade[0].(string)
		price, _ := strconv.ParseFloat(priceString, 64)

		volumeString := trade[1].(string)
		volume, _ := strconv.ParseFloat(trade[1].(string), 64)

		tradeInfo := TradeInfo{
			Price:         priceString,
			PriceFloat:    price,
			Volume:        volumeString,
			VolumeFloat:   volume,
			Time:          int64(trade[2].(float64)),
			Buy:           trade[3].(string) == BUY,
			Sell:          trade[3].(string) == SELL,
			Market:        trade[4].(string) == MARKET,
			Limit:         trade[4].(string) == LIMIT,
			Miscellaneous: trade[5].(string),
		}

		result.Trades = append(result.Trades, tradeInfo)
	}

	return result, nil
}

// Spread
// https://www.kraken.com/help/api#get-recent-spread-data

// Balance returns all account asset balances
// https://www.kraken.com/help/api#get-account-balance
func (api *KrakenApi) Balance() (*BalanceResponse, error) {
	resp, err := api.queryPrivate("Balance", nil, &BalanceResponse{})
	if err != nil {
		return nil, err
	}

	return resp.(*BalanceResponse), nil
}

// TradeBalance
// https://www.kraken.com/help/api#get-trade-balance

// OpenOrders returns all open orders
// https://www.kraken.com/help/api#get-open-orders
func (api *KrakenApi) OpenOrders(args map[string]string) (*OpenOrdersResponse, error) {
	values, err := prepareValues(OpenOrdersParams, args)
	if err != nil {
		return nil, err
	}

	resp, err := api.queryPublic("OpenOrders", values, &OpenOrdersResponse{})
	if err != nil {
		return nil, err
	}

	return resp.(*OpenOrdersResponse), nil
}

// ClosedOrders returns all closed orders
// https://www.kraken.com/help/api#get-closed-orders
func (api *KrakenApi) ClosedOrders(args map[string]string) (*ClosedOrdersResponse, error) {
	values, err := prepareValues(ClosedOrdersParams, args)
	if err != nil {
		return nil, err
	}

	resp, err := api.queryPrivate("ClosedOrders", values, &ClosedOrdersResponse{})
	if err != nil {
		return nil, err
	}

	return resp.(*ClosedOrdersResponse), nil
}

// QueryOrders shows order
// https://www.kraken.com/help/api#query-orders-info
func (api *KrakenApi) QueryOrders(args map[string]string) (*QueryOrdersResponse, error) {
	values, err := prepareValues(QueryOrdersParams, args)
	if err != nil {
		return nil, err
	}

	resp, err := api.queryPrivate("QueryOrders", values, &QueryOrdersResponse{})
	if err != nil {
		return nil, err
	}

	return resp.(*QueryOrdersResponse), nil
}

// TradesHistory
// https://www.kraken.com/help/api#get-trades-history

// QueryTrades
// https://www.kraken.com/help/api#query-trades-info

// OpenPositions
// https://www.kraken.com/help/api#get-open-positions

// Ledgers
// https://www.kraken.com/help/api#get-ledgers-info

// QueryLedgers
// https://www.kraken.com/help/api#query-ledgers

// TradeVolume
// https://www.kraken.com/help/api#get-trade-volume

// AddOrder adds new order
// https://www.kraken.com/help/api#add-standard-order
func (api *KrakenApi) AddOrder(args map[string]string) (*AddOrderResponse, error) {
	values, err := prepareValues(AddOrderParams, args)
	if err != nil {
		return nil, err
	}

	resp, err := api.queryPrivate("AddOrder", values, &AddOrderResponse{})
	if err != nil {
		return nil, err
	}

	return resp.(*AddOrderResponse), nil
}

// CancelOrder cancels order
// https://www.kraken.com/help/api#cancel-open-order
func (api *KrakenApi) CancelOrder(args map[string]string) (*CancelOrderResponse, error) {
	values, err := prepareValues(CancelOrderParams, args)
	if err != nil {
		return nil, err
	}

	resp, err := api.queryPrivate("CancelOrder", values, &CancelOrderResponse{})
	if err != nil {
		return nil, err
	}

	return resp.(*CancelOrderResponse), nil
}
