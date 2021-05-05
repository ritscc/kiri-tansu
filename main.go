package main

import (
	"net/http"
	"ritscc/kiri-tansu/controller"
	migration "ritscc/kiri-tansu/db/migration"
)

func main() {
	// マイグレーション
	if err := migration.Migrate(); err != nil {
		return
	}

	// ルーティング
	controller.Routing()

	// サーバ起動
	http.ListenAndServe(":8080", nil)
}
