package repo

import (
	"context"

	"github.com/trantho123/saleswebsite/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (i *Imp) CreateAccessToken(token models.AccessToken) error {
	_, err := i.imp.Collection("access_tokens").InsertOne(context.TODO(), token)
	if err != nil {
		return err
	}
	return nil
}

func (i *Imp) GetAccessToken(token string) (models.AccessToken, error) {
	var accessToken models.AccessToken
	err := i.imp.Collection("access_tokens").FindOne(context.TODO(), bson.M{"accesstoken": token}).Decode(&accessToken)
	if err != nil {
		return accessToken, err
	}
	return accessToken, nil
}

func (i *Imp) DeleteAccessToken(token string) error {
	_, err := i.imp.Collection("access_tokens").DeleteOne(context.TODO(), bson.M{"accesstoken": token})
	if err != nil {
		return err
	}
	return nil
}
