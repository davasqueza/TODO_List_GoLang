package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"todo-list/config"
	"todo-list/models"
)

type TodoRepository struct {
	db *mongo.Collection
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{
		db: config.AppConfig.DatabaseClient.Collection("todo"),
	}
}

func (todoRepository *TodoRepository) FindAll(limit int) ([]*models.Todo, error) {
	var list = []*models.Todo{}

	var cur, err = todoRepository.db.Aggregate(context.TODO(), []bson.D{
		{{"$match", bson.D{}}},
		{{"$limit", limit}},
	})

	if err != nil {
		return nil, err
	}

	defer func() { _ = cur.Close(context.TODO()) }()
	for cur.Next(context.TODO()) {
		var result = &models.Todo{}
		err = cur.Decode(result)
		if err != nil {
			return nil, err
		}
		list = append(list, result)
	}

	return list, err
}

func (todoRepository *TodoRepository) FindById(id string) (*models.Todo, error) {
	var result = &models.Todo{}

	var objectId, err = primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	err = todoRepository.db.FindOne(context.TODO(), bson.D{{"_id", objectId}}).Decode(result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (todoRepository *TodoRepository) Create(todo *models.Todo) (*mongo.InsertOneResult, error) {
	return todoRepository.db.InsertOne(context.TODO(), todo)
}
