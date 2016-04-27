package krakenapi

import (
	"encoding/base64"
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
	resp, err := publicAPI.Ticker(XXBTZGBP, XXBTZEUR)
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
