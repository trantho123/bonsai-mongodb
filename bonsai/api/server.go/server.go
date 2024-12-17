package server

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/trantho123/saleswebsite/repo"
	"github.com/trantho123/saleswebsite/utils"
)

type Server struct {
	router *gin.Engine
	repo   repo.Repo
	config *utils.Config
}

func NewServer(repo repo.Repo, config *utils.Config) *Server {
	server := &Server{
		router: gin.Default(),
		repo:   repo,
		config: config,
	}

	server.Router()

	return server
}

func (server *Server) Router() {
	server.corsConfig()
	server.router.POST("/register", server.CreateUser)
	server.router.POST("/login", server.Login)
	server.router.POST("/forgot-password", server.ForgotPassword)
	server.router.POST("/reset-password", server.ResetPasswordWithToken)
	server.router.GET("/products", server.GetListProducts)
	server.router.GET("/product/:id", server.GetProduct)
	server.router.POST("/products/tags", server.GetProductByTags)
	server.router.GET("/verify/:code", server.VerifyEmail)
	server.router.GET("/products/:id/comments", server.GetProductComments)

	authRoutes := server.router.Group("/auth").Use(authMiddleware(server.config.AccessTokenKey))
	authRoutes.GET("/profile", server.GetProfile)
	authRoutes.PUT("/profile", server.UpdateProfile)
	authRoutes.POST("/cart", server.CreateCart)
	authRoutes.GET("/cart", server.GetCart)
	authRoutes.DELETE("/cart", server.DeleteProductInCart)
	authRoutes.PUT("/cart/quantity", server.UpdateQuantityProductCart)
	authRoutes.PUT("/resetpassword", server.ResetPassword)
	authRoutes.GET("/logout", server.Logout)
	authRoutes.POST("/orders", server.CreateOrder)
	authRoutes.GET("/orders", server.GetUserOrders)
	authRoutes.GET("/orders/:id", server.GetOrder)
	authRoutes.POST("/comments", server.CreateComment)
	authRoutes.PUT("/comments", server.UpdateComment)
	authRoutes.DELETE("/comments/:id", server.DeleteComment)

}

func (server *Server) corsConfig() {

	server.router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,                                                // Cho phép tất cả origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // Các phương thức cho phép
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Các header cho phép
		AllowCredentials: true,                                                // Cho phép gửi cookie hoặc thông tin xác thực
		MaxAge:           12 * time.Hour,                                      // Thời gian tối đa cho phép lưu CORS trong cache của trình duyệt
	}))
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
