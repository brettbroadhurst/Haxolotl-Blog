// data/init_test.go - Test database connection
// Written by Brett Broadhurst <brettbroadhurst@gmail.com>

package data

import (
	log "log"
	os "os"
	testing "testing"
)

// Test Database.Init
func TestInit(t *testing.T) {
	l := log.New(os.Stdout, "TestInit: ", log.LstdFlags)
	db := CreateDatabaseClient(l)

	// Attempt to connect
	if err := db.Connect("brett", "brett", "172.17.0.3", "posts"); err != nil {
		t.Fatal("Could not connect to the database.")
	}
}
