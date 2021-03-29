// data/get_all_posts_test.go - Test getting all posts
// Written by Brett Broadhurst <brettbroadhurst@gmail.com>

package data

import (
	log "log"
	os "os"
	testing "testing"
)

// Test getting all posts from the database
func TestGetAllPosts(t *testing.T) {
	l := log.New(os.Stdout, "TestGetAllPosts", log.LstdFlags)

	db := CreateDatabaseClient(l)

	// Attempt to connect
	if err := db.Connect("brett", "brett", "172.17.0.2", "posts"); err != nil {
		t.Fatal("Could not connect to the database.")
	}

	// Posts
	posts, err := db.GetAllPosts()
	if err != nil {
		t.Fatal("Could not get posts")
	}

	if len(posts) <= 0 {
		t.Fatal("There are no posts")
	}
}
