package request

import (
	"context"
	"github.com/omniinfer/golang-sdk/types"
	"os"
	"testing"
	"time"
)

func TestOmniClient_Models(t *testing.T) {
	client, err := NewOmniClient(os.Getenv("OMNI_API_KEY"))
	if err != nil {
		t.Error(err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()
	models, err := client.Models(ctx, WithRefresh())
	if err != nil {
		t.Error(err)
		return
	}
	// test filtering and sorting
	t.Log(models)
	top := models.FilterType(types.Checkpoint).TopN(10, func(m *types.Model) float32 {
		return float32(m.CivitaiDownloadCount)
	})
	t.Log(top)
}
