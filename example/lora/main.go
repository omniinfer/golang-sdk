package main

import (
	"context"
	"fmt"
	"github.com/omniinfer/golang-sdk/request"
	"github.com/omniinfer/golang-sdk/types"
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
		fmt.Printf("get types list failed, %v\n", err)
		return
	}
	// Anything V5/Ink, https://civitai.com/models/9409/or-anything-v5ink
	modelName := modelList.FilterCivitaiVersionId(90854).SdName
	// Detail Tweaker LoRA, https://civitai.com/models/58390/detail-tweaker-lora-lora
	loraName := modelList.FilterCivitaiVersionId(62833).SdName
	txt2ImgReq := types.NewTxt2ImgRequest(fmt.Sprintf("a dog flying in the sky, <lora:%s:%d>", loraName, 1), "", modelName)
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
