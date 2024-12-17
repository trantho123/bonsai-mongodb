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
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Dob       string `json:"dob" binding:"required"`
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
	newUser := models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
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
	accesstoken, err := jwt.CreateToken(newUser.Email, roleIdUser, s.config.AccessTokenKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "User created successfully",
		"success":   true,
		"authToken": accesstoken,
	})
}

type userLoginParams struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s *Server) Login(c *gin.Context) {
	var user userLoginParams
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userData, err := s.repo.GetUserByEmail(user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email/Password is incorrect"})
		return
	}
	if !utils.CheckPasswordHash(user.Password, userData.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email/Password is incorrect"})
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

	c.JSON(http.StatusOK, gin.H{
		"message":   "Login successfully",
		"success":   true,
		"authToken": accesstoken,
	})
}

type userGetProfileResponse struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Dob       string `json:"dob"`
	City      string `json:"city"`
	Postal    string `json:"zipCode"`
	State     string `json:"userState"`
}

func (s *Server) GetProfile(c *gin.Context) {
	email, _, _, exits := authenUser(c)
	if !exits {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not login"})
		return
	}
	userData, err := s.repo.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	userRes := userGetProfileResponse{
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Email:     userData.Email,
		Phone:     userData.Phone,
		Address:   userData.Address,
		Dob:       userData.Dob,
		City:      userData.City,
		Postal:    userData.Postal,
		State:     userData.State,
	}
	c.JSON(http.StatusOK, gin.H{"data": userRes})
}

type userUpdateProfileParam struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Dob       string `json:"dob"`
	City      string `json:"city"`
	Postal    string `json:"zipCode"`
	State     string `json:"userState"`
}

func (s *Server) UpdateProfile(c *gin.Context) {
	email, _, _, exits := authenUser(c)
	if !exits {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not login"})
		return
	}
	var userReq userUpdateProfileParam
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userData, err := s.repo.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	userUpdate := models.User{}
	if userReq.FirstName != userData.FirstName {
		userUpdate.FirstName = userReq.FirstName
	} else {
		userUpdate.FirstName = userData.FirstName
	}
	if userReq.LastName != userData.LastName {
		userUpdate.LastName = userReq.LastName
	} else {
		userUpdate.LastName = userData.LastName
	}
	if userReq.Phone != userData.Phone {
		userUpdate.Phone = userReq.Phone
	} else {
		userUpdate.Phone = userData.Phone
	}
	if userReq.Address != userData.Address {
		userUpdate.Address = userReq.Address
	} else {
		userUpdate.Address = userData.Address
	}
	if userReq.Dob != userData.Dob {
		userUpdate.Dob = userReq.Dob
	} else {
		userUpdate.Dob = userData.Dob
	}
	if userReq.City != userData.City {
		userUpdate.City = userReq.City
	} else {
		userUpdate.City = userData.City
	}
	if userReq.Postal != userData.Postal {
		userUpdate.Postal = userReq.Postal
	} else {
		userUpdate.Postal = userData.Postal
	}
	if userReq.State != userData.State {
		userUpdate.State = userReq.State
	} else {
		userUpdate.State = userData.State
	}
	userUpdate.Email = userData.Email
	userUpdate.Password = userData.Password
	userUpdate.UpdatedAt = time.Now()
	userUpdate.Role = userData.Role
	err = s.repo.UpdateUser(userUpdate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
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

	c.JSON(http.StatusOK, gin.H{"message": "Logout successfully"})
}

type resetPasswordParam struct {
	CurrentPassword string `json:"currentPassword" binding:"required"`
	NewPassword     string `json:"newPassword" binding:"required"`
}

func (s *Server) ResetPassword(c *gin.Context) {
	email, _, _, exits := authenUser(c)
	if !exits {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not login"})
		return
	}
	var userReq resetPasswordParam
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !utils.CheckPasswordHash(userReq.CurrentPassword, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Current password is incorrect"})
		return
	}
	err = s.repo.ResetPassword(email, utils.HashingPassword(userReq.NewPassword))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (s *Server) DeleteUser(c *gin.Context) {
	email, _, _, exits := authenUser(c)
	if !exits {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not login"})
		return
	}
	err := s.repo.DeleteUser(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (s *Server) VerifyEmail(c *gin.Context) {
	code := c.Param("code")

	// Tìm user với verification code
	user, err := s.repo.GetUserByVerificationCode(code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verification code"})
		return
	}
	err = s.repo.UpdateUserVerification(user.ID.Hex(), true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Email verified successfully",
	})
}

type forgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func (s *Server) ForgotPassword(c *gin.Context) {
	var req forgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Kiểm tra email có tồn tại không
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not found"})
		return
	}

	// Tạo reset token
	resetToken, err := jwt.CreateResetToken(user.Email, s.config.AccessTokenKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reset token"})
		return
	}

	// Tạo reset link
	resetLink := fmt.Sprintf("http://localhost:3000/reset-password/%s", resetToken)

	// Chuẩn bị data cho email
	emailData := utils.ResetPasswordEmailData{
		Username:  user.FirstName + " " + user.LastName,
		ResetLink: resetLink,
	}

	// Cấu hình email
	emailConfig := &utils.EmailConfig{
		SMTPHost:    s.config.SMTPHost,
		SMTPPort:    s.config.SMTPPort,
		SenderEmail: s.config.EmailFrom,
		SenderPass:  s.config.EmailPassword,
	}

	// Gửi email
	if err := utils.SendResetPasswordEmail(user.Email, emailData, emailConfig); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send reset password email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Reset password link has been sent to your email",
	})
}

type resetPasswordRequest struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

func (s *Server) ResetPasswordWithToken(c *gin.Context) {
	var req resetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate token
	claims, err := jwt.ValidateToken(req.Token, s.config.AccessTokenKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired token"})
		return
	}

	// Validate password
	if err := utils.IsPasswordValid(req.NewPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash new password
	hashedPassword := utils.HashingPassword(req.NewPassword)

	// Update password in database
	if err := s.repo.ResetPassword(claims.Email, hashedPassword); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reset password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Password has been reset successfully",
	})
}
