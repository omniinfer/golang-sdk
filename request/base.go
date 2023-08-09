package request

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	omni_error "github.com/omniinfer/golang-sdk/error"
	"github.com/omniinfer/golang-sdk/model"
	"io"
	"net/http"
	"net/url"
	"os"
)

// omniSyncImageGeneration Common omni sync image generation procedure
func omniSyncImageGeneration[RequestT any](ctx context.Context, request RequestT, opts []WithGenerateImageOption,
	async func(context.Context, RequestT) (*model.AsyncResponse, error),
	progress func(context.Context, *model.ProgressRequest, ...WithGenerateImageOption) (*model.ProgressResponse, error)) (*model.ProgressResponse, error) {
	// execute async generate image function, returns result that contains field `task_id`
	middleRsp, err := async(ctx, request)
	if err != nil {
		return nil, err
	}
	// use `task_id` to get images from progress
	progressRsp, err := progress(ctx, &model.ProgressRequest{
		TaskId: middleRsp.Data.TaskID,
	}, opts...)
	if err != nil {
		return nil, err
	}
	return progressRsp, nil
}

// GenerateImageOption Generation Option
type GenerateImageOption struct {
	DownloadImage              bool
	SaveImage                  bool
	SaveImageDir               string
	SaveImagePerm              os.FileMode
	SaveImageFileNameConverter func(taskId string, fileIndex int, fileName string) string
}

// WithGenerateImageOption Option Pattern
type WithGenerateImageOption func(*GenerateImageOption)

// WithDownloadImage if this is set, you can get image raw bytes in `Progress.Data.ImgsBytes`
func WithDownloadImage() WithGenerateImageOption {
	return func(opt *GenerateImageOption) {
		opt.DownloadImage = true
	}
}

func WithSaveImage(dir string, perm os.FileMode, filenameConvert func(taskId string, fileIndex int, fileName string) string) WithGenerateImageOption {
	return func(opt *GenerateImageOption) {
		// download first
		opt.DownloadImage = true
		opt.SaveImage = true
		opt.SaveImageDir = dir
		opt.SaveImagePerm = perm
		// identity
		if filenameConvert == nil {
			filenameConvert = func(taskId string, fileIndex int, fileName string) string {
				return fileName
			}
		}
		opt.SaveImageFileNameConverter = filenameConvert
	}
}

// NewGenerateImageOption Option Pattern
func newGenerateImageOption(opts ...WithGenerateImageOption) *GenerateImageOption {
	all := &GenerateImageOption{}
	for _, opt := range opts {
		opt(all)
	}
	return all
}

// omniRequest Omni common request procedure.
func omniRequest[RequestT any, ResponseT model.BasicResponse](ctx context.Context, httpCli *http.Client, method, omniApiUrl, apiKey string,
	query map[string]interface{}, reqObj *RequestT) (*ResponseT, error) {
	// compare with nil
	// build url
	u, err := url.Parse(omniApiUrl)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	for key, value := range query {
		q.Set(key, fmt.Sprintf("%v", value))
	}
	u.RawQuery = q.Encode()
	// build body
	var bodyReader io.Reader
	if reqObj != nil {
		bs, err := json.Marshal(reqObj)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewBuffer(bs)
	}
	// build request
	req, err := http.NewRequestWithContext(ctx, method, u.String(), bodyReader)
	if err != nil {
		return nil, err
	}
	// build header
	headers := map[string]string{
		"Accept":          "application/json",
		"Content-Type":    "application/json",
		"X-Omni-Key":      apiKey,
		"User-Agent":      fmt.Sprintf("omniinfer-go-sdk/%s", "v0.1.0"),
		"Accept-Encoding": "gzip, deflate",
	}
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	// send request
	rsp, err := httpCli.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	// validate http response
	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status = %d", rsp.StatusCode)
	}
	var reader io.Reader
	var rs ResponseT
	if rsp.Header.Get("Content-Encoding") == "gzip" {
		grd, err := gzip.NewReader(rsp.Body)
		if err != nil {
			return nil, err
		}
		reader = grd
	} else {
		reader = rsp.Body
	}
	// unmarshal
	if err := json.NewDecoder(reader).Decode(&rs); err != nil {
		return nil, err
	}
	if rs.GetCode() != omni_error.CodeNormal {
		return nil, omni_error.NewOmniError(rs.GetCode(), rs.GetMsg())
	}
	return &rs, nil
}
