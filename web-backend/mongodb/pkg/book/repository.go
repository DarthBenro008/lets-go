package book

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"newsfeeder/pkg/entities"
)

type Repository interface {
}
type repository struct {
	Collection *mongo.Collection
}

func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}
func (r *repository) CreateBook(book *entities.Book) (*entities.Book, error) {
	book.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(context.Background(), book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *repository) ReadBook() (*[]entities.Book, error) {
	var books []entities.Book
	cursor, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var book entities.Book
		_ = cursor.Decode(&book)
		books = append(books, book)
	}
	return &books, nil
}

func (r *repository) UpdateBook(book *entities.Book) (*entities.Book, error) {
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": book.ID}, bson.M{"$set": book})
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *repository) DeleteBook(ID string) error {
	bookID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": bookID})
	if err != nil {
		return err
	}
	return nil
}
