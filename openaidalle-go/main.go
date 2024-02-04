package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/sashabaranov/go-openai"
)

var (
	token = flag.String("token", "", "openai token")
	prompt = flag.String("prompt", "", "prompt")
)

func main() {
	flag.Parse()
	if len(*token) == 0 {
		log.Fatalf("please give me -token")
	}
	if len(*prompt) == 0 {
		log.Fatalf("please give me -prompt")
	}
	fmt.Println("start")

	client := openai.NewClient(*token)

	ctx := context.Background()
	res, err := client.CreateImage(ctx, openai.ImageRequest{
		Prompt:         *prompt,
		Model:          openai.CreateImageModelDallE3,
		Size:           openai.CreateImageSize1024x1024,
		ResponseFormat: openai.CreateImageResponseFormatURL,
		N:              1,
	})
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	fmt.Printf("%v\n", res)
	fmt.Printf("url: %s\n", res.Data[0].URL)
}
