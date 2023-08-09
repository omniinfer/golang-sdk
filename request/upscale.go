package request

import (
	"context"
	"github.com/omniinfer/golang-sdk/model"
	"net/http"
)

func (c *OmniClient) Upscale(ctx context.Context, request *model.UpscaleRequest) (*model.AsyncResponse, error) {
	responseData, err := omniRequest[model.UpscaleRequest, model.AsyncResponse](ctx, c.httpCli, http.MethodPost, BaseURL+"/upscale", c.apiKey, nil, request)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}

func (c *OmniClient) SyncUpscale(ctx context.Context, request *model.UpscaleRequest, opts ...WithGenerateImageOption) (*model.ProgressResponse, error) {
	return omniSyncImageGeneration[*model.UpscaleRequest](ctx, request, opts, c.Upscale, c.waitForTask)
}
