package request

import (
	"context"
	"github.com/omniinfer/golang-sdk/model"
	"net/http"
)

func (c *OmniClient) Img2Img(ctx context.Context, request *model.Img2ImgRequest) (*model.AsyncResponse, error) {
	responseData, err := omniRequest[model.Img2ImgRequest, model.AsyncResponse](ctx, c.httpCli, http.MethodPost, BaseURL+"/img2img", c.apiKey, nil, request)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}

func (c *OmniClient) SyncImg2img(ctx context.Context, request *model.Img2ImgRequest, opts ...WithGenerateImageOption) (*model.ProgressResponse, error) {
	return omniSyncImageGeneration[*model.Img2ImgRequest](ctx, request, opts, c.Img2Img, c.waitForTask)
}
