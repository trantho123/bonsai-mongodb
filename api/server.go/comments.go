package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trantho123/saleswebsite/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type createCommentRequest struct {
	ProductID string  `json:"productId" binding:"required"`
	Content   string  `json:"content" binding:"required"`
	Rating    float32 `json:"rating" binding:"required,min=1,max=5"`
}

type commentResponse struct {
	ID        primitive.ObjectID `json:"_id"`
	UserID    primitive.ObjectID `json:"userId"`
	ProductID primitive.ObjectID `json:"productId"`
	Content   string             `json:"content"`
	Rating    float32            `json:"rating"`
	CreatedAt string             `json:"createdAt"`
	UpdatedAt string             `json:"updatedAt"`
	User      struct {
		ID        primitive.ObjectID `json:"_id"`
		Email     string             `json:"email"`
		FirstName string             `json:"firstName"`
		LastName  string             `json:"lastName"`
	} `json:"user"`
}

type updateCommentRequest struct {
	ID      string  `json:"id" binding:"required"`
	Content string  `json:"comment" binding:"required"`
	Rating  float32 `json:"rating" binding:"required,min=1,max=5"`
}

func (s *Server) CreateComment(c *gin.Context) {
	var req createCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request format: " + err.Error(),
		})
		return
	}

	// Get authenticated user
	email, _, _, exists := authenUser(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User not authenticated",
		})
		return
	}

	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to get user details",
		})
		return
	}

	productID, err := primitive.ObjectIDFromHex(req.ProductID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid product ID format",
		})
		return
	}

	// Verify product exists
	if !s.repo.IsProductExist(req.ProductID) {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Product not found",
		})
		return
	}

	comment := models.Comment{
		UserID:    user.ID,
		ProductID: productID,
		Content:   req.Content,
		Rating:    req.Rating,
	}

	if err := s.repo.CreateComment(comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to create comment: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Comment created successfully",
		"data": commentResponse{
			ID:        comment.ID,
			UserID:    comment.UserID,
			ProductID: comment.ProductID,
			Content:   comment.Content,
			Rating:    comment.Rating,
			CreatedAt: comment.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt: comment.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
			User: struct {
				ID        primitive.ObjectID `json:"_id"`
				Email     string             `json:"email"`
				FirstName string             `json:"firstName"`
				LastName  string             `json:"lastName"`
			}{
				ID:        user.ID,
				Email:     user.Email,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			},
		},
	})
}

func (s *Server) GetProductComments(c *gin.Context) {
	productID := c.Param("id")

	// Validate product exists
	if !s.repo.IsProductExist(productID) {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Product not found",
		})
		return
	}

	comments, err := s.repo.GetCommentsByProductID(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to get comments: " + err.Error(),
		})
		return
	}

	// Enhance comments with user details
	var enhancedComments []commentResponse
	for _, comment := range comments {
		user, err := s.repo.GetUserByID(comment.UserID.Hex())
		if err != nil {
			// Skip comments with missing users
			continue
		}

		enhancedComments = append(enhancedComments, commentResponse{
			ID:        comment.ID,
			UserID:    comment.UserID,
			ProductID: comment.ProductID,
			Content:   comment.Content,
			Rating:    comment.Rating,
			CreatedAt: comment.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt: comment.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
			User: struct {
				ID        primitive.ObjectID `json:"_id"`
				Email     string             `json:"email"`
				FirstName string             `json:"firstName"`
				LastName  string             `json:"lastName"`
			}{
				ID:        user.ID,
				Email:     user.Email,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"comments": enhancedComments,
	})
}

func (s *Server) UpdateComment(c *gin.Context) {
	var req updateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request format: " + err.Error(),
		})
		return
	}

	// Get authenticated user
	email, _, _, exists := authenUser(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User not authenticated",
		})
		return
	}

	// Get the comment to verify ownership
	comment, err := s.repo.GetCommentByID(req.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Comment not found",
		})
		return
	}

	// Get user details
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to get user details",
		})
		return
	}

	// Verify ownership or admin status
	if comment.UserID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "Not authorized to edit this comment",
		})
		return
	}

	// Update the comment
	err = s.repo.UpdateComment(req.ID, req.Content, req.Rating)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to update comment: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "Comment updated successfully",
	})
}

func (s *Server) DeleteComment(c *gin.Context) {
	commentID := c.Param("id")
	if commentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Comment ID is required",
		})
		return
	}

	// Get authenticated user
	email, _, _, exists := authenUser(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User not authenticated",
		})
		return
	}

	// Get the comment to verify ownership
	comment, err := s.repo.GetCommentByID(commentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Comment not found",
		})
		return
	}

	// Get user details
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to get user details",
		})
		return
	}

	// Verify ownership or admin status
	if comment.UserID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "Not authorized to delete this comment",
		})
		return
	}

	// Delete the comment
	err = s.repo.DeleteComment(commentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to delete comment: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "Comment deleted successfully",
	})
}
