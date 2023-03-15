package main

import (
	"fmt"

	"github.com/rprtr258/simpdb"
)

type Emote struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

func (e Emote) ID() string {
	return e.Name
}

func (Emote) TableName() string {
	return "emotes"
}

func main() {
	db := simpdb.New("db")
	emotes := simpdb.GetTable[Emote](db)

	_, err := emotes.GetAll()
	if err != nil {
		fmt.Println(err)
	}

	err = emotes.Insert(Emote{
		Name: "Madge",
		Link: "https://cdn.7tv.app/emote/60a95f109d598ea72fad13bd/4x.webp",
	})

	if err != nil {
		fmt.Println(err)
	}
}
