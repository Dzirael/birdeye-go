package birdeye

type (
	sortBy     string
	sortType   string
	chain      string
	timeUpdate string
	querry     map[string]string
)

var (
	SortByRank      sortBy = "rank"
	SortByLiquidity sortBy = "liquidity"
	SortByVolume24H sortBy = "volume24hUSD"

	SortTypeAsc  sortType = "asc"
	SortTypeDesc sortType = "desc"

	H1  timeUpdate = "1h"
	H4  timeUpdate = "4h"
	H8  timeUpdate = "8h"
	H24 timeUpdate = "24h"
)

var (
	Solana    chain = "solana"
	Ethereum  chain = "ethereum"
	Arbitrum  chain = "arbitrum"
	Avalanche chain = "avalanche"
	Bsc       chain = "bsc"
	Optimism  chain = "optimism"
	Polygon   chain = "polygon"
	Base      chain = "base"
	ZkSync    chain = "zksync"
	Sui       chain = "sui"
)

// BirdeyeResponse is a generic response from Birdeye API
type BirdeyeResponse[T any] struct {
	Data    T    `json:"data"`
	Success bool `json:"success"`
}

type SupportedNetworks []chain

// https://docs.birdeye.so/reference/get_defi-price
type Price struct {
	Value           float64 `json:"value"`
	UpdateUnixTime  int     `json:"updateUnixTime"`
	UpdateHumanTime string  `json:"updateHumanTime"`
	Liquidity       float64 `json:"liquidity"`
}

type PriceMultiple map[string]Price

type PriceHistoricalUnix struct {
	Value          float64 `json:"value"`
	UpdateUnixTime int     `json:"updateUnixTime"`
	PriceChange24H float64 `json:"priceChange24h"`
}

type Trade struct {
	Quote         TokenData   `json:"quote"`
	Base          TokenData   `json:"base"`
	BasePrice     interface{} `json:"basePrice"`
	QuotePrice    interface{} `json:"quotePrice"`
	TxHash        string      `json:"txHash"`
	Source        string      `json:"source"`
	BlockUnixTime int         `json:"blockUnixTime"`
	TxType        string      `json:"txType"`
	Owner         string      `json:"owner"`
	Side          string      `json:"side"`
	Alias         interface{} `json:"alias"`
	PricePair     float64     `json:"pricePair"`
	From          TokenData   `json:"from"`
	To            TokenData   `json:"to"`
	TokenPrice    interface{} `json:"tokenPrice"`
	PoolID        string      `json:"poolId"`
}

type TokenData struct {
	Symbol         string      `json:"symbol"`
	Decimals       int         `json:"decimals"`
	Address        string      `json:"address"`
	Amount         int         `json:"amount"`
	FeeInfo        interface{} `json:"feeInfo"`
	UIAmount       float64     `json:"uiAmount"`
	Price          interface{} `json:"price"`
	NearestPrice   float64     `json:"nearestPrice"`
	ChangeAmount   int         `json:"changeAmount"`
	UIChangeAmount float64     `json:"uiChangeAmount"`
}

// https://docs.birdeye.so/reference/get_defi-token-trending
type TrendingList struct {
	UpdateUnixTime int     `json:"updateUnixTime"`
	UpdateTime     string  `json:"updateTime"`
	Tokens         []token `json:"tokens"`
	Total          int     `json:"total"`
}

type token struct {
	Address      string  `json:"address"`
	Decimals     int     `json:"decimals"`
	Liquidity    float64 `json:"liquidity"`
	LogoURI      string  `json:"logoURI"`
	Name         string  `json:"name"`
	Symbol       string  `json:"symbol"`
	Volume24HUSD float64 `json:"volume24hUSD"`
	Rank         int     `json:"rank"`
	Price        float64 `json:"price"`
}

// https://docs.birdeye.so/reference/get_defi-v2-tokens-new-listing
type NewListing struct {
	Items []listingItem `json:"items"`
}

type listingItem struct {
	Address          string      `json:"address"`
	Symbol           string      `json:"symbol"`
	Name             string      `json:"name"`
	Decimals         int         `json:"decimals"`
	LiquidityAddedAt string      `json:"liquidityAddedAt"`
	LogoURI          interface{} `json:"logoURI"`
	Liquidity        float64     `json:"liquidity"`
}
