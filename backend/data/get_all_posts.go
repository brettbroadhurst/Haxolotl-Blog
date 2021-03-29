// get_all_posts.go - Get all posts from the DB
// Written by Brett Broadhurst <brettbroadhurst@gmail.com>

package data

import (
	context "context"
	bson "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
	time "time"
)

// Get all posts from the database
func (db *Database) GetAllPosts() ([]Post, error) {
	var (
		database *mongo.Database
		posts    *mongo.Collection
		ctx      context.Context
		allPosts []Post
		err      error
	)

	database = db.client.Database("blog")
	posts = database.Collection("posts")

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	// Get all documents
	cursor, err := posts.Find(ctx, bson.M{})
	if err != nil {
		db.logger.Println(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	// Get all documents in the collection
	for cursor.Next(ctx) {
		var post Post

		if err = cursor.Decode(&post); err != nil {
			db.logger.Println(err)
			return nil, err
		}

		allPosts = append(allPosts, post)
	}

	return allPosts, nil
}
