package service

import (
	"GO_IoT_Server/model"
	"errors"
)

type ProductService struct{}

func (p *ProductService) GetAllProducts() ([]model.Product, error) {
	productList := []model.Product{}

	return productList, errors.New("Not implement")
}

func (p *ProductService) GetProduct(id string) (model.Product, error) {
	product := model.Product{}

	return product, errors.New("Not implement")
}

func (p *ProductService) AddProduct(product model.Product) error {
	// Database 에 Insert 과정 실행

	return errors.New("Not implement")
}
