package krakenapi

import (
	"encoding/base64"
	"github.com/jarcoal/httpmock"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

var publicAPI = New("", "")

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

func TestKrakenApi_TradeVolume(t *testing.T) {
	type fields struct {
		key    string
		secret string
		client *http.Client
	}
	type args map[string]string
	tests := []struct {
		name     string
		fields   fields
		args     args
		response string
		want     *TradeVolumeResponse
		wantErr  bool
	}{
		{
			name: "test_response",
			fields: fields{
				key:    "my_key",
				secret: "my_secret",
				client: http.DefaultClient,
			},
			args: args{
				"pair":     XXBTZEUR,
				"fee-info": "",
			},
			response: `{"error":[],"result":{"currency":"ZUSD","volume":"0.0000","fees":{"XXBTZEUR":{"fee":"0.2600","minfee":"0.1000","maxfee":"0.2600","nextfee":"0.2400","nextvolume":"50000.0000","tiervolume":"0.0000"}},"fees_maker":{"XXBTZEUR":{"fee":"0.1600","minfee":"0.0000","maxfee":"0.1600","nextfee":"0.1400","nextvolume":"50000.0000","tiervolume":"0.0000"}}}}`,
			want: &TradeVolumeResponse{
				Volume:   0.0000,
				Currency: "ZUSD",
				Fees: Fees{
					XXBTZEUR: FeeInfo{
						Fee:        0.2600,
						MinFee:     0.1000,
						MaxFee:     0.2600,
						NextFee:    0.2400,
						NextVolume: 50000.0000,
						TierVolume: 0.0000,
					},
				},
				FeesMaker: Fees{
					XXBTZEUR: FeeInfo{
						Fee:        0.1600,
						MinFee:     0.0000,
						MaxFee:     0.1600,
						NextFee:    0.1400,
						NextVolume: 50000.0000,
						TierVolume: 0.0000,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test_error",
			fields: fields{
				key:    "my_key",
				secret: "my_secret",
				client: http.DefaultClient,
			},
			args: args{
				"pair":     XXBTZEUR,
				"fee-info": "",
			},
			response: `{"error":["my_error"]}`,
			want:     nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		httpmock.Activate()
		t.Run(tt.name, func(t *testing.T) {
			api := &KrakenApi{
				key:    tt.fields.key,
				secret: tt.fields.secret,
				client: tt.fields.client,
			}
			httpmock.RegisterResponder("POST", "https://api.kraken.com/0/private/TradeVolume", func(req *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(200, tt.response)
				resp.Header.Add("Content-Type", "application/json")
				return resp, nil
			})
			got, err := api.TradeVolume(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("TradeVolume() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TradeVolume() got = %v, want %v", got, tt.want)
			}
		})
	}
}
