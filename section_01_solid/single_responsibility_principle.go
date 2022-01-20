package solid

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {

}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// separation of concerns

// these methods brake the principle. It should be declared on separate package

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {

}

func (j *Journal) LoadFromWeb(url *url.URL) {

}

// this should be moved on another package

var LineSeparator = "\n"

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

type Persistence struct {
	lineSeparator string
}

func SingleResponsibilityPrinciple() {
	fmt.Println("Single Responsibility Principle...")
	j := Journal{}
	j.AddEntry("I cried today")
	j.AddEntry("I ate a bug")
	fmt.Println(j.String())

	// SaveToFile()

	//p := Persistence{lineSeparator: "\r\n"}
	//p.SaveToFile(&j, "journal.txt")
	fmt.Println("#############################################################")
	fmt.Println()
}
