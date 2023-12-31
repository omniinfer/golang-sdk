package request

import (
	"context"
	"github.com/omniinfer/golang-sdk/types"
	"net/http"
)

func (c *OmniClient) Progress(ctx context.Context, request *types.ProgressRequest, opts ...WithGenerateImageOption) (*types.ProgressResponse, error) {
	responseData, err := omniRequest[types.ProgressRequest, types.ProgressResponse](ctx, c.httpCli, http.MethodGet, BaseURL+"/progress", c.apiKey, map[string]interface{}{
		"task_id": request.TaskId,
	}, nil)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}
