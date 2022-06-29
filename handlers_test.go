package main

import (
	"mend-crud/databases"
	"testing"
)

func TestCrud(t *testing.T) {
	c := databases.GetDatabase(databases.Array)
	s := databases.Superhero{
		Name:     "Chicken man",
		Age:      42,
		Power:    "Can talk to chickens",
		Weakness: "His mom",
	}

	if err := c.AddSuperhero(s); err != nil {
		t.Errorf("Failed to save hero: %s", err)
	}

	clone, err := c.GetSuperhero(s.Name)
	if err != nil {
		t.Errorf("Failed to get hero %s", err)
	}

	if clone.Name != s.Name || clone.Age != s.Age || clone.Power != s.Power || clone.Weakness != s.Weakness {
		t.Errorf("retrived hero not as saved.\n original: %+v\n clone: %+v", s, clone)
	}

	s.Power = "Can transform into chicken"

	if err := c.UpdateSuperhero(s.Name, s); err != nil {
		t.Errorf("Couldn't update hero: %s", err)
	}

	clone, err = c.GetSuperhero(s.Name)
	if err != nil {
		t.Errorf("Failed to get hero %s", err)
	}

	if clone.Name != s.Name || clone.Age != s.Age || clone.Power != s.Power || clone.Weakness != s.Weakness {
		t.Errorf("retrived hero not as updated.\n original: %+v\n clone: %+v", s, clone)
	}

}
