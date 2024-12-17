package repo

import (
	"github.com/trantho123/saleswebsite/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo interface {
	CreateUser(user models.User) error
	GetUserByEmail(email string) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	GetUserByID(id string) (models.User, error)
	GetProduct(id string) (models.Product, error)
	GetListProducts() ([]models.Product, error)
	GetAllRoles() ([]models.Role, error)
	IsEmailExist(email string) bool
	UpdateUser(user models.User) error
	IsUserExist(id string) bool
	IsProductExist(id string) bool
	GetCartByUserId(userID string) (models.Cart, error)
	GetItemsInCart(userID string) ([]models.Item, error)
	IsItemInCart(userID, productID string) bool
	CreateCart(cart models.Cart) error
	UpdateCart(cart models.Cart) error
	IsCartExist(userID string) bool
	IsAdminRole(id string) bool
	CreateAccessToken(token models.AccessToken) error
	IsUserRole(id string) bool
	ResetPassword(email, newPassword string) error
	DeleteUser(email string) error
	GetUserByVerificationCode(code string) (models.User, error)
	UpdateUserVerification(userID string, verified bool) error
	GetProductsByTags(tags []string) ([]models.Product, error)
	CreateOrder(order models.Order) (string, error)
	GetOrderByID(orderID string) (models.Order, error)
	GetOrdersByUserID(userID string) ([]models.Order, error)
	UpdateOrderStatus(orderID string, status string) error
	UpdatePaymentStatus(orderID string, status string, transactionID string) error
	CreateComment(comment models.Comment) error
	GetCommentsByProductID(productID string) ([]models.Comment, error)
	UpdateComment(commentID string, content string, rating float32) error
	DeleteComment(commentID string) error
	GetCommentByID(commentID string) (models.Comment, error)
}

type Imp struct {
	imp *mongo.Database
}

func NewRepo(i *mongo.Database) Repo {
	return &Imp{imp: i}
}
