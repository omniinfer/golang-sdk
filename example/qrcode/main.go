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
