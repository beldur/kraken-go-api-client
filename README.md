Kraken GO API Client
====================

A simple API Client for the [Kraken](https://www.kraken.com/ "Kraken") Trading platform.

Example usage:

```go
package main

import (
	"fmt"
	"github.com/Beldur/kraken-go-api-client"
)

func main() {
	api := krakenapi.New("KEY", "SECRET")
	result, err := api.Query("Ticker", map[string]string{
		"pair": "XXBTZEUR",
	})

	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Printf("Result: %+v\n", result)
}
```

If you find this useful, you can send me a fraction of a bitcoin!

1Q3P96LcTkbS9VwZkV5ndQa6t4EcR4GzSL
