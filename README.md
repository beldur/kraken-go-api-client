Kraken GO API Client
====================

[![build status](https://img.shields.io/travis/beldur/kraken-go-api-client/master.svg)](https://travis-ci.org/beldur/kraken-go-api-client)

A simple API Client for the [Kraken](https://www.kraken.com/ "Kraken") Trading platform.

Example usage:

```go
package main

import (
	"fmt"
	"log"

	"github.com/beldur/kraken-go-api-client"
)

func main() {
	api := krakenapi.New("KEY", "SECRET")
	result, err := api.Query("Ticker", map[string]string{
		"pair": "XXBTZEUR",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %+v\n", result)

	// There are also some strongly typed methods available
	ticker, err := api.Ticker(krakenapi.XXBTZEUR)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ticker.XXBTZEUR.OpeningPrice)
}
```

## Contributors
 - Piega
 - Glavic
 - MarinX
 - bjorand
 - [khezen](https://github.com/khezen)
 