package server

import (
	"net/http"
	"time"

	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/trantho123/saleswebsite/models"
	"github.com/trantho123/saleswebsite/utils"
)

type createOrderRequest struct {
	ShippingDetails struct {
		FirstName   string `json:"firstName" binding:"required"`
		LastName    string `json:"lastName" binding:"required"`
		Email       string `json:"email" binding:"required,email"`
		PhoneNumber string `json:"phoneNumber" binding:"required"`
		Address     string `json:"address" binding:"required"`
		City        string `json:"city" binding:"required"`
		State       string `json:"state" binding:"required"`
		ZipCode     string `json:"zipCode" binding:"required"`
	} `json:"shippingDetails" binding:"required"`
	PaymentMethod string `json:"paymentMethod" binding:"required,oneof=COD Card"`
}

func (s *Server) CreateOrder(c *gin.Context) {
	var req createOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Get authenticated user
	email, _, _, exists := authenUser(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	// Get user's cart
	cart, err := s.repo.GetCartByUserId(user.ID.Hex())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart not found"})
		return
	}

	if len(cart.Items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart is empty"})
		return
	}

	// Convert cart items to order items
	var orderItems []models.OrderItem
	for _, cartItem := range cart.Items {
		orderItem := models.OrderItem{
			ProductID: cartItem.ProductID,
			Name:      cartItem.Name,
			Price:     cartItem.Price,
			Quantity:  cartItem.Quantity,
			Image:     cartItem.Image,
		}
		orderItems = append(orderItems, orderItem)
	}

	// Create order from cart
	order := models.Order{
		UserID:      user.ID,
		Items:       orderItems,
		TotalAmount: cart.Totals,
		Status:      "pending",
		ShippingDetails: models.ShippingDetails{
			FirstName:   req.ShippingDetails.FirstName,
			LastName:    req.ShippingDetails.LastName,
			Email:       req.ShippingDetails.Email,
			PhoneNumber: req.ShippingDetails.PhoneNumber,
			Address:     req.ShippingDetails.Address,
			City:        req.ShippingDetails.City,
			State:       req.ShippingDetails.State,
			ZipCode:     req.ShippingDetails.ZipCode,
		},
		PaymentDetails: models.PaymentDetails{
			Method: req.PaymentMethod,
			Status: "pending",
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	orderID, err := s.repo.CreateOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	// Clear cart after creating order
	cart.Items = []models.Item{}
	cart.Totals = 0
	cart.UpdatedAt = time.Now()

	if err := s.repo.UpdateCart(cart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear cart"})
		return
	}
	// Prepare email data
	var emailItems []utils.OrderItem
	for _, item := range order.Items {
		emailItems = append(emailItems, utils.OrderItem{
			Name:     item.Name,
			Quantity: item.Quantity,
			Price:    item.Price,
			Total:    item.Price * item.Quantity,
		})
	}

	emailData := utils.OrderEmailData{
		CustomerName: order.ShippingDetails.FirstName + " " + order.ShippingDetails.LastName,
		OrderID:      orderID,
		Items:        emailItems,
		TotalAmount:  order.TotalAmount,
		ShippingAddress: fmt.Sprintf("%s, %s, %s, %s",
			order.ShippingDetails.Address,
			order.ShippingDetails.City,
			order.ShippingDetails.State,
			order.ShippingDetails.ZipCode,
		),
	}

	config := &utils.EmailConfig{
		SMTPHost:    s.config.SMTPHost,
		SMTPPort:    s.config.SMTPPort,
		SenderEmail: s.config.EmailFrom,
		SenderPass:  s.config.EmailPassword,
	}
	// Send confirmation email
	go func() {
		if err := utils.SendOrderConfirmationEmail(order.ShippingDetails.Email, emailData, config); err != nil {
			log.Printf("Failed to send order confirmation email: %v", err)
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Order created successfully",
		"orderId": orderID,
	})
}

func (s *Server) GetUserOrders(c *gin.Context) {

}

func (s *Server) GetOrder(c *gin.Context) {

}
