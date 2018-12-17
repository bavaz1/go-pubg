package database

import (
	"testing"
)

func TestPlayers(t *testing.T) {
	m := New()
	p := m.CreatePlayer(Player{
		Name:       "TestName",
		Email:      "test2@gmail.com",
		InGameName: "bROkedli",
	})
	if p.Name != "TestName" {
		t.Errorf("Name should be 'TestName', instead of %s", p.Name)
	}
	if p.Email != "test2@gmail.com" {
		t.Errorf("Email should be 'test2@gmail.com', instead of %s", p.Email)
	}
	if p.InGameName != "bROkedli" {
		t.Errorf("InGameName should be 'bROkedli', instead of %s", p.InGameName)
	}
	t.Log(p)

	p2 := m.GetPlayer(p)
	if p2.Name != "TestName" {
		t.Errorf("Name should be 'TestName', instead of %s", p2.Name)
	}
	if p2.Email != "test2@gmail.com" {
		t.Errorf("Email should be 'test2@gmail.com', instead of %s", p2.Email)
	}
	if p2.InGameName != "bROkedli" {
		t.Errorf("InGameName should be 'bROkedli', instead of %s", p2.InGameName)
	}
	t.Log(p2)

	p3 := m.UpdatePlayer(Player{
		ID:         p2.ID,
		Name:       "ChangedName",
		Email:      p2.Email,
		InGameName: "bro",
	})
	if p3.Name != "ChangedName" {
		t.Errorf("Name should be 'ChangedName', instead of %s", p3.Name)
	}
	if p3.Email != "test2@gmail.com" {
		t.Errorf("Email should be 'test2@gmail.com', instead of %s", p3.Email)
	}
	if p3.InGameName != "bro" {
		t.Errorf("InGameName should be 'bro', instead of %s", p3.InGameName)
	}
	t.Log(p3)

	p4 := m.GetPlayer(Player{
		ID: p2.ID,
	})
	if p4.Name != "ChangedName" {
		t.Errorf("Name should be 'ChangedName', instead of %s", p4.Name)
	}
	if p4.Email != "test2@gmail.com" {
		t.Errorf("Email should be 'test2@gmail.com', instead of %s", p4.Email)
	}
	if p4.InGameName != "bro" {
		t.Errorf("InGameName should be 'bro', instead of %s", p4.InGameName)
	}
	t.Log(p4)

	p5 := m.DeletePlayer(Player{
		ID: p2.ID,
	})
	if p5.Name == "ChangedName" {
		t.Errorf("Name should be %s, instead of 'ChangedName'", p4.Name)
	}
	if p5.Email == "test2@gmail.com" {
		t.Errorf("Email should be %s, instead of 'test2@gmail.com'", p4.Email)
	}
	if p5.InGameName == "bro" {
		t.Errorf("InGameName should be %s, instead of 'bro'", p4.InGameName)
	}
	t.Log(p5)

	p6 := m.GetPlayers()
	if len(p6) == 0 {
		t.Errorf("Players array length should be not 0")
	}
	t.Log(p6)
}
