package types

import (
	"context"
	"io"
	"net/http"
	"sync"
)

type ProgressData struct {
	Status        ProgressResponseStatusCode `json:"status"`
	Progress      float32                    `json:"progress"`
	ETARelative   float64                    `json:"eta_relative"`
	Imgs          []string                   `json:"imgs"`
	FailedReason  string                     `json:"failed_reason"`
	CurrentImages []string                   `json:"current_images"`
	SubmitTime    string                     `json:"submit_time"`
	ExecutionTime string                     `json:"execution_time"`
	Txt2ImgTime   string                     `json:"txt2img_time"`
	FinishTime    string                     `json:"finish_time"`
	ImgsBytes     [][]byte
}

type ProgressResponseStatusCode int

func (p ProgressResponseStatusCode) IsFinish() bool {
	return p == Successful || p == Failed || p == Timeout
}

const (
	Initializing ProgressResponseStatusCode = 0
	Running                                 = 1
	Successful                              = 2
	Failed                                  = 3
	Timeout                                 = 4
)

type ProgressRequest struct {
	TaskId string `json:"task_id"`
}

type ProgressResponse struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data *ProgressData `json:"data"`
}

func (p ProgressResponse) GetCode() int {
	return p.Code
}

func (p ProgressResponse) GetMsg() string {
	return p.Msg
}

func (p *ProgressData) DownloadImages(ctx context.Context) error {
	errs := make([]error, len(p.Imgs), len(p.Imgs))
	rs := make([][]byte, len(p.Imgs), len(p.Imgs))
	wg := sync.WaitGroup{}
	for idx, img := range p.Imgs {
		wg.Add(1)
		go func(idx int, u string) {
			defer wg.Done()
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
			if err != nil {
				errs[idx] = err
				return
			}
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				errs[idx] = err
				return
			}
			defer res.Body.Close()
			bs, err := io.ReadAll(res.Body)
			if err != nil {
				errs[idx] = err
				return
			}
			rs[idx] = bs
		}(idx, img)
	}
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	p.ImgsBytes = rs
	return nil
}
