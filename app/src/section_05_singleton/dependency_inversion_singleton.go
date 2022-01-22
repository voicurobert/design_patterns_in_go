package singleton

import "fmt"

type Database interface {
	GetPopulation(name string) int
}

func GetTotalPopulation2(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city)
	}
	return result
}

type dummyDatabase struct {
	dummyData map[string]int
}

func (db *dummyDatabase) GetPopulation(name string) int {
	if len(db.dummyData) == 0 {
		db.dummyData = map[string]int{
			"alpha": 1,
			"beta":  2,
			"gamma": 3,
		}
	}
	return db.dummyData[name]
}

func DependencyInversionSingletonExample() {
	names := []string{"alpha", "gamma"}
	tp := GetTotalPopulation2(&dummyDatabase{}, names)
	fmt.Println(tp)
}
