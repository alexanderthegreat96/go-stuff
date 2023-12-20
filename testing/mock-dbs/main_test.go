package main

import "testing"

func TestAddUser(t *testing.T) {
	db := NewDabase()

	err := db.AddUser(User{
		ID:       1,
		Username: "admin",
		Password: "admin",
		Email:    "admin@localhost.com",
		Age:      27,
	})

	if err != nil {
		t.Errorf("we got a problem. %v", err)
	}

	if _, ok := db.GetUser(1); !ok {
		t.Errorf("want %v, got: nothing", 1)
	}

}
