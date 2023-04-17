# Busha Commerce Golang SDK 

This is the official Go client library for accessing the Busha Commerce API.

## Installation ️
To install the package, use the command below.

```sh
go get github.com/bushaHQ/busha-commerce-go
```

## Usage
```go

package main

import (
    commerce "github.com/bushaHQ/busha-commerce-go"
    "encoding/json"
    "log"
    "os"
)

func main() {
    //You can get your Secret API Key from the Business Dashboard
    secretKey := "test_b748a900222829292222222222"
    
    //Initiate Client by passing your secret key and an optional 
    //param to override the default http client 
    commerceClient, err := commerce.New(secretKey, nil)
    if err != nil {
        log.Fatal(err)
        return
    }
    
    //Create Checkout with fixed price
    chargeCreated, err := commerceClient.Checkout.Create(&commerce.CheckoutRequest{
        Name:          "iPhone 14 Pro",
        Description:   "This is a test checkout to sell my iPhone 14",
        CheckoutType:  commerce.FixedPrice,
        RequestedInfo: []string{"name", "email", "phone"},
        LocalAmount:   800,
        LocalCurrency: "NGN",
    })
    if err != nil {
        log.Fatal(err)
        return
    }
    json.NewEncoder(os.Stdout).Encode(chargeCreated)
	
	
    //List Charges
    charges, err := commerceClient.Charge.List(commerce.ListParameters{Page: 1, Sort: "asc", Limit :10}))
    if err != nil {
        log.Fatal(err)
        return
    }

    json.NewEncoder(os.Stdout).Encode(charges)

}
```

## TODO
- [ ] Update Documentation