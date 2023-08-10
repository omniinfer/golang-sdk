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
	// get all models
	modelList, err := client.Models(ctx)
	if err != nil {
		fmt.Printf("get model list failed, %v\n", err)
		return
	}
	// top 10 checkpoint
	modelList = modelList.FilterType(types.Checkpoint).TopN(10, func(m *types.Model) float32 {
		return m.CivitaiRating
	})
	for _, m := range modelList {
		fmt.Println(m.Name)
	}
}
