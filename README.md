Kraken GO API Client
====================

A simple API Client for the [Kraken](https://www.kraken.com/ "Kraken") Trading platform.

```go
package main

import (
	"fmt"
	"https://github.com/Beldur/kraken-go-api-client"
)

func main() {
	api := NewKrakenApi("KEY", "SECRET")
	result, err := api.Query("Ticker", map[string]string{
		"pair": "XXBTZEUR",
	})

	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Printf("Result: %+v\n", result)
```