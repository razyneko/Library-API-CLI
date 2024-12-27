package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// model for book

type Book struct {
	ID 		primitive.ObjectID		`bson:"_id,omitempty" json:"id"`
	Title 	string  				`bson:"title" json:"title"`
	Author 	string					`bson:"author" json:"author"`
	Stock 	int  					`bson:"stock" json:"stock"`
}
