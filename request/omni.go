package request

import (
	"context"
	"errors"
	"github.com/omniinfer/golang-sdk/model"
	"net/http"
	"time"
)

// OmniClientInterface define the behavior or OmniClient
type OmniClientInterface interface {
	// Txt2Img Asynchronously generate images from request. Returns AsyncResponse, use `task_id` to get Progress.
	Txt2Img(context.Context, *model.Txt2ImgRequest) (*model.AsyncResponse, error)
	// SyncTxt2img Synchronously generate images from request. You can get image url in `Progress.Data.Imgs`.
	SyncTxt2img(context.Context, *model.Txt2ImgRequest, ...WithGenerateImageOption) (*model.ProgressResponse, error)
	// Img2Img  Asynchronously generate images from request. Returns AsyncResponse, use `task_id` to get Progress.
	Img2Img(context.Context, *model.Img2ImgRequest) (*model.AsyncResponse, error)
	// SyncImg2img Synchronously generate images from request. You can get image url in `Progress.Data.Imgs`.
	SyncImg2img(context.Context, *model.Img2ImgRequest, ...WithGenerateImageOption) (*model.ProgressResponse, error)
	// Progress Task Progress, use `task_id` to get progress.
	Progress(context.Context, *model.ProgressRequest, ...WithGenerateImageOption) (*model.ProgressResponse, error)
	// Upscale Asynchronously upscale images from request. Returns AsyncResponse, use `task_id` to get Progress.
	Upscale(context.Context, *model.Img2ImgRequest) (*model.AsyncResponse, error)
	// SyncUpscale Synchronously upscale images from request. You can get image url in `Progress.Data.Imgs`.
	SyncUpscale(context.Context, *model.Img2ImgRequest, ...WithGenerateImageOption) (*model.ProgressResponse, error)
	// Models List all models, including checkpoint, lora, vae and other models. Return model info by type.
	Models(context.Context, ...WithModelOption) (map[model.ModelType]*model.Model, error)
}

const BaseURL = "https://api.omniinfer.io/v2"

type OmniClient struct {
	apiKey     string
	httpCli    *http.Client
	modelCache model.ModelList
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
