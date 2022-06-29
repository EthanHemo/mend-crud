package databases

import "fmt"

type ArrayDB struct {
	superHeroList []Superhero
}

func (d *ArrayDB) Init() {
	d.superHeroList = make([]Superhero, 0)
}

func (d *ArrayDB) AddSuperhero(s Superhero) error {
	d.superHeroList = append(d.superHeroList, s)
	fmt.Printf("I added %s, where can i find %s?\n", s.Name, s.Weakness)
	return nil
}

func (d *ArrayDB) DeleteSuperhero(name string) error {
	for i := range d.superHeroList {
		if d.superHeroList[i].Name == name {
			d.superHeroList = append(d.superHeroList[:i], d.superHeroList[i+1:]...)
			fmt.Printf("%s is DEAD to us now!\n", name)
			return nil
		}
	}

	return fmt.Errorf("Could not found %s in our party :( /n", name)
}

func (d *ArrayDB) UpdateSuperhero(name string, s Superhero) error {
	for i := range d.superHeroList {
		if d.superHeroList[i].Name == name {
			d.superHeroList[i] = s
			fmt.Printf("%s is updated.. wow.\n", name)
			return nil
		}
	}

	return fmt.Errorf("Could not found %s in our party :( /n", name)
}

func (d *ArrayDB) GetSuperhero(name string) (Superhero, error) {
	for i := range d.superHeroList {
		if d.superHeroList[i].Name == name {
			return d.superHeroList[i], nil
		}
	}

	return Superhero{}, fmt.Errorf("Could not found %s in our party :( /n", name)
}
