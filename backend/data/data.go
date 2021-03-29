// data/data.go - Data definitions
// Written by Brett Broadhurst <brettbroadhurst@gmail.com>

package data

import (
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	//AuthorId  primitive.ObjectID  `bson:"authorId,omitempty"`
	Title     string              `json:"title" bson:"title"`
	Slug      string              `json:"slug" bson:"slug,omitempty"`
	Content   string              `json:"content" bson:"content"`
	Tags      []string            `json:"tags" bson:"tags,omitempty"`
	CreatedAt primitive.Timestamp `bson:"createdAt,omitempty"`
	UpdatedAt primitive.Timestamp `bson:"updatedAt,omitempty"`
}
