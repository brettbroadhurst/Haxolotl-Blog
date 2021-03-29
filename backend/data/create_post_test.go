// data/create_post_test.go - Create post test
// Written by Brett Broadhurst <brettbroadhurst@gmail.com>

package data

import (
	log "log"
	os "os"
	testing "testing"
)

// Test Create Post
func TestCreatePost(t *testing.T) {
	title := "Test Post"
	content := "This is a test post"

	l := log.New(os.Stdout, "TestCreatePost: ", log.LstdFlags)
	db := CreateDatabaseClient(l)

	// Attempt to connect
	if err := db.Connect("172.17.0.2", "posts"); err != nil {
		t.Fatal("Could not connect to the database.")
	}

	// Create a new post
	post := db.CreatePost(title, content)
	if post == nil {
		t.Fatal("Could not create post.")
	}

	// Test values
	if post.Title != title {
		t.Fatal("Post title does not pass.")
	} else if post.Content != content {
		t.Fatal("Post content does not pass.")
	}

	l.Printf("%+v\n", post)
}
