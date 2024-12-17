package repo

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/trantho123/saleswebsite/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestGetUserByUsername(t *testing.T) {
	username := "john_doe"
	var user models.User
	err := i.imp.Collection("users").FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	require.NoError(t, err)
}

func TestGetUserByEmail(t *testing.T) {
	email := "john.doe@example.com"
	var user models.User
	err := i.imp.Collection("users").FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	require.NoError(t, err)
}

func TestGetProduct(t *testing.T) {
	id := "64b1b7a9f75a9a7b7e9f5c20"
	productId, err := primitive.ObjectIDFromHex(id)
	require.NoError(t, err)
	var product models.Product
	err = i.imp.Collection("products").FindOne(context.TODO(), bson.M{"_id": productId}).Decode(&product)
	require.NoError(t, err)
	require.NotEmpty(t, product)
	require.Equal(t, product.ID, productId)
}

func TestGetListProducts(t *testing.T) {
	page := 1
	pageSize := 10
	var products []models.Product

	skip := (page - 1) * pageSize

	filter := bson.M{} // Thay đổi điều kiện nếu cần
	cursor, err := i.imp.Collection("products").Find(context.TODO(), filter,
		options.Find().SetSkip(int64(skip)).SetLimit(int64(pageSize)))
	require.NoError(t, err)
	require.NoError(t, cursor.All(context.Background(), &products))
	require.NotEmpty(t, products)
	total, err := i.imp.Collection("products").CountDocuments(context.TODO(), filter)
	require.NoError(t, err)
	require.NotEmpty(t, total)
}

func TestGetAllRoles(t *testing.T) {
	var roles []models.Role
	cursor, err := i.imp.Collection("roles").Find(context.TODO(), bson.M{})
	require.NoError(t, err)
	require.NoError(t, cursor.All(context.Background(), &roles))
	require.NotEmpty(t, roles)
}

func TestCreateCart(t *testing.T) {
	user := models.Cart{
		User: primitive.NewObjectID(),
		Product: []models.Item{
			{
				ProductID: primitive.NewObjectID(),
				Price:     1000,
				Quantity:  1,
				CreatedAt: time.Now(),
			},
			{
				ProductID: primitive.NewObjectID(),
				Price:     2000,
				Quantity:  2,
				CreatedAt: time.Now(),
			},
		},
		Totals:    5000,
		CreatedAt: time.Now(),
	}
	err := i.CreateCart(user)
	require.NoError(t, err)

}
