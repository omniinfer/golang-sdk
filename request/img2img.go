package request

import (
	"context"
	"github.com/omniinfer/golang-sdk/types"
	"net/http"
)

func (c *OmniClient) Img2Img(ctx context.Context, request *types.Img2ImgRequest) (*types.AsyncResponse, error) {
	responseData, err := omniRequest[types.Img2ImgRequest, types.AsyncResponse](ctx, c.httpCli, http.MethodPost, BaseURL+"/img2img", c.apiKey, nil, request)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}

func (c *OmniClient) SyncImg2img(ctx context.Context, request *types.Img2ImgRequest, opts ...WithGenerateImageOption) (*types.ProgressResponse, error) {
	return omniSyncImageGeneration[*types.Img2ImgRequest](ctx, request, opts, c.Img2Img, c.waitForTask)
}
