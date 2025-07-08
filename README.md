Kraken GO API Client
====================

<p align="center">:sparkles: Please use the new official go-lang api client from Kraken: https://github.com/krakenfx/api-go :sparkles:</p>

A simple API Client for the [Kraken](https://www.kraken.com/ "Kraken") Trading platform.

Example usage:

```go
package main

import (
	"fmt"
	"log"

	"github.com/massigerardi/kraken-go-api-client"
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
 
