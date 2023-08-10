package request

import (
	"context"
	"github.com/omniinfer/golang-sdk/types"
	"net/http"
)

func (c *OmniClient) Upscale(ctx context.Context, request *types.UpscaleRequest) (*types.AsyncResponse, error) {
	responseData, err := omniRequest[types.UpscaleRequest, types.AsyncResponse](ctx, c.httpCli, http.MethodPost, BaseURL+"/upscale", c.apiKey, nil, request)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}

func (c *OmniClient) SyncUpscale(ctx context.Context, request *types.UpscaleRequest, opts ...WithGenerateImageOption) (*types.ProgressResponse, error) {
	return omniSyncImageGeneration[*types.UpscaleRequest](ctx, request, opts, c.Upscale, c.waitForTask)
}
