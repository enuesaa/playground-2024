package main

import (
	"fmt"

	"github.com/enuesaa/.devdev/go-fx-httpcall/callerfx"
	"github.com/enuesaa/.devdev/go-fx-httpcall/clifx"
	"go.uber.org/fx"
)

// http リクエストする CLI ツール。curl みたいに。
func main() {
	fxapp := fx.New(
		clifx.Module,
		callerfx.Module,

		// メインロジック (上から順に実行される)
		fx.Invoke(func(cli clifx.ICli) error {
			// cli を立ち上げ
			return cli.Launch()
		}),
		fx.Invoke(func(cli clifx.ICli, caller callerfx.ICaller, shutdowner fx.Shutdowner) error {
			// -url フラグの値を取得
			url := cli.GetUrl()

			// http リクエスト
			body, err := caller.Run(url)
			if err != nil {
				return err
			}
			fmt.Printf("%s", body)

			// fx.App をシャットダウン
			if err := shutdowner.Shutdown(); err != nil {
				return err
			}

			return nil
		}),

		// fx.App のログを出力をしない
		fx.NopLogger,
	)
	fxapp.Run()
}
