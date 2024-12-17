package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trantho123/saleswebsite/models"
)

type productIdReq struct {
	ID string `uri:"id" binding:"required"`
}

type productResp struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Price       int32             `json:"price"`
	Quantity    int32             `json:"quantity"`
	Image       string            `json:"image"`
	Rating      float32           `json:"rating"`
	Tags        []models.ListTags `json:"tags"`
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
		Rating:      product.Rating,
		Tags:        product.Tags,
	}
	ctx.JSON(http.StatusOK, productResp)
}

func (s *Server) GetListProducts(ctx *gin.Context) {
	products, err := s.repo.GetListProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var productResps []productResp
	for _, product := range products {
		productResp := productResp{
			ID:          product.ID.Hex(),
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Quantity:    product.Quantity,
			Image:       product.Image,
			Rating:      product.Rating,
			Tags:        product.Tags,
		}
		productResps = append(productResps, productResp)
	}

	ctx.JSON(http.StatusOK, productResps)
}

type productTagsRequest struct {
	Tags []string `json:"tags" binding:"required"`
}

func (s *Server) GetProductByTags(ctx *gin.Context) {
	var req productTagsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Lấy sản phẩm theo nhiều tags
	products, err := s.repo.GetProductsByTags(req.Tags)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var productResps []productResp
	for _, product := range products {
		productResp := productResp{
			ID:          product.ID.Hex(),
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Quantity:    product.Quantity,
			Image:       product.Image,
			Rating:      product.Rating,
			Tags:        product.Tags,
		}
		productResps = append(productResps, productResp)
	}

	ctx.JSON(http.StatusOK, productResps)
}
