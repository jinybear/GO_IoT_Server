package controller

import (
	"GO_IoT_Server/model"
	"GO_IoT_Server/service"
	"encoding/json"
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

var productService = new(service.ProductService)

type ScootControl struct {
}

// 전체 제품 정보 제공 함수
func (s *ScootControl) GetAllProducts(ctx *fasthttp.RequestCtx) {
	products, err := productService.GetAllProducts()

	if err != nil {
		log.Println(err)
		fmt.Fprintf(ctx, "%s", "products are not exist")
		ctx.SetStatusCode(404)
	} else {
		result, err := json.Marshal(products)
		if err != nil {
			log.Println(err)
		}
		fmt.Fprintf(ctx, string(result))
		ctx.SetStatusCode(200)
	}

	ctx.SetContentType("application/json")
}

// 특정 제품 정보 제공 함수
func (s *ScootControl) GetProduct(ctx *fasthttp.RequestCtx) {
	query_id := string(ctx.QueryArgs().Peek("id"))
	product, err := productService.GetProduct(query_id)

	if err != nil {
		log.Println(err)
		fmt.Fprintf(ctx, "%s", "Failed to find product")
		ctx.SetStatusCode(404)
	} else {
		result, err := json.Marshal(product)
		if err != nil {
			log.Println(err)
		}
		fmt.Fprintf(ctx, string(result))
		ctx.SetStatusCode(200)
	}

	ctx.SetContentType("application/json")
}

// 제품 추가 처리 함수
func (s *ScootControl) AddProduct(ctx *fasthttp.RequestCtx) {
	var product model.Product
	if err := json.Unmarshal(ctx.PostBody(), &product); err != nil {
		log.Println(err)
		ctx.SetStatusCode(400)
		fmt.Fprintf(ctx, "%s", "invalid request")
	} else {
		productService.AddProduct(product)
	}

	ctx.SetContentType("application/json")
}
