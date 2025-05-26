package repository

import (
	"context"

	"github.com/Masum-Osman/lex-scope/modules/text/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TextRepository interface {
	Save(ctx context.Context, text *domain.Text) error
	GetByID(ctx context.Context, id string) (*domain.Text, error)
	Update(ctx context.Context, text *domain.Text) error
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

func (r *mongoTextRepository) Save(ctx context.Context, text *domain.Text) error {
	text.ID = ""
	_, err := r.collection.InsertOne(ctx, text)
	return err
}

func (r *mongoTextRepository) GetByID(ctx context.Context, id string) (*domain.Text, error) {
	var text domain.Text
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&text)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &text, nil
}

func (r *mongoTextRepository) Update(ctx context.Context, text *domain.Text) error {
	filter := bson.M{"_id": text.ID}
	update := bson.M{"$set": bson.M{"content": text.Content, "user_id": text.UserID}}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *mongoTextRepository) Delete(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
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
