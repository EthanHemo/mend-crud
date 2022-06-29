package databases

import "fmt"

type Crud interface {
	Init()

	AddSuperhero(s Superhero) error

	DeleteSuperhero(name string) error

	UpdateSuperhero(name string, s Superhero) error

	GetSuperhero(name string) (Superhero, error)
}

type DatabaseType int64

const (
	Array DatabaseType = iota
	File
	SQL
)

func GetDatabase(dbType DatabaseType) Crud {
	switch dbType {
	case Array:
		c := ArrayDB{}
		c.Init()
		return &c
	case File:
		c := FileDB{}
		c.Init()
		return c
	default:
		fmt.Println("No db selected")
		return nil
	}
}
