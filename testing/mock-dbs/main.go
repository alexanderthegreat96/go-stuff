package main

import (
	"fmt"
	"log"
)

// inspired by the original exercise-68 by Todd McLeod

// gonna create a way to actually store data
// just like it would be stored in a database
// for this, the stuff will be stored in memory
// so we are going to use some type of pointer for this

// define table structure for user data

type User struct {
	ID       int
	Username string
	Password string
	Email    string
	Age      int
}

// define table structure for Orders

type Order struct {
	ID          int
	UserID      int
	ProductID   int
	ProductName string
	Qty         int
	UnitPrice   float64
}

// Database
type Database struct {
	Users  map[int]User
	Orders map[int]Order
}

// Initiate a new database

func NewDabase() *Database {
	return &Database{
		Users:  make(map[int]User),
		Orders: make(map[int]Order),
	}
}

// add user will he hooked to Database
func (db *Database) AddUser(user User) error {
	_, ok := db.Users[user.ID]

	if ok {
		return fmt.Errorf("User %v, already exists", user.ID)
	}
	db.Users[user.ID] = user
	return nil
}

// add the order
func (db *Database) AddOrder(order Order) {
	db.Orders[order.ID] = order
}

// gets the user if exists
func (db *Database) GetUser(userId int) (User, bool) {
	user, exists := db.Users[userId]
	return user, exists
}

// let's get the orders

func (db *Database) GetUserOrders(userId int) (Order, bool) {
	order, exists := db.Orders[userId]
	return order, exists
}

// create a new database

func main() {
	db := NewDabase()

	err := db.AddUser(User{
		ID:       1,
		Username: "admin",
		Password: "admin",
		Email:    "admin@localhost.com",
		Age:      27,
	})

	if err != nil {
		log.Fatalf("we got a problem. %v", err)
	}

	user, ok := db.GetUser(1)

	if ok {
		fmt.Println(user)
	}

}
