package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/trantho123/saleswebsite/jwt"
	"github.com/trantho123/saleswebsite/models"
	"github.com/trantho123/saleswebsite/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userRegisterParams struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Dob      string `json:"dob" binding:"required"`
}

// CreateUser creates a new user
func (s *Server) CreateUser(c *gin.Context) {
	var user userRegisterParams
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := utils.IsEmailValid(user.Email); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := utils.IsPasswordValid(user.Password); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := utils.IsValidDOB(user.Dob); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if s.repo.IsEmailExist(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}
	if s.repo.IsUserExist(user.Username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}
	roleIdUser := s.AddRoles("User")
	if roleIdUser == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Role not found"})
		return
	}
	roleId, err := primitive.ObjectIDFromHex(roleIdUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	newUser := models.CreateUser{
		Username:  user.Username,
		Email:     user.Email,
		Password:  utils.HashingPassword(user.Password),
		Dob:       user.Dob,
		CreatedAt: time.Now(),
		Role:      roleId,
	}
	if err := s.repo.CreateUser(newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

type userLoginParams struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s *Server) Login(c *gin.Context) {
	var user userLoginParams
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userData, err := s.repo.GetUserByUsername(user.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	if !utils.CheckPasswordHash(user.Password, userData.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is incorrect"})
		return
	}
	rolesUserId := s.AddRoles("User")
	if rolesUserId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Role not found"})
		return
	}
	accesstoken, err := jwt.CreateToken(userData.Email, rolesUserId, s.config.AccessTokenKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = s.repo.CreateAccessToken(models.AccessToken{AccessToken: accesstoken})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     "Login successfully",
		"accesstoken": accesstoken})
}

type userResponse struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Dob       string `json:"dob"`
	CreatedAt time.Time
}

func (s *Server) GetProfile(c *gin.Context) {
	email, role, token, exits := authenUser(c)
	if !exits {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not login"})
		return
	}
	if role != s.AddRoles("User") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token not found"})
		return
	}
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	userRes := userResponse{
		Username:  user.Username,
		Email:     user.Email,
		Dob:       user.Dob,
		CreatedAt: user.CreatedAt,
	}
	c.JSON(http.StatusOK, userRes)
}

func (s *Server) AddRoles(roleString string) string {
	roles, err := s.repo.GetAllRoles()
	if err != nil {
		return ""
	}
	for _, role := range roles {
		if role.Name == roleString {
			return role.ID.Hex()
		}
	}
	return ""
}

func authenUser(c *gin.Context) (string, string, string, bool) {
	email, exists := c.Get("email")
	if !exists {
		return "", "", "", false
	}
	role, exists := c.Get("role")
	if !exists {
		return "", "", "", false
	}
	token, exists := c.Get("token")
	if !exists {
		return "", "", "", false
	}
	fmt.Println(email, role, token)
	return email.(string), role.(string), token.(string), true
}

func (s *Server) Logout(c *gin.Context) {
	email, role, token, exits := authenUser(c)
	if !exits {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not login"})
		return
	}
	if role != s.AddRoles("User") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not found"})
		return
	}
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token not found"})
		return
	}
	err := s.repo.DeleteAccessToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Logout successfully"})
}
