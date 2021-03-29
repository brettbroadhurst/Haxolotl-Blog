// data/create_post.go
// Written by Brett Broadhurst <brettbroadhurst@gmail.com>

package data

import (
	context "context"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	time "time"
)

// Create a new post in the database
func (db *Database) CreatePost(title, content string) *Post {
	database := db.client.Database("blog")
	posts := database.Collection("posts")

	// New Post
	newPost := Post{
		Title:   title,
		Content: content,
	}

	// Create a new execution context
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	// Insert the post into the table
	res, err := posts.InsertOne(ctx, newPost)
	if err != nil {
		db.logger.Println(err)
		return nil
	}

	return &Post{Id: res.InsertedID.(primitive.ObjectID), Title: title, Content: content}
}
