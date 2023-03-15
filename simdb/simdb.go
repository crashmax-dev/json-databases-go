package main

import (
	"fmt"

	"github.com/sonyarouje/simdb"
)

type Emote struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

func (e Emote) ID() (jsonField string, value interface{}) {
	value = e.Name
	jsonField = "name"
	return
}

func main() {
	driver, err := simdb.New("db/simdb")
	if err != nil {
		panic(err)
	}

	err = driver.Insert(Emote{
		Name: "Madge",
		Link: "https://cdn.7tv.app/emote/60a95f109d598ea72fad13bd/4x.webp",
	})
	if err != nil {
		fmt.Println(err)
	}

	var emote Emote
	err = driver.
		Open(Emote{}).
		Where("name", "=", "Madge").
		First().
		AsEntity(&emote)
	if err != nil {
		fmt.Println(err)
	}

	var emotes []Emote
	err = driver.Get().AsEntity(&emotes)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(emotes)
	fmt.Println(emote)
}
