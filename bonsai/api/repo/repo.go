package repo

import (
	"github.com/trantho123/saleswebsite/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo interface {
	CreateUser(user models.CreateUser) error
	GetUserByEmail(email string) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	GetProduct(id string) (models.Product, error)
	GetListProducts(page, pageSize int) ([]models.Product, int64, error)
	GetAllRoles() ([]models.Role, error)
	IsEmailExist(email string) bool
	IsUserExist(id string) bool
	IsProductExist(id string) bool
	GetCartByUserId(userID string) (models.Cart, error)
	GetProductsInCart(userID string) ([]models.Item, error)
	IsProductExitsInCart(userID, productID string) bool
	CreateCart(cart models.Cart) error
	UpdateCart(cart models.Cart) error
	IsCartExist(userID string) bool
	IsAdminRole(id string) bool
	CreateAccessToken(token models.AccessToken) error
	IsUserRole(id string) bool
	GetAccessToken(token string) (models.AccessToken, error)
	DeleteAccessToken(token string) error
}

type Imp struct {
	imp *mongo.Database
}

func NewRepo(i *mongo.Database) Repo {
	return &Imp{imp: i}
}
