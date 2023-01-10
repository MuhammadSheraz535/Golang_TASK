package services

import (
	"context"
	"errors"

	"githhub.com/shery535/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImpl) GetUser(*string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "user_name", Value: user.Name}}
	u.usercollection.FindOne(u.ctx, query).Decode(&user)
	return user, nil
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	var users []*models.User
	cursor, err := u.usercollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user models.User
		err = cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("document not found")
	}

	cursor.Close(u.ctx)
	return nil, nil

}
func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	filter := bson.D{bson.E{Key: "user_name", Value: user.Name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "user_name", Value: user.Name}, bson.E{Key: "user_age", Value: user.Age}, bson.E{Key: "user_address", Value: user.Address}}}}
	res, _ := u.usercollection.UpdateOne(u.ctx, filter, update)
	if res.MatchedCount != 1 {
		return errors.New("matched documents not found for update")
	}
	return nil

}
func (u *UserServiceImpl) DeleteUser(*string) error {
	filter := bson.D{bson.E{Key: "user_name", Value: u.usercollection.Name}}
	res, _ := u.usercollection.DeleteOne(u.ctx, filter)
	if res.DeletedCount != 1 {
		return errors.New("matched documents not found for delete")
	}

	return nil

}
