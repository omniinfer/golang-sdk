package request

import (
	"context"
	"errors"
	"github.com/omniinfer/golang-sdk/types"
	"net/http"
	"time"
)

// OmniClientInterface define the behavior or OmniClient
type OmniClientInterface interface {
	// Txt2Img Asynchronously generate images from request. Returns AsyncResponse, use `task_id` to get Progress.
	Txt2Img(context.Context, *types.Txt2ImgRequest) (*types.AsyncResponse, error)
	// SyncTxt2img Synchronously generate images from request. You can get image url in `Progress.Data.Imgs`.
	SyncTxt2img(context.Context, *types.Txt2ImgRequest, ...WithGenerateImageOption) (*types.ProgressResponse, error)
	// Img2Img  Asynchronously generate images from request. Returns AsyncResponse, use `task_id` to get Progress.
	Img2Img(context.Context, *types.Img2ImgRequest) (*types.AsyncResponse, error)
	// SyncImg2img Synchronously generate images from request. You can get image url in `Progress.Data.Imgs`.
	SyncImg2img(context.Context, *types.Img2ImgRequest, ...WithGenerateImageOption) (*types.ProgressResponse, error)
	// Progress Task Progress, use `task_id` to get progress.
	Progress(context.Context, *types.ProgressRequest, ...WithGenerateImageOption) (*types.ProgressResponse, error)
	// Upscale Asynchronously upscale images from request. Returns AsyncResponse, use `task_id` to get Progress.
	Upscale(context.Context, *types.Img2ImgRequest) (*types.AsyncResponse, error)
	// SyncUpscale Synchronously upscale images from request. You can get image url in `Progress.Data.Imgs`.
	SyncUpscale(context.Context, *types.Img2ImgRequest, ...WithGenerateImageOption) (*types.ProgressResponse, error)
	// Models List all models, including checkpoint, lora, vae and other models. Return types info by type.
	Models(context.Context, ...WithModelOption) (map[types.ModelType]*types.Model, error)
}

const BaseURL = "https://api.omniinfer.io/v2"

type OmniClient struct {
	apiKey     string
	httpCli    *http.Client
	modelCache types.ModelList
}

func NewOmniClient(apiKey string) (*OmniClient, error) {
	if apiKey == "" {
		return nil, errors.New("apiKey is not set, you can get api key refer to https://docs.omniinfer.io/get-started")
	}
	client := &OmniClient{
		apiKey: apiKey,
		httpCli: &http.Client{
			Timeout: 30 * time.Second,
		},
		modelCache: nil,
	}
	return client, nil
}
