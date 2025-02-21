package main

import (
	"fmt"
	"time"

	"github.com/zivlakmilos/rstax/api"
)

func main() {
	val, err := api.GetExchangeRate("USD", time.Now())
	if err != nil {
		fmt.Printf("error: %s", err)
	}

	fmt.Printf("%#v\n", val)
}
