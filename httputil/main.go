package httputil

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

type Res struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Success(ctx iris.Context) {

}
func BadRequest(ctx iris.Context, mes string) {
	res := Res{Status: 400, Message: mes}
	ctx.JSON(res)
	// fmt.Fprintf("u")
}
func UnAuthenticate(ctx iris.Context, mes string) {
	res := Res{Status: 401, Message: mes}
	ctx.JSON(res)
	fmt.Sprintf("h")
	// return false
}
func ServerErr(ctx iris.Context, mes string) {
	res := Res{Status: 500, Message: mes}
	ctx.JSON(res)
	fmt.Sprintf("k")
}
