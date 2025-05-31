package repository

import (
	"context"
	"time"

	"github.com/Masum-Osman/lex-scope/modules/text/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TextRepository interface {
	Save(ctx context.Context, text *domain.Text) (string, error)
	GetByID(ctx context.Context, id string) (*domain.Text, error)
	Update(ctx context.Context, id string, text *domain.Text) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.Text, error)
}

type mongoTextRepository struct {
	collection *mongo.Collection
}

func NewTextRepository(db *mongo.Database) TextRepository {
	return &mongoTextRepository{
		collection: db.Collection("texts"),
	}
}

/* This was for TDD mock Insert
func (r *mongoTextRepository) Save(ctx context.Context, text *domain.Text) (string, error) {

	text.ID = ""
	_, err := r.collection.InsertOne(ctx, text)
	return err

}
*/

func (r *mongoTextRepository) Save(ctx context.Context, text *domain.Text) (string, error) {
	text.ID = primitive.NewObjectID()
	text.CreatedAt = time.Now()
	text.UpdatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, text)
	return text.ID.Hex(), err
}

func (r *mongoTextRepository) GetByID(ctx context.Context, id string) (*domain.Text, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	var t domain.Text
	err := r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&t)
	return &t, err
}

/*
Mock Update function for TDD
func (r *mongoTextRepository) Update(ctx context.Context, text *domain.Text) error {
	filter := bson.M{"_id": text.ID}
	update := bson.M{"$set": bson.M{"content": text.Content, "user_id": text.UserID}}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}
*/

func (r *mongoTextRepository) Update(ctx context.Context, id string, updated *domain.Text) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	updated.UpdatedAt = time.Now()
	_, err := r.collection.ReplaceOne(ctx, bson.M{"_id": objID}, updated)
	return err
}

/*
Mock Delete function for TDD
func (r *mongoTextRepository) Delete(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
*/

func (r *mongoTextRepository) Delete(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func (r *mongoTextRepository) List(ctx context.Context) ([]domain.Text, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var texts []domain.Text
	for cursor.Next(ctx) {
		var t domain.Text
		if err := cursor.Decode(&t); err == nil {
			texts = append(texts, t)
		}
	}
	return texts, nil
}
