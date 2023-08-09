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
	// get all model
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
