package login

import (
	"context"
	// "encoding/json"
	"fmt"
	httputil "learn-project/httputil"

	"github.com/kataras/iris/v12"
)

type Res struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func Login(ctx iris.Context) {
	// fmt.Printf("login")
	// data:={res:"1999"}
	res := Res{200, "success"}
	fmt.Println("Received:", res)
	httputil.BadRequest(ctx, "err")
	httputil.UnAuthenticate(ctx, "unauthen")
	// ctx.JSONP(res)
}
func Logout(ctx context.Context) {
	fmt.Println("logout")
}
