package main

import (
	"context"
	"fmt"

	"github.com/Dzirael/birdeye-go"
)

func main() {

	bird := birdeye.New("API_KEY", birdeye.Solana)

	ctx := context.Background()
	networks, err := bird.SupportedNetworks(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Supported networks: %v\n", networks)
}
