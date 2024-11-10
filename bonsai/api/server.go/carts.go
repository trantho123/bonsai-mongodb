package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/trantho123/saleswebsite/models"
)

type cartReq struct {
	ProductID string `uri:"id" binding:"required"`
	Quantity  int32  `uri:"quantity" binding:"required"`
}

func (s *Server) CreateCart(c *gin.Context) {
	var cartReq cartReq
	if err := c.ShouldBindUri(&cartReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	email, role, token, exits := authenUser(c)
	if !exits {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not login"})
		return
	}
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token not found"})
		return
	}
	tk, err := s.repo.GetAccessToken(token)
	if err != nil || tk.AccessToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please login"})
		return
	}
	if !s.repo.IsUserRole(role) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not have permission"})
		return
	}
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	product, err := s.repo.GetProduct(cartReq.ProductID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}
	if !s.repo.IsCartExist(user.ID.Hex()) {
		item := models.Item{
			ProductID: product.ID,
			Price:     product.Price,
			Quantity:  cartReq.Quantity,
			CreatedAt: time.Now(),
		}
		newCart := models.Cart{
			User:      user.ID,
			Product:   []models.Item{item},
			Totals:    cartReq.Quantity * product.Price,
			CreatedAt: time.Now(),
		}
		if err := s.repo.CreateCart(newCart); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		cart, err := s.repo.GetCartByUserId(user.ID.Hex())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if s.repo.IsProductExitsInCart(user.ID.Hex(), product.ID.Hex()) {
			for i, item := range cart.Product {
				if item.ProductID.Hex() == product.ID.Hex() {
					cart.Product[i].Quantity += cartReq.Quantity
					cart.Product[i].UpdatedAt = time.Now()
					cart.Totals += cartReq.Quantity * product.Price
					break
				}
			}
		} else {
			item := models.Item{
				ProductID: product.ID,
				Price:     product.Price,
				Quantity:  cartReq.Quantity,
				CreatedAt: time.Now(),
			}
			cart.Product = append(cart.Product, item)
			cart.Totals += cartReq.Quantity * product.Price
		}
		cart.UpdatedAt = time.Now()
		if err := s.repo.UpdateCart(cart); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cart created successfully"})
}

func (s *Server) GetCart(c *gin.Context) {
	email, role, token, exits := authenUser(c)
	if !exits {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not login"})
		return
	}
	if !s.repo.IsUserRole(role) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not have permission"})
		return
	}
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token not found"})
		return
	}
	tk, err := s.repo.GetAccessToken(token)
	if err != nil || tk.AccessToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please login"})
		return
	}
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	cart, err := s.repo.GetCartByUserId(user.ID.Hex())
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			c.JSON(http.StatusOK, gin.H{"message": "Cart is empty"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cart)
}

type deleteProductInCartReq struct {
	ProductID string `json:"productid" binding:"required"`
}

func (s *Server) DeleteProductInCart(c *gin.Context) {
	var deleteReq deleteProductInCartReq
	if err := c.ShouldBindJSON(&deleteReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	email, role, token, exits := authenUser(c)
	if !exits {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not login"})
		return
	}
	if !s.repo.IsUserRole(role) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not have permission"})
		return
	}
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token not found"})
		return
	}
	tk, err := s.repo.GetAccessToken(token)
	if err != nil || tk.AccessToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please login"})
		return
	}
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	if !s.repo.IsProductExitsInCart(user.ID.Hex(), deleteReq.ProductID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not exits in cart"})
		return
	}
	cart, err := s.repo.GetCartByUserId(user.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for i, item := range cart.Product {
		if item.ProductID.Hex() == deleteReq.ProductID {
			cart.Totals -= item.Price * item.Quantity
			cart.Product = append(cart.Product[:i], cart.Product[i+1:]...)
			break
		}
	}
	cart.UpdatedAt = time.Now()
	if err := s.repo.UpdateCart(cart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

type updateQuantityProductInCartReq struct {
	ProductID string `json:"id" binding:"required"`
	Quantity  int32  `json:"quantity" binding:"required" min:"1"`
}

func (s *Server) UpdateQuantityProductCart(c *gin.Context) {
	var updateReq updateQuantityProductInCartReq
	if err := c.ShouldBindJSON(&updateReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	email, role, token, exits := authenUser(c)
	if !exits {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not login"})
		return
	}
	if !s.repo.IsUserRole(role) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not have permission"})
		return
	}
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token not found"})
		return
	}
	tk, err := s.repo.GetAccessToken(token)
	if err != nil || tk.AccessToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please login"})
		return
	}
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	if !s.repo.IsProductExitsInCart(user.ID.Hex(), updateReq.ProductID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not exits in cart"})
		return
	}

	product, err := s.repo.GetProduct(updateReq.ProductID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}
	if product.Quantity < updateReq.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not enough quantity"})
		return
	}

	cart, err := s.repo.GetCartByUserId(user.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for i, item := range cart.Product {
		if item.ProductID.Hex() == updateReq.ProductID {
			cart.Totals -= item.Price * item.Quantity
			cart.Product[i].Quantity = updateReq.Quantity
			cart.Totals += item.Price * updateReq.Quantity
			break
		}
	}
	cart.UpdatedAt = time.Now()
	if err := s.repo.UpdateCart(cart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	cartAfterUpdate, err := s.repo.GetCartByUserId(user.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cartAfterUpdate)
}
