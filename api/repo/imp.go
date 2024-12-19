package repo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/trantho123/saleswebsite/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (i *Imp) UpdateUserVerification(userID string, verified bool) error {
	_, err := i.imp.Collection("users").UpdateOne(context.TODO(), bson.M{"_id": userID}, bson.M{"$set": bson.M{"verified": verified}})
	if err != nil {
		return err
	}
	return nil
}

func (i *Imp) GetUserByVerificationCode(code string) (models.User, error) {
	var user models.User
	err := i.imp.Collection("users").FindOne(context.TODO(), bson.M{"verificationcode": code}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (i *Imp) UpdateUser(user models.User) error {
	_, err := i.imp.Collection("users").UpdateOne(context.TODO(), bson.M{"email": user.Email}, bson.M{"$set": user})
	if err != nil {
		return err
	}
	return nil
}

func (i *Imp) DeleteUser(email string) error {
	_, err := i.imp.Collection("users").DeleteOne(context.TODO(), bson.M{"email": email})
	if err != nil {
		return err
	}
	return nil
}

func (i *Imp) ResetPassword(email, newPassword string) error {
	_, err := i.imp.Collection("users").UpdateOne(context.TODO(), bson.M{"email": email}, bson.M{"$set": bson.M{"password": newPassword}})
	if err != nil {
		return err
	}
	return nil
}

func (i *Imp) CreateUser(user models.User) error {
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

func (i *Imp) GetListProducts() ([]models.Product, error) {
	var products []models.Product
	cursor, err := i.imp.Collection("products").Find(context.TODO(), bson.M{})
	if err != nil {
		return products, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil {
			return products, err
		}
		products = append(products, product)
	}
	return products, nil
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

func (i *Imp) GetItemsInCart(userID string) ([]models.Item, error) {
	cart, err := i.GetCartByUserId(userID)
	if err != nil {
		return nil, err
	}
	return cart.Items, nil
}

func (i *Imp) IsItemInCart(userID, productID string) bool {
	items, err := i.GetItemsInCart(userID)
	if err != nil {
		return false
	}
	for _, item := range items {
		if item.ProductID.Hex() == productID {
			return true
		}
	}
	return false
}

func (i *Imp) UpdateCart(cart models.Cart) error {
	_, err := i.imp.Collection("carts").UpdateOne(
		context.TODO(),
		bson.M{"_id": cart.ID},
		bson.M{"$set": cart},
	)
	return err
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
func (i *Imp) GetAllOrders() ([]models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := i.imp.Collection("orders")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []models.Order
	if err = cursor.All(ctx, &orders); err != nil {
		return nil, err
	}

	return orders, nil
}

func (i *Imp) IsAdminRole(id string) bool {
	role, err := i.GetRole(id)
	if err != nil {
		return false
	}
	if role.ID.Hex() == "" || role.Name != "Admin" {
		return false
	}
	return role.Name == "Admin"
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

func (i *Imp) GetProductsByTags(tags []string) ([]models.Product, error) {
	var products []models.Product

	filter := bson.M{
		"tags.name": bson.M{
			"$in": tags,
		},
	}

	cursor, err := i.imp.Collection("products").Find(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to query products: %v", err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, fmt.Errorf("failed to decode product: %v", err)
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return products, nil
}

func (i *Imp) CreateOrder(order models.Order) (string, error) {
	order.ID = primitive.NewObjectID()
	_, err := i.imp.Collection("orders").InsertOne(context.TODO(), order)
	if err != nil {
		return "", err
	}
	return order.ID.Hex(), nil
}

func (i *Imp) GetOrderByID(orderID string) (models.Order, error) {
	var order models.Order
	id, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return order, err
	}
	err = i.imp.Collection("orders").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&order)
	return order, err
}

func (i *Imp) GetOrdersByUserID(userID string) ([]models.Order, error) {
	var orders []models.Order
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return orders, err
	}
	cursor, err := i.imp.Collection("orders").Find(context.TODO(), bson.M{"user_id": id})
	if err != nil {
		return orders, err
	}
	err = cursor.All(context.TODO(), &orders)
	return orders, err
}

func (i *Imp) UpdateOrderStatus(orderID string, status string) error {
	id, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return err
	}
	_, err = i.imp.Collection("orders").UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.M{"$set": bson.M{
			"status":     status,
			"updated_at": time.Now(),
		}},
	)
	return err
}

func (i *Imp) UpdatePaymentStatus(orderID string, status string, transactionID string) error {
	id, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return err
	}
	_, err = i.imp.Collection("orders").UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.M{"$set": bson.M{
			"payment_details.status":         status,
			"payment_details.transaction_id": transactionID,
			"payment_details.paid_at":        time.Now(),
			"updated_at":                     time.Now(),
		}},
	)
	return err
}

func (i *Imp) CreateComment(comment models.Comment) error {
	comment.ID = primitive.NewObjectID()
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()

	_, err := i.imp.Collection("reviews").InsertOne(context.TODO(), comment)
	return err
}

func (i *Imp) GetCommentsByProductID(productID string) ([]models.Comment, error) {
	var comments []models.Comment

	id, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return nil, err
	}

	cursor, err := i.imp.Collection("reviews").Find(
		context.TODO(),
		bson.M{"product_id": id},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	err = cursor.All(context.TODO(), &comments)
	return comments, err
}

func (i *Imp) UpdateComment(commentID string, content string, rating float32) error {
	id, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return err
	}

	_, err = i.imp.Collection("reviews").UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.M{
			"$set": bson.M{
				"content":    content,
				"rating":     rating,
				"updated_at": time.Now(),
			},
		},
	)
	return err
}

func (i *Imp) DeleteComment(commentID string) error {
	id, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return err
	}

	_, err = i.imp.Collection("reviews").DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

func (i *Imp) GetUserByID(id string) (models.User, error) {
	var user models.User

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	err = i.imp.Collection("users").FindOne(
		context.TODO(),
		bson.M{"_id": objectID},
	).Decode(&user)

	return user, err
}

func (i *Imp) GetCommentByID(commentID string) (models.Comment, error) {
	var comment models.Comment

	id, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return comment, err
	}

	err = i.imp.Collection("reviews").FindOne(
		context.TODO(),
		bson.M{"_id": id},
	).Decode(&comment)

	return comment, err
}

func (i *Imp) GetAllUsers() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := i.imp.Collection("users")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}
