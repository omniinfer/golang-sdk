package request

import (
	"context"
	"github.com/omniinfer/golang-sdk/model"
	"net/http"
)

func (c *OmniClient) Models(ctx context.Context, opts ...WithModelOption) (model.ModelList, error) {
	modelOpt := newModelOption(opts...)
	if c.modelCache == nil || len(c.modelCache) == 0 || modelOpt.Refresh {
		responseData, err := omniRequest[*model.ModelRequest, model.ModelsResponse](ctx, c.httpCli, http.MethodGet, BaseURL+"/models", c.apiKey, nil, nil)
		if err != nil {
			return nil, err
		}
		c.modelCache = responseData.Data.Models
	}
	return c.modelCache, nil
}

type ModelOption struct {
	Refresh bool
}

func newModelOption(opts ...WithModelOption) *ModelOption {
	all := &ModelOption{}
	for _, opt := range opts {
		opt(all)
	}
	return all
}

type WithModelOption func(opt *ModelOption)

func WithRefresh() WithModelOption {
	return func(opt *ModelOption) {
		opt.Refresh = true
	}
}
