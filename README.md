go-coincheck
==========

go-coincheck is an implementation of the coincheck.com API in Golang.

It's just for me to handle coincheck for now, so now not support all of the APIs.

Based off of https://github.com/toorop/go-bittrex/

## Import
```
import "github.com/rakd/go-coincheck"
```

## Usage
~~~ go
package main

import (
	"fmt"
	"github.com/rakd/go-coincheck"
)

const (
	API_KEY    = "YOUR_API_KEY"
	API_SECRET = "YOUR_API_SECRET"
)

func main() {
	// coincheck client
	client := coincheck.New(API_KEY, API_SECRET)

	// Get balance
	balance, err := client.GetBalance()
	fmt.Println(err, balance)
}
~~~


## Stay tuned

- [Follow me on Twitter](https://twitter.com/kaz_lavender)

## Donate

- BTC: 1Ah8sarQ4w9FnsCs8LoG6JuYiFHmrAAy6F
