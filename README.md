# Omniinfer Golang SDK

This SDK is based on the official [API documentation](https://docs.omniinfer.io/).

**Join our discord server for help**

[![](https://dcbadge.vercel.app/api/server/nzqq8UScpx)](https://discord.gg/nzqq8UScpx)

## Installation

```bash
go get -u github.com/omniinfer/golang-sdk
```

## Quick Start

**Get api key refer to [https://docs.omniinfer.io/get-started](https://docs.omniinfer.io/get-started/)**

```golang
package main

import (
	"context"
	"fmt"
	"github.com/omniinfer/golang-sdk/model"
	"github.com/omniinfer/golang-sdk/request"
	"time"
)

func main() {
	// get your api key refer to https://docs.omniinfer.io/get-started/
	const apiKey = "your-key"
	client, err := request.NewOmniClient(apiKey)
	if err != nil {
		fmt.Printf("new omniclient failed, %v\n", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()
	txt2ImgReq := model.NewTxt2ImgRequest("a dog flying in the sky", "", "AnythingV5_v5PrtRE.safetensors")
	res, err := client.SyncTxt2img(ctx, txt2ImgReq,
		request.WithSaveImage("out", 0777, func(taskId string, fileIndex int, fileName string) string {
			return "test_txt2img_sync.png"
		}))
	if err != nil {
		fmt.Printf("generate image failed, %v\n", err)
		return
	}
	for _, s3Url := range res.Data.Imgs {
		fmt.Printf("generate image url: %v\n", s3Url)
	}
}
```

## Examples

### Txt2Img with LoRA

[example_txt2img_with_lora](example/lora/main.go)

```golang
package main

import (
	"context"
	"fmt"
	"github.com/omniinfer/golang-sdk/types"
	"github.com/omniinfer/golang-sdk/request"
	"time"
)

func main() {
	// get your api key refer to https://docs.omniinfer.io/get-started/
	const apiKey = "your-key"
	client, err := request.NewOmniClient(apiKey)
	if err != nil {
		fmt.Printf("new omniclient failed, %v\n", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()
	modelList, err := client.Models(ctx)
	if err != nil {
		fmt.Printf("get model list failed, %v\n", err)
		return
	}
	// Anything V5/Ink, https://civitai.com/models/9409/or-anything-v5ink
	modelName := modelList.FilterCivitaiVersionId(90854).SdName
	// Detail Tweaker LoRA, https://civitai.com/models/58390/detail-tweaker-lora-lora
	loraName := modelList.FilterCivitaiVersionId(62833).SdName
	txt2ImgReq := model.NewTxt2ImgRequest(fmt.Sprintf("a dog flying in the sky, <lora:%s:%d>", loraName, 1), "", modelName)
	res, err := client.SyncTxt2img(ctx, txt2ImgReq,
		request.WithSaveImage("out", 0777, func(taskId string, fileIndex int, fileName string) string {
			return "test_txt2img_sync.png"
		}))
	if err != nil {
		fmt.Printf("generate image failed, %v\n", err)
		return
	}
	for _, s3Url := range res.Data.Imgs {
		fmt.Printf("generate image url: %v\n", s3Url)
	}
}
```

### Model Search

```golang
package main

import (
	"context"
	"fmt"
	"github.com/omniinfer/golang-sdk/model"
	"github.com/omniinfer/golang-sdk/request"
	"time"
)

func main() {
	// get your api key refer to https://docs.omniinfer.io/get-started/
	const apiKey = "your-key"
	client, err := request.NewOmniClient(apiKey)
	if err != nil {
		fmt.Printf("new omniclient failed, %v\n", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()
	// get all models
	modelList, err := client.Models(ctx)
	if err != nil {
		fmt.Printf("get model list failed, %v\n", err)
		return
	}
	// top 10 checkpoint
	modelList = modelList.FilterType(model.Checkpoint).TopN(10, func(m *model.Model) float32 {
		return m.CivitaiRating
	})
	for _, m := range modelList {
		fmt.Println(m.Name)
	}
}
```

### ControlNet QRCode

```golang
package main

import (
	"context"
	"fmt"
	"github.com/omniinfer/golang-sdk/model"
	"github.com/omniinfer/golang-sdk/request"
	"github.com/omniinfer/golang-sdk/util"
	"time"
)

func main() {
	// get your api key refer to https://docs.omniinfer.io/get-started/
	const apiKey = "your-key"
	client, err := request.NewOmniClient(apiKey)
	if err != nil {
		fmt.Printf("new omniclient failed, %v\n", err)
		return
	}
	initImage := "./example/qrcode/qrcode.png"
	initImageBase64, err := util.ReadImageToBase64(initImage)
	if err != nil {
		fmt.Printf("read image failed, %v\n", err)
		return
	}
	txt2ImgReq := &model.Txt2ImgRequest{
		Prompt:      "a beautify butterfly in the colorful flowers, best quality, best details, masterpiece",
		ModelName:   "AnythingV5_v5PrtRE.safetensors",
		SamplerName: model.DPMPPMKarras,
		BatchSize:   1,
		NIter:       1,
		Steps:       30,
		CfgScale:    7,
		Height:      512,
		Width:       512,
		Seed:        -1,
		ControlNetUnits: []*model.ControlNetUnit{
			{
				Model:         "control_v1p_sd15_qrcode_monster_v2",
				Weight:        2.0,
				Module:        model.None,
				InputImage:    initImageBase64,
				ControlMode:   model.Balanced,
				ResizeMode:    model.JustResize,
				GuidanceStart: 0,
				GuidanceEnd:   1,
			},
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()
	res, err := client.SyncTxt2img(ctx, txt2ImgReq)
	if err != nil {
		fmt.Printf("read image failed, %v\n", err)
		return
	}
	for _, s3Url := range res.Data.Imgs {
		fmt.Printf("generate image url: %v\n", s3Url)
	}
}
```

## Testing

```
OMNI_API_KEY=<your-key> go test ./...
```