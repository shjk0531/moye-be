package repository

import (
	"context"
	"time"

	"github.com/shjk0531/moye-be/internal/domain/message/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	InsertMessage(msg *model.Message) error
	FindMessagesByChatRoomID(chatRoomID uint) ([]*model.Message, error)
}

type repository struct {
	collection *mongo.Collection
}

func NewRepository(mongoClient *mongo.Client, dbName string) Repository {
	collection := mongoClient.Database(dbName).Collection("messages")
	return &repository{collection: collection}
}

func (r *repository) InsertMessage(msg *model.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := r.collection.InsertOne(ctx, msg)
	return err
}

func (r *repository) FindMessagesByChatRoomID(chatRoomID uint) ([]*model.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"chatroom_id": chatRoomID}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var messages []*model.Message
	for cursor.Next(ctx) {
		var msg model.Message
		if err := cursor.Decode(&msg); err != nil {
			return nil, err
		}
		messages = append(messages, &msg)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return messages, nil
}
