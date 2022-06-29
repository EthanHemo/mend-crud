package databases

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const DB_FILE_NAME = "db_file.json"

type FileDB struct {
	superHeroList []Superhero
}

func (d FileDB) Init() {
	f, err := os.OpenFile(DB_FILE_NAME, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	byteValue, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("Error reading from file: %s\n")
		d.superHeroList = make([]Superhero, 0)
		return
	}
	if len(byteValue) == 0 {
		d.superHeroList = make([]Superhero, 0)
		return
	}

	err = json.Unmarshal(byteValue, &d.superHeroList)
	if err != nil {
		fmt.Printf("Error reading from file: %s\n")
		d.superHeroList = make([]Superhero, 0)
		return
	}

}

func (d FileDB) AddSuperhero(s Superhero) error {
	d.superHeroList = append(d.superHeroList, s)
	fmt.Printf("I added %s, where can i find %s?\n", s.Name, s.Weakness)
	return d.saveListToFile()
}

func (d FileDB) DeleteSuperhero(name string) error {
	for i := range d.superHeroList {
		if d.superHeroList[i].Name == name {
			d.superHeroList = append(d.superHeroList[:i], d.superHeroList[i+1:]...)
			fmt.Printf("%s is DEAD to us now!\n", name)
			return d.saveListToFile()
		}
	}

	return fmt.Errorf("Could not found %s in our party :( /n", name)
}

func (d FileDB) UpdateSuperhero(name string, s Superhero) error {
	for i := range d.superHeroList {
		if d.superHeroList[i].Name == name {
			d.superHeroList[i] = s
			fmt.Printf("%s is updated.. wow.\n", name)
			return d.saveListToFile()
		}
	}

	return fmt.Errorf("Could not found %s in our party :( /n", name)
}

func (d FileDB) GetSuperhero(name string) (Superhero, error) {
	for i := range d.superHeroList {
		if d.superHeroList[i].Name == name {
			return d.superHeroList[i], nil
		}
	}

	return Superhero{}, fmt.Errorf("Could not found %s in our party :( /n", name)
}

func (d FileDB) saveListToFile() error {
	data, _ := json.MarshalIndent(d.superHeroList, "", " ")
	err := ioutil.WriteFile(DB_FILE_NAME, data, 644)
	if err != nil {
		return fmt.Errorf("Error saving file: %s\n", err)
	}
	return nil
}
