package busha_commerce_go

import (
	"log"
)

var c *Client

func init() {
	apiKey := mustHaveTestKeyEnv()
	c, _ = New(apiKey, nil)
	log.Println(apiKey)
}
