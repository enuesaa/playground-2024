package main

import (
	"context"
	"fmt"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/guregu/dynamo/v2"
)

func main() {
	if err := handle(); err != nil {
		panic(err)
	}
	fmt.Print("created!\n")
}

type Note struct {
	BinderName string
	NoteName string
}
func handle() error {
	ctx := context.Background()
	cfg, err := awsconfig.LoadDefaultConfig(ctx,
		awsconfig.WithRegion("ap-northeast-1"),
	)
	if err != nil {
		return err
	}

	db := dynamo.New(cfg)
	table := db.Table("pinit")

	n := Note{
		BinderName: "test",
		NoteName: "testest",
	}
	return table.Put(n).Run(ctx)
}
