package request

import (
	"context"
	"github.com/omniinfer/golang-sdk/model"
	"net/http"
)

func (c *OmniClient) Progress(ctx context.Context, request *model.ProgressRequest, opts ...WithGenerateImageOption) (*model.ProgressResponse, error) {
	responseData, err := omniRequest[model.ProgressRequest, model.ProgressResponse](ctx, c.httpCli, http.MethodGet, BaseURL+"/progress", c.apiKey, map[string]interface{}{
		"task_id": request.TaskId,
	}, nil)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}
