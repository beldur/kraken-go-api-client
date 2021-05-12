package krakenapi

import (
	"encoding/base64"
	"encoding/json"
	"net/url"
	"reflect"
	"testing"
)

var publicAPI = New("", "")

func TestKrakenApi(t *testing.T) {
	var kk interface{} = KrakenApi{
		key:    "key",
		secret: "secret",
	}

	name := reflect.TypeOf(kk).Name()
	if name != "KrakenAPI" {
		t.Errorf("Unexpected struct, got %s want %s", name, "KrakenAPI")
	}
}

func TestCreateSignature(t *testing.T) {
	expectedSig := "Uog0MyIKZmXZ4/VFOh0g1u2U+A0ohuK8oCh0HFUiHLE2Csm23CuPCDaPquh/hpnAg/pSQLeXyBELpJejgOftCQ=="
	urlPath := "/0/private/"
	secret, _ := base64.StdEncoding.DecodeString("SECRET")
	values := url.Values{
		"TestKey": {"TestValue"},
	}

	sig := createSignature(urlPath, values, secret)

	if sig != expectedSig {
		t.Errorf("Expected Signature to be %s, got: %s\n", expectedSig, sig)
	}
}

func TestTime(t *testing.T) {
	resp, err := publicAPI.Time()
	if err != nil {
		t.Errorf("Time() should not return an error, got %s", err)
	}

	if resp.Unixtime <= 0 {
		t.Errorf("Time() should return valid Unixtime, got %d", resp.Unixtime)
	}
}

func TestAssets(t *testing.T) {
	_, err := publicAPI.Assets()
	if err != nil {
		t.Errorf("Assets() should not return an error, got %s", err)
	}
}

func TestAssetPairs(t *testing.T) {
	resp, err := publicAPI.AssetPairs()
	if err != nil {
		t.Errorf("AssetPairs() should not return an error, got %s", err)
	}

	if resp.XXBTZEUR.Base+resp.XXBTZEUR.Quote != XXBTZEUR {
		t.Errorf("AssetPairs() should return valid response, got %+v", resp.XXBTZEUR)
	}
}

func TestTicker(t *testing.T) {
	resp, err := publicAPI.Ticker(XXBTZEUR, XXRPZEUR)
	if err != nil {
		t.Errorf("Ticker() should not return an error, got %s", err)
	}

	if resp.XXBTZEUR.OpeningPrice == 0 {
		t.Errorf("Ticker() should return valid OpeningPrice, got %+v", resp.XXBTZEUR.OpeningPrice)
	}
}

func TestOHLCWithInterval(t *testing.T) {
	resp, err := publicAPI.OHLCWithInterval(XXBTZEUR, "15")
	if err != nil {
		t.Errorf("OHLCWithInterval() should not return an error, got %s", err)
	}

	if resp.Pair == "" {
		t.Errorf("OHLCWithInterval() should return valid Pair, got %+v", resp.Pair)
	}
}

func TestOHLC(t *testing.T) {
	resp, err := publicAPI.OHLC(XXBTZEUR)
	if err != nil {
		t.Errorf("OHLC() should not return an error, got %s", err)
	}

	if resp.Pair == "" {
		t.Errorf("OHLC() should return valid Pair, got %+v", resp.Pair)
	}
}

func TestQueryTime(t *testing.T) {
	result, err := publicAPI.Query("Time", map[string]string{})
	resultKind := reflect.TypeOf(result).Kind()

	if err != nil {
		t.Errorf("Query should not return an error, got %s", err)
	}
	if resultKind != reflect.Map {
		t.Errorf("Query `Time` should return a Map, got: %s", resultKind)
	}
}

func TestQueryTicker(t *testing.T) {
	result, err := publicAPI.Query("Ticker", map[string]string{
		"pair": "XXBTZEUR",
	})
	resultKind := reflect.TypeOf(result).Kind()

	if err != nil {
		t.Errorf("Query should not return an error, got %s", err)
	}

	if resultKind != reflect.Map {
		t.Errorf("Query `Ticker` should return a Map, got: %s", resultKind)
	}
}

func TestQueryTrades(t *testing.T) {
	result, err := publicAPI.Trades(XXBTZEUR, 1495777604391411290)

	if err != nil {
		t.Errorf("Trades should not return an error, got %s", err)
	}

	if result.Last == 0 {
		t.Errorf("Returned parameter `last` should always have a value...")
	}

	if len(result.Trades) > 0 {
		for _, trade := range result.Trades {
			if trade.Buy == trade.Sell {
				t.Errorf("Trade should be buy or sell")
			}
			if trade.Market == trade.Limit {
				t.Errorf("Trade type should be market or limit")
			}
		}
	}
}

func TestQueryDepth(t *testing.T) {
	pair := "XETHZEUR"
	count := 10
	result, err := publicAPI.Depth(pair, count)
	if err != nil {
		t.Errorf("Depth should not return an error, got %s", err)
	}

	resultType := reflect.TypeOf(result)

	if resultType != reflect.TypeOf(&OrderBook{}) {
		t.Errorf("Depth should return an OrderBook, got %s", resultType)
	}

	if len(result.Asks) > count {
		t.Errorf("Asks length must be less than count , got %d > %d", len(result.Asks), count)
	}

	if len(result.Bids) > count {
		t.Errorf("Bids length must be less than count , got %d > %d", len(result.Bids), count)
	}
}

func TestUnmarshalClosedOrder(t *testing.T) {
	body := []byte(`
	{
		"error": [],
		"result": {
			"closed": {
				"AAAAAA-BBBBB-CCCCCC": {
					"refid": null,
					"userref": 1000000000,
					"status": "closed",
					"reason": null,
					"opentm": 1000000000.5194,
					"closetm": 1000000000.5288,
					"starttm": 0,
					"expiretm": 1200000000,
					"descr": {
						"pair": "ETHEUR",
						"type": "sell",
						"ordertype": "market",
						"price": "0",
						"price2": "0",
						"leverage": "none",
						"order": "sell 0.02000000 ETHEUR @ market",
						"close": ""
					},
					"vol": "0.02000000",
					"vol_exec": "0.02000000",
					"cost": "2.95",
					"fee": "0",
					"price": "147.55",
					"stopprice": "0.00000",
					"limitprice": "0.00000",
					"misc": "",
					"oflags": "fciq",
					"trades": [
						"DDDDDD-EEEEE-FFFFFF"
					]
				}
			},
			"count": 1
		}
	}
	`)

	var jsonData KrakenResponse
	if err := json.Unmarshal(body, &jsonData); err != nil {
		t.Error(err)
	}

	result, err := json.Marshal(jsonData.Result)
	if err != nil {
		t.Error(err)
	}

	var closedOrders ClosedOrdersResponse
	if err := json.Unmarshal(result, &closedOrders); err != nil {
		t.Error(err)
	}

	if closedOrders.Count != 1 {
		t.Error("count should be 1")
	}

	order, ok := closedOrders.Closed["AAAAAA-BBBBB-CCCCCC"]
	if !ok {
		t.Error("order id AAAAAA-BBBBB-CCCCCC should exist")
	}

	if order.TradesID[0] != "DDDDDD-EEEEE-FFFFFF" {
		t.Error("trades[0] id should be DDDDDD-EEEEE-FFFFFF")
	}

}
