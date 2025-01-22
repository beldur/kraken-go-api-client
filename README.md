# Kraken GO API Client

A simple API Client for the [Kraken](https://www.kraken.com/ "Kraken") Trading platform.

## Forked from github.com/beldur/kraken-go-api-client

The main changes to the original code is removal of hardcoded Assets, AssetPairs, etc. This eliminates the need of
manually adding new assets and pairs. Also, it allows for iterating over api responses without using reflection.

Example usage:

```go
package main

import (
	"fmt"
	"log"

	"github.com/henkvanramshorst/kraken-go-api-client"
)

func main() {
	api := krakenapi.New("KEY", "SECRET")
	resp, err := api.Query("Ticker", map[string]string{
		"pair": "XXBTZEUR",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Response: %+v\n", resp)

	// There are also some strongly typed methods available
	resp, err = api.Ticker("XXBTZEUR")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
```

## Contributors
 - Piega
 - Glavic
 - MarinX
 - bjorand
 - [khezen](https://github.com/khezen)
