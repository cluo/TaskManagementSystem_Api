package blls

import (
	"TaskManagementSystem_Api/models/dals"
	"TaskManagementSystem_Api/models/types"
)

type ProductBLL struct {
}

func (bll *ProductBLL) GetProducts(pageSize, pageNumber int) (t []*types.ProductHeader_Get, err error) {
	t, err = (&dals.ProductDAL{}).GetProductHeaders(pageSize, pageNumber)
	return
}
func (bll *ProductBLL) GetProductCount() (counts map[string]int, err error) {
	counts, err = (&dals.ProductDAL{}).GetProductCount()
	return
}
func (bll *ProductBLL) GetProductDetail(id string) (t *types.Product_Get, err error) {
	t, err = (&dals.ProductDAL{}).GetProductDetail(id)
	return
}
func (bll *ProductBLL) AddProduct(productPost types.Product_Post, user types.UserInfo_Get) (err error) {
	err = (&dals.ProductDAL{}).AddProduct(productPost, user)
	return
}

func (bll *ProductBLL) DeleteProduct(id string, user types.UserInfo_Get) (err error) {
	err = (&dals.ProductDAL{}).DeleteProduct(id, user)
	return
}
func (bll *ProductBLL) UpdateProduct(id string, product types.Product_Post, user types.UserInfo_Get) (err error) {
	err = (&dals.ProductDAL{}).UpdateProduct(id, product, user)
	return
}
