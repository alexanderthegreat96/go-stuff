package main

import (
	"fmt"
	"log"
)

type Player struct {
	ID       int
	Username string
	Platform string
}

// store all players inside a map
type PlayerStore struct {
	Players map[int]Player
}

// store the data
func (ps PlayerStore) SavePlayer(p Player) error {
	if _, ok := ps.Players[p.ID]; !ok {
		ps.Players[p.ID] = p
		return nil
	}
	return fmt.Errorf("player with ID %d already exists", p.ID)
}

// if found, returns a value of Type player
func (ps PlayerStore) GetPlayer(id int) (Player, error) {
	player, ok := ps.Players[id]

	if ok {
		return player, nil
	}
	return Player{}, fmt.Errorf("player with ID %d, was not found", id)
}

// whatever implements this, MUST IMPLMENET THE 2 METHODS
type PlayerInterface interface {
	SavePlayer(p Player) error
	GetPlayer(id int) (Player, error)
}

// implement the interface inside the service
// so that by calling the service we are able
// to add, get and list data

type PlayerService struct {
	ps PlayerInterface
}

// create the methods and hook them into the service

func (s PlayerService) SavePlayer(p Player) error {
	return s.ps.SavePlayer(p)
}

func (s PlayerService) GetPlayer(id int) (Player, error) {
	return s.ps.GetPlayer(id)
}

func main() {
	players := PlayerStore{
		Players: make(map[int]Player),
	}

	PlayerService := PlayerService{
		ps: players,
	}

	save_player := PlayerService.SavePlayer(Player{
		ID:       1,
		Username: "alexanderdth",
		Platform: "psn",
	})

	if save_player != nil {
		log.Fatalf("error %s", save_player)
	}

	player, err := PlayerService.GetPlayer(1)

	if err == nil {
		fmt.Println(player)
	} else {
		log.Fatalf("error %s", err)
	}

}
