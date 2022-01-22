package singleton

import (
	"fmt"
	"sync"
)

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

// thread safety
// sync.Once
// init

// laziness

var once sync.Once

var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {
	// this is thread safety and lazy instantiation
	once.Do(func() {
		db := singletonDatabase{map[string]int{"Bucuresti": 4445550, "Berlin": 44553311}}
		instance = &db
	})
	return instance
}

func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDatabase().GetPopulation(city)
	}
	return result
}

func SingletonExample() {
	db := GetSingletonDatabase()
	fmt.Println(db.GetPopulation("Berlin"))
	tp := GetTotalPopulation([]string{"Berlin", "Bucuresti"})
	ok := tp == (4445550 + 44553311)
	fmt.Println(ok)
}
