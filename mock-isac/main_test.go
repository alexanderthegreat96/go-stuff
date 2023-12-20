package main

import "testing"

func TestGetPlayer(t *testing.T) {
	md := &PlayerStore{
		Players: map[int]Player{
			2: {ID: 2, Username: "seth121", Platform: "uplay"},
		},
	}
	s := &PlayerService{
		ps: md,
	}

	p, err := s.GetPlayer(2)

	if err != nil {
		t.Errorf("go error %v", err)
	}

	if p.Username != "seth121" {
		t.Errorf("got %s. want: %s", p.Username, "seth121")
	}

	if p.Platform != "uplay" {
		t.Errorf("got %s. want: %s", p.Platform, "uplay")
	}

}

func TestSavePlayer(t *testing.T) {
	md := &PlayerStore{
		Players: map[int]Player{
			2: {ID: 2, Username: "seth121", Platform: "uplay"},
		},
	}

	s := &PlayerService{
		ps: md,
	}

	err := s.SavePlayer(Player{
		ID:       3,
		Username: "Alex",
		Platform: "uplay",
	})

	if err != nil {
		t.Errorf("go error: %v", err)
	}
}
