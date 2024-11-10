package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type productIdReq struct {
	ID string `uri:"id" binding:"required"`
}

type productResp struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       int32    `json:"price"`
	Quantity    int32    `json:"quantity"`
	Image       []string `json:"image"`
}

func (s *Server) GetProduct(ctx *gin.Context) {
	var req productIdReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	product, err := s.repo.GetProduct(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	productResp := productResp{
		ID:          product.ID.Hex(),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
		Image:       product.Image,
	}
	ctx.JSON(http.StatusOK, productResp)
}

type productListReq struct {
	Page     int `uri:"page" `
	PageSize int `uri:"page_size" binding:"max=10"`
}

type productListResp struct {
	PageId   int           `json:"page_id"`
	PageSize int           `json:"page_size"`
	Total    int64         `json:"total"`
	Products []productResp `json:"products"`
}

func (s *Server) GetListProducts(ctx *gin.Context) {
	var req productListReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	products, total, err := s.repo.GetListProducts(req.Page, req.PageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	listproduct := []productResp{}
	for _, product := range products {
		productResp := productResp{
			ID:          product.ID.Hex(),
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Quantity:    product.Quantity,
			Image:       product.Image,
		}
		listproduct = append(listproduct, productResp) // Thêm phần tử vào slice
	}
	ctx.JSON(http.StatusOK, productListResp{
		PageId:   req.Page,
		PageSize: req.PageSize,
		Total:    total,
		Products: listproduct,
	})
}
