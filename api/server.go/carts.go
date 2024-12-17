package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/trantho123/saleswebsite/models"
)

type cartRequest struct {
	ProductID string `json:"productId" binding:"required"`
	Quantity  int32  `json:"quantity" binding:"required"`
}

func (s *Server) CreateCart(c *gin.Context) {
	var req cartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
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

	if !s.repo.IsUserRole(role) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not have permission"})
		return
	}
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	product, err := s.repo.GetProduct(req.ProductID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}
	if product.Quantity < req.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough product quantity"})
		return
	}
	if !s.repo.IsCartExist(user.ID.Hex()) {
		item := models.Item{
			ProductID: product.ID,
			Name:      product.Name,
			Price:     product.Price,
			Quantity:  req.Quantity,
			Image:     product.Image,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		newCart := models.Cart{
			User:      user.ID,
			Items:     []models.Item{item},
			Totals:    req.Quantity * product.Price,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
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
		if s.repo.IsItemInCart(user.ID.Hex(), product.ID.Hex()) {
			for i, item := range cart.Items {
				if item.ProductID.Hex() == product.ID.Hex() {
					cart.Items[i].Quantity = req.Quantity
					cart.Items[i].UpdatedAt = time.Now()
					cart.Totals = calculateTotal(cart.Items)
					break
				}
			}
		} else {
			item := models.Item{
				ProductID: product.ID,
				Name:      product.Name,
				Price:     product.Price,
				Quantity:  req.Quantity,
				Image:     product.Image,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			cart.Items = append(cart.Items, item)
			cart.Totals = calculateTotal(cart.Items)
		}
		cart.UpdatedAt = time.Now()
		if err := s.repo.UpdateCart(cart); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cart updated successfully"})
}

func calculateTotal(items []models.Item) int32 {
	var total int32
	for _, item := range items {
		total += item.Price * item.Quantity
	}
	return total
}

func (s *Server) GetCart(c *gin.Context) {
	email, role, _, exits := authenUser(c)
	if !exits {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not login"})
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
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Authentication checks
	email, role, _, exits := authenUser(c)
	if !exits {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not login"})
		return
	}
	if !s.repo.IsUserRole(role) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not have permission"})
		return
	}

	// Get user
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// Check if item exists in cart
	if !s.repo.IsItemInCart(user.ID.Hex(), deleteReq.ProductID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found in cart"})
		return
	}

	// Get cart
	cart, err := s.repo.GetCartByUserId(user.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Remove item from cart
	for i, item := range cart.Items {
		if item.ProductID.Hex() == deleteReq.ProductID {
			cart.Totals -= item.Price * item.Quantity
			cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
			break
		}
	}

	// Update cart
	cart.UpdatedAt = time.Now()
	if err := s.repo.UpdateCart(cart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product removed from cart"})
}

type updateQuantityProductInCartReq struct {
	ProductID string `json:"productId" binding:"required"`
	Quantity  int32  `json:"newQuantity" binding:"required" min:"1"`
}

func (s *Server) UpdateQuantityProductCart(c *gin.Context) {
	var updateReq updateQuantityProductInCartReq
	if err := c.ShouldBindJSON(&updateReq); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Authentication checks
	email, role, _, exits := authenUser(c)
	if !exits {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not login"})
		return
	}
	if !s.repo.IsUserRole(role) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not have permission"})
		return
	}

	// Get user
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// Check if item exists in cart
	if !s.repo.IsItemInCart(user.ID.Hex(), updateReq.ProductID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found in cart"})
		return
	}

	// Check product quantity
	product, err := s.repo.GetProduct(updateReq.ProductID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}
	if product.Quantity < updateReq.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough product quantity"})
		return
	}

	// Get cart
	cart, err := s.repo.GetCartByUserId(user.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update item quantity
	for i, item := range cart.Items {
		if item.ProductID.Hex() == updateReq.ProductID {
			cart.Totals -= item.Price * item.Quantity
			cart.Items[i].Quantity = updateReq.Quantity
			cart.Items[i].UpdatedAt = time.Now()
			cart.Totals += item.Price * updateReq.Quantity
			break
		}
	}

	// Update cart
	cart.UpdatedAt = time.Now()
	if err := s.repo.UpdateCart(cart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get updated cart
	cartAfterUpdate, err := s.repo.GetCartByUserId(user.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cartAfterUpdate)
}
