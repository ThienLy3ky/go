package main

import (
	login "learn-project/api/login"

	"github.com/kataras/iris/v12"
)

type PingResponse struct {
	Message string `json:"message"`
}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
func main() {
	app := iris.New()
	app.Use(myMiddleware)

	app.Get("/ping", login.Login)
	app.Listen(":8080")
}
