package request

import (
	"context"
	"github.com/omniinfer/golang-sdk/types"
	"github.com/omniinfer/golang-sdk/util"
	"os"
	"testing"
	"time"
)

func TestOmniClient_SyncUpscale(t *testing.T) {
	client, err := NewOmniClient(os.Getenv("OMNI_API_KEY"))
	if err != nil {
		t.Error(err)
		return
	}
	initImage := "out/test_txt2img_sync.png"
	initImageBase64, err := util.ReadImageToBase64(initImage)
	if err != nil {
		t.Error(err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()
	upscale := types.NewUpscaleRequest(initImageBase64, 2)
	res, err := client.SyncUpscale(ctx, upscale,
		WithSaveImage("out", 0777, func(taskId string, fileIndex int, fileName string) string {
			return "test_upscale_sync.png"
		}))
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("status = %d", res.Data.Status)
}
