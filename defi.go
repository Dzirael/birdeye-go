package birdeye

import (
	"context"
	"net/http"
	"strconv"
)

func (b *birdeye) SupportedNetworks(ctx context.Context) (result BirdeyeResponse[SupportedNetworks], err error) {
	req := b.client.R().
		SetContext(ctx).
		SetResult(&result)

	err = b.call(req, http.MethodGet, "/birdeye/networks")
	return
}

// Optional query parameters for Price
type PriceOpt struct {
	CheckLiquidty    bool
	IncludeLiquidity bool
	Type             timeUpdate
}

func (b *birdeye) Price(ctx context.Context, address string, opt *PriceOpt) (result BirdeyeResponse[Price], err error) {
	params := querry{
		"address": address,
	}

	if opt != nil {
		params["check_liquidity"] = strconv.FormatBool(opt.CheckLiquidty)
		params["include_liquidity"] = strconv.FormatBool(opt.IncludeLiquidity)
	}

	req := b.client.R().
		SetQueryParams(params).
		SetContext(ctx).
		SetResult(&result)

	err = b.call(req, http.MethodGet, "/birdeye/price")
	return
}

func (b *birdeye) PriceMultipleGet(ctx context.Context, addresses []string, opt *PriceOpt) (result BirdeyeResponse[PriceMultiple], err error) {
	params := querry{
		"list_address": toString(addresses),
	}

	if opt != nil {
		params["check_liquidity"] = strconv.FormatBool(opt.CheckLiquidty)
		params["include_liquidity"] = strconv.FormatBool(opt.IncludeLiquidity)
	}

	req := b.client.R().
		SetQueryParams(params).
		SetContext(ctx).
		SetResult(&result)

	err = b.call(req, http.MethodGet, "/birdeye/multi_price")
	return
}

func (b *birdeye) PriceMultiplePost(ctx context.Context, addresses []string, opt *PriceOpt) (result BirdeyeResponse[PriceMultiple], err error) {
	params := querry{
		"list_address": toString(addresses),
	}

	if opt != nil {
		params["check_liquidity"] = strconv.FormatBool(opt.CheckLiquidty)
		params["include_liquidity"] = strconv.FormatBool(opt.IncludeLiquidity)
		params["type"] = string(opt.Type)
	}

	req := b.client.R().
		SetBody(params).
		SetContext(ctx).
		SetResult(&result)

	err = b.call(req, http.MethodPost, "/birdeye/multi_price")
	return
}

type PriceHistoricalOpt struct {
	Address     string `json:"address"`
	AddressType string `json:"address_type"`
	Type        string `json:"type"`
	TimeFrom    int    `json:"time_from"`
	TimeTo      int    `json:"time_to"`
}

func (b *birdeye) PriceHistorical(ctx context.Context, opt PriceHistoricalOpt) (result BirdeyeResponse[Price], err error) {
	req := b.client.R().
		SetBody(opt).
		SetContext(ctx).
		SetResult(&result)

	err = b.call(req, http.MethodGet, "/birdeye/history_price")
	return
}

type PriceHistoricalUnixOpt struct {
	UnixTime int
}

func (b *birdeye) PriceHistoricalUnix(ctx context.Context, address string, opt *PriceHistoricalUnixOpt) (result BirdeyeResponse[PriceHistoricalUnix], err error) {
	params := querry{
		"address": address,
	}

	if opt != nil {
		params["unixtime"] = strconv.Itoa(opt.UnixTime)
	}

	req := b.client.R().
		SetQueryParams(params).
		SetContext(ctx).
		SetResult(&result)

	err = b.call(req, http.MethodGet, "/birdeye/historical_price_unix")
	return
}

type Pagination struct {
	Offset int
	Limit  int
}

func (b *birdeye) TokenTrades(ctx context.Context, address, tx_type string, sort sortType, opt *Pagination) (result BirdeyeResponse[Trade], err error) {
	params := querry{
		"address":   address,
		"tx_type":   tx_type,
		"sort_type": string(sort),
	}

	if opt != nil {
		params["offset"] = strconv.Itoa(opt.Offset)
		params["limit"] = strconv.Itoa(opt.Limit)
	}

	req := b.client.R().
		SetQueryParams(params).
		SetContext(ctx).
		SetResult(&result)

	err = b.call(req, http.MethodGet, "/birdeye/txs/token")
	return
}

func (b *birdeye) PairTrades(ctx context.Context, address, tx_type string, sort sortType, opt *Pagination) (result BirdeyeResponse[Trade], err error) {
	params := querry{
		"address":   address,
		"tx_type":   tx_type,
		"sort_type": string(sort),
	}

	if opt != nil {
		params["offset"] = strconv.Itoa(opt.Offset)
		params["limit"] = strconv.Itoa(opt.Limit)
	}

	req := b.client.R().
		SetQueryParams(params).
		SetContext(ctx).
		SetResult(&result)

	err = b.call(req, http.MethodGet, "/birdeye/txs/token")
	return
}
