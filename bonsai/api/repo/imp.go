package repo

import (
	"context"
	"errors"

	"github.com/trantho123/saleswebsite/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (i *Imp) CreateUser(user models.CreateUser) error {
	_, err := i.imp.Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

func (i *Imp) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := i.imp.Collection("users").FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (i *Imp) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := i.imp.Collection("users").FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (i *Imp) GetAllRoles() ([]models.Role, error) {
	var roles []models.Role
	cursor, err := i.imp.Collection("roles").Find(context.TODO(), bson.M{})
	if err != nil {
		return roles, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var role models.Role
		err := cursor.Decode(&role)
		if err != nil {
			return roles, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}

func (i *Imp) GetProduct(id string) (models.Product, error) {
	productId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Product{}, errors.New("invalid product id")
	}
	var product models.Product
	err = i.imp.Collection("products").FindOne(context.TODO(), bson.M{"_id": productId}).Decode(&product)
	if err != nil {
		return product, errors.New("failed to get product")
	}
	return product, nil
}

func (i *Imp) GetListProducts(page, pageSize int) ([]models.Product, int64, error) {
	var products []models.Product

	skip := (page - 1) * pageSize

	filter := bson.M{} // Thay đổi điều kiện nếu cần
	cursor, err := i.imp.Collection("products").Find(context.TODO(), filter,
		options.Find().SetSkip(int64(skip)).SetLimit(int64(pageSize)))
	if err != nil {
		return products, 0, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil {
			return products, 0, err
		}
		products = append(products, product)
	}

	total, err := i.imp.Collection("products").CountDocuments(context.TODO(), filter)
	if err != nil {
		return products, 0, err
	}

	return products, total, nil
}

// func (i *Imp) GetProductByCategory(category string) ([]models.Product, error) {
// 	var products []models.Product
// 	cursor, err := i.imp.Collection("products").Find(context.TODO(), bson.M{"category": category})
// 	if err != nil {
// 		return products, err
// 	}
// 	defer cursor.Close(context.Background())

// }

func (i *Imp) CreateCart(cart models.Cart) error {
	_, err := i.imp.Collection("carts").InsertOne(context.TODO(), cart)
	if err != nil {
		return err
	}
	return nil
}

func (i *Imp) GetCartByUserId(userID string) (models.Cart, error) {
	var cart models.Cart
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return cart, errors.New("invalid user id")
	}
	err = i.imp.Collection("carts").FindOne(context.TODO(), bson.M{"user": id}).Decode(&cart)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return cart, errors.New("mongo: no documents in result")
		}
		return cart, errors.New("failed to get cart")
	}
	return cart, nil
}

func (i *Imp) GetProductsInCart(userID string) ([]models.Item, error) {
	cart, err := i.GetCartByUserId(userID)
	if err != nil {
		return nil, err
	}
	return cart.Product, nil
}

func (i *Imp) UpdateCart(cart models.Cart) error {
	_, err := i.imp.Collection("carts").UpdateOne(context.TODO(), bson.M{"_id": cart.ID}, bson.M{"$set": cart})
	if err != nil {
		return err
	}

	return nil
}

func (i *Imp) IsProductExitsInCart(userID, productID string) bool {
	products, err := i.GetProductsInCart(userID)
	if err != nil {
		return false
	}
	for _, item := range products {
		if item.ProductID.Hex() == productID {
			return true
		}
	}
	return false
}

func (i *Imp) GetRole(id string) (models.Role, error) {
	roleID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Role{}, errors.New("invalid role id")
	}
	var role models.Role
	err = i.imp.Collection("roles").FindOne(context.TODO(), bson.M{"_id": roleID}).Decode(&role)
	if err != nil {
		return role, errors.New("failed to get role")
	}
	return role, nil
}

func (i *Imp) IsUserRole(id string) bool {
	role, err := i.GetRole(id)
	if err != nil {
		return false
	}
	if role.ID.Hex() == "" || role.Name != "User" {
		return false
	}
	return true
}

func (i *Imp) IsAdminRole(id string) bool {
	role, err := i.GetRole(id)
	if err != nil {
		return false
	}
	if role.ID.Hex() == "" || role.Name != "Admin" {
		return false
	}
	return role.Name == "admin"
}
func (i *Imp) IsCartExist(userID string) bool {
	cart, err := i.GetCartByUserId(userID)
	if err != nil {
		return false
	}
	if cart.ID.Hex() == "" {
		return false
	}
	return true
}

func (i *Imp) IsProductExist(id string) bool {
	product, err := i.GetProduct(id)
	if err != nil {
		return false
	}
	if product.ID.Hex() == "" {
		return false
	}
	return true
}

func (i *Imp) IsUserExist(id string) bool {
	user, err := i.GetUserByUsername(id)
	if err != nil {
		return false
	}
	if user.ID.Hex() == "" {
		return false
	}
	return true
}

func (i *Imp) IsEmailExist(email string) bool {
	user, err := i.GetUserByEmail(email)
	if err != nil {
		return false
	}
	if user.ID.Hex() == "" {
		return false
	}
	return true
}
