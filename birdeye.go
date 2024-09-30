package birdeye

import "context"

type Birdeye interface {
	// Defi APIs

	// SupportedNetworks retrieves a list of all supported networks from the Birdeye API.
	// It accepts a context (`ctx`) to allow for cancellation and timeout control during the API call.
	//
	// Parameters:
	//   - ctx: context.Context - used to control the lifecycle of the API request
	//
	// Returns:
	//   - BirdeyeResponse[SupportedNetworks]: response containing the supported networks information
	//   - error: any error encountered during the API request
	//
	// Example usage:
	//   networks, err := birdeye.SupportedNetworks(ctx)
	//   if err != nil {
	//       log.Fatalf("failed to retrieve networks: %v", err)
	//   }
	//   fmt.Printf("Supported networks: %+v\n", networks)
	SupportedNetworks(ctx context.Context) (BirdeyeResponse[SupportedNetworks], error)

	// Price retrieves the latest price update for a given token from the Birdeye API.
	// The token is identified by its address, and an optional `PriceOpt` parameter can be provided
	// to customize the query (e.g., checking liquidity).
	//
	// Parameters:
	//   - ctx: context.Context - used to control the lifecycle of the API request
	//   - address: string - the address of the token for which the price is requested
	//   - opt: *PriceOpt - optional parameters to customize the request (e.g., liquidity options)
	//
	// Returns:
	//   - BirdeyeResponse[Price]: response containing the price information of the token
	//   - error: any error encountered during the API request
	//
	// Example usage:
	//   price, err := birdeye.Price(ctx, "So11111111111111111111111111111111111111112", &PriceOpt{
	//       CheckLiquidity:    true,
	//       IncludeLiquidity:  true,
	//   })
	//   if err != nil {
	//       log.Fatalf("failed to retrieve token price: %v", err)
	//   }
	//   fmt.Printf("Token price: %+v\n", price)
	Price(ctx context.Context, address string, opt *PriceOpt) (BirdeyeResponse[Price], error)

	// PriceMultipleGet retrieves price updates for multiple tokens in a single API call using the Birdeye API.
	// You can request prices for up to 100 tokens at once by providing their addresses.
	// An optional `PriceOpt` parameter allows customization of the query (e.g., checking liquidity).
	//
	// Parameters:
	//   - ctx: context.Context - used to control the lifecycle of the API request
	//   - addresses: []string - a slice of token addresses (maximum of 100) for which prices are requested
	//   - opt: *PriceOpt - optional parameters to customize the request (e.g., liquidity options)
	//
	// Returns:
	//   - BirdeyeResponse[PriceMultiple]: response containing price information for the requested tokens
	//   - error: any error encountered during the API request
	//
	// Example usage:
	//   prices, err := birdeye.PriceMultipleGet(ctx, []string{
	//       "So11111111111111111111111111111111111111112",
	//       "mSoLzYCxHdYgdzU16g5QSh3i5K3z3KZK7ytfqcJm7So"}, &PriceOpt{
	//       CheckLiquidity:    true,
	//       IncludeLiquidity:  true,
	//   })
	//   if err != nil {
	//       log.Fatalf("failed to retrieve multiple token prices: %v", err)
	//   }
	//   fmt.Printf("Token prices: %+v\n", prices)
	PriceMultipleGet(ctx context.Context, addresses []string, opt *PriceOpt) (BirdeyeResponse[PriceMultiple], error)

	// PriceMultiplePost retrieves price updates for multiple tokens in a single API call using the Birdeye API (via POST request).
	// You can request prices for up to 100 tokens at once by providing their addresses.
	// An optional `PriceOpt` parameter can be provided to customize the query, including options such as checking liquidity and selecting the price type.
	//
	// Parameters:
	//   - ctx: context.Context - used to control the lifecycle of the API request
	//   - addresses: []string - a slice of token addresses (maximum of 100) for which prices are requested
	//   - opt: *PriceOpt - optional parameters to customize the request (e.g., liquidity options, price type)
	//
	// Returns:
	//   - BirdeyeResponse[PriceMultiple]: response containing price information for the requested tokens
	//   - error: any error encountered during the API request
	//
	// Example usage:
	//   prices, err := birdeye.PriceMultiplePost(ctx, []string{
	//       "So11111111111111111111111111111111111111112",
	//       "mSoLzYCxHdYgdzU16g5QSh3i5K3z3KZK7ytfqcJm7So"}, &PriceOpt{
	//       CheckLiquidity:    true,
	//       IncludeLiquidity:  true,
	//       Type:             H1,
	//   })
	//   if err != nil {
	//       log.Fatalf("failed to retrieve multiple token prices: %v", err)
	//   }
	//   fmt.Printf("Token prices: %+v\n", prices)
	PriceMultiplePost(ctx context.Context, addresses []string, opt *PriceOpt) (BirdeyeResponse[PriceMultiple], error)

	// PriceHistorical retrieves the historical price data of a token from the Birdeye API, typically for use in a line chart.
	// The historical data is fetched based on the parameters provided in the `PriceHistoricalOpt` struct.
	//
	// Parameters:
	//   - ctx: context.Context - used to control the lifecycle of the API request
	//   - opt: PriceHistoricalOpt - required options for customizing the historical price query (e.g., token address, time range)
	//
	// Returns:
	//   - BirdeyeResponse[Price]: response containing the historical price information of the token
	//   - error: any error encountered during the API request
	//
	// Example usage:
	//   historicalPrice, err := birdeye.PriceHistorical(ctx, PriceHistoricalOpt{
	//       Address: "So11111111111111111111111111111111111111112",
	//       Interval: "1h",
	//       StartTime: 1620000000,
	//       EndTime:   1623600000,
	//   })
	//   if err != nil {
	//       log.Fatalf("failed to retrieve historical price data: %v", err)
	//   }
	//   fmt.Printf("Historical token prices: %+v\n", historicalPrice)
	PriceHistorical(ctx context.Context, opt PriceHistoricalOpt) (result BirdeyeResponse[Price], err error)

	// PriceHistoricalUnix retrieves the historical price of a token for a specific Unix timestamp using the Birdeye API.
	// The request can be customized using optional parameters provided in the `PriceHistoricalUnixOpt` struct.
	//
	// Parameters:
	//   - ctx: context.Context - used to control the lifecycle of the API request
	//   - address: string - the token address for which historical price data is requested
	//   - opt: *PriceHistoricalUnixOpt - optional parameters, such as specifying a Unix timestamp to fetch the price at a specific time
	//
	// Returns:
	//   - BirdeyeResponse[PriceHistoricalUnix]: response containing the historical price information for the specified token and time
	//   - error: any error encountered during the API request
	//
	// Example usage:
	//   historicalPrice, err := birdeye.PriceHistoricalUnix(ctx, "So11111111111111111111111111111111111111112", &PriceHistoricalUnixOpt{
	//       UnixTime: 1634025600,  // Specify the Unix timestamp (e.g., 1634025600)
	//   })
	//   if err != nil {
	//       log.Fatalf("failed to retrieve historical price by Unix timestamp: %v", err)
	//   }
	//   fmt.Printf("Historical price: %+v\n", historicalPrice)
	PriceHistoricalUnix(ctx context.Context, address string, opt *PriceHistoricalUnixOpt) (result BirdeyeResponse[PriceHistoricalUnix], err error)

	// TokenTrades retrieves a list of trades for a specific token using the Birdeye API.
	// You can filter trades by type (e.g., "buy" or "sell") and specify sorting preferences.
	// Optional pagination parameters can be used to limit the number of results and set an offset.
	//
	// Parameters:
	//   - ctx: context.Context - used to control the lifecycle of the API request
	//   - address: string - the token address for which trades are requested
	//   - tx_type: string - the type of trade (e.g., "buy", "sell"), required
	//   - sort: sortType - the sorting order for trades (e.g., ascending or descending), required
	//   - opt: *Pagination - optional pagination settings to limit and offset the result set
	//
	// Returns:
	//   - BirdeyeResponse[Trade]: response containing trade data for the specified token
	//   - error: any error encountered during the API request
	//
	// Example usage:
	//   trades, err := birdeye.TokenTrades(ctx, "So11111111111111111111111111111111111111112", "buy", SortTypeDesc, &Pagination{
	//       Offset: 0,
	//       Limit:  10,
	//   })
	//   if err != nil {
	//       log.Fatalf("failed to retrieve token trades: %v", err)
	//   }
	//   fmt.Printf("Token trades: %+v\n", trades)
	TokenTrades(ctx context.Context, address, tx_type string, sort sortType, opt *Pagination) (result BirdeyeResponse[Trade], err error)

	// PairTrades retrieves a list of trades for a specific trading pair or market using the Birdeye API.
	// The function allows filtering by trade type (e.g., "buy", "sell") and sorting the results based on the specified criteria.
	// Optional pagination can be applied to control the number of results returned and the starting point.
	//
	// Parameters:
	//   - ctx: context.Context - manages the API request lifecycle
	//   - address: string - the trading pair or market address for which trades are requested
	//   - tx_type: string - the type of trade (e.g., "buy", "sell"), required
	//   - sort: sortType - the order in which trades are sorted (e.g., ascending or descending), required
	//   - opt: *Pagination - optional settings for result pagination (e.g., offset, limit)
	//
	// Returns:
	//   - BirdeyeResponse[Trade]: response containing trade data for the specified trading pair or market
	//   - error: any error encountered during the API request
	//
	// Example usage:
	//   trades, err := birdeye.PairTrades(ctx, "So11111111111111111111111111111111111111112", "buy", SortTypeDesc, &Pagination{
	//       Offset: 0,
	//       Limit:  10,
	//   })
	//   if err != nil {
	//       log.Fatalf("failed to retrieve pair trades: %v", err)
	//   }
	//   fmt.Printf("Pair trades: %+v\n", trades)
	PairTrades(ctx context.Context, address, tx_type string, sort sortType, opt *Pagination) (result BirdeyeResponse[Trade], err error)
}
