package router

import (
	"GO_IoT_Server/controller"

	"github.com/valyala/fasthttp"
)

var scootController = new(controller.ScootControl)

func RequestHandler(ctx *fasthttp.RequestCtx) {
	switch {
	case string(ctx.Path()) == "/product" && string(ctx.Method()) == "GET":
		scootController.GetProduct(ctx)
	case string(ctx.Path()) == "/product" && string(ctx.Method()) == "POST":
		scootController.AddProduct(ctx)
	case string(ctx.Path()) == "/products" && string(ctx.Method()) == "GET":
		scootController.GetAllProducts(ctx)
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
}
