package birdeye

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// Optional query parameters for NewListing
type NewListingOpt struct {
	ToTime int
	Limit  int
}

func (b *birdeye) NewListing(ctx context.Context, toTime time.Time, opt *NewListingOpt) (result BirdeyeResponse[NewListing], err error) {
	params := querry{
		"time_to": strconv.Itoa(int(toTime.Unix())),
		"limit":   "10",
	}

	req := b.client.R().
		SetQueryParams(params).
		SetContext(ctx).
		SetResult(&result)

	err = b.call(req, http.MethodGet, "/defi/v2/tokens/new_listing")
	return
}

type TrandingListParam struct {
	SortBy   sortBy
	SortType sortType
	Offset   int
	Limit    int
}

func (b *birdeye) TrendingList(ctx context.Context, param TrandingListParam) (result BirdeyeResponse[TrendingList], err error) {
	if param.SortBy == "" || param.SortType == "" {
		return result, errors.New("missing required parameters: TrandingListParam.SortBy or TrandingListParam.SortType")
	}

	if param.Offset < 0 {
		return result, errors.New("invalid parameter: TrandingListParam.Offset")
	}

	if param.Limit == 0 {
		param.Limit = 10
	}

	req := b.client.R().
		SetQueryParams(map[string]string{
			"sort_by":   string(param.SortBy),
			"sort_type": string(param.SortType),
			"offset":    strconv.Itoa(param.Offset),
			"limit":     strconv.Itoa(param.Limit),
		}).
		SetContext(ctx).
		SetResult(&result)

	err = b.call(req, http.MethodGet, "/defi/token_trending")
	return
}
