package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/trantho123/saleswebsite/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserResponse struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	Postal    string    `json:"postal"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
}

func (s *Server) GetAllUsers(c *gin.Context) {
	// Get email and role ID from auth middleware
	email, roleId, token, exists := authenUser(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Authentication failed",
			"details": "No valid authentication token found",
			"code":    "AUTH_REQUIRED",
		})
		return
	}

	// Log authentication details for debugging
	fmt.Printf("Auth check - Email: %s, RoleID: %s, Token exists: %v\n", email, roleId, token != "")

	// Check if user has admin role
	isAdmin := s.repo.IsAdminRole(roleId)
	if !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   "Access denied",
			"details": fmt.Sprintf("User %s does not have admin privileges. Role ID: %s", email, roleId),
			"code":    "ADMIN_REQUIRED",
		})
		return
	}

	// Get all users from database
	users, err := s.repo.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch users",
			"details": err.Error(),
			"code":    "DATABASE_ERROR",
		})
		return
	}

	// Convert users to response format
	var usersResponse []UserResponse
	for _, user := range users {
		// Get role name
		role, err := s.repo.GetRole(user.Role.Hex())
		if err != nil {
			fmt.Printf("Failed to get role for user %s: %v\n", user.Email, err)
			continue
		}

		userResponse := UserResponse{
			ID:        user.ID.Hex(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
			Address:   user.Address,
			City:      user.City,
			State:     user.State,
			Postal:    user.Postal,
			Role:      role.Name,
			CreatedAt: user.CreatedAt,
		}
		usersResponse = append(usersResponse, userResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    usersResponse,
		"count":   len(usersResponse),
	})
}

type PaymentData struct {
	ID        string    `json:"id"`
	User      string    `json:"user"`
	Amount    int32     `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

type ChartData struct {
	Orders   []PaymentData `json:"orders"`
	Products []productResp `json:"products"`
}

func (s *Server) GetChartData(c *gin.Context) {
	// Get email and role ID from auth middleware

	orders, err := s.repo.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to fetch orders: " + err.Error(),
		})
		return
	}

	var payments []PaymentData
	for _, order := range orders {
		user, err := s.repo.GetUserByID(order.UserID.Hex())
		if err != nil {
			fmt.Printf("Failed to get user for order %s: %v\n", order.ID.Hex(), err)
			continue
		}
		payment := PaymentData{
			ID:        order.ID.Hex(),
			User:      user.FirstName + " " + user.LastName,
			Amount:    order.TotalAmount,
			Status:    order.Status,
			CreatedAt: order.CreatedAt,
		}
		payments = append(payments, payment)
	}

	products, err := s.repo.GetListProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to fetch products: " + err.Error(),
		})
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
	rep := ChartData{
		Orders:   payments,
		Products: productResps,
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    rep,
	})

}

type CreateUserAdminRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
}

func (s *Server) CreateUserAdmin(c *gin.Context) {
	email, roleId, token, exists := authenUser(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Authentication failed",
			"details": "No valid authentication token found",
			"code":    "AUTH_REQUIRED",
		})
		return
	}

	// Log authentication details for debugging
	fmt.Printf("Auth check - Email: %s, RoleID: %s, Token exists: %v\n", email, roleId, token != "")

	// Check if user has admin role
	isAdmin := s.repo.IsAdminRole(roleId)
	if !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   "Access denied",
			"details": fmt.Sprintf("User %s does not have admin privileges. Role ID: %s", email, roleId),
			"code":    "ADMIN_REQUIRED",
		})
		return
	}
	var req CreateUserAdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"details": err.Error(),
		})
		return
	}
	if s.repo.IsEmailExist(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}
	roleAdminId, err := primitive.ObjectIDFromHex(roleId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Phone:     req.Phone,
		Role:      roleAdminId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Verified:  true,
	}

	err = s.repo.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
