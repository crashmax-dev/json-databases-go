package main

import (
	"fmt"
	"jsondb-go/util"

	"github.com/rprtr258/simpdb"
)

type Emote struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

func (e Emote) ID() string {
	return e.Name
}

func main() {
	db := simpdb.New("db/simpdb")
	emotes, err := simpdb.GetTable[Emote](db, "emotes", simpdb.TableConfig{
		Indent: true,
	})
	if err != nil {
		fmt.Println(err)
	}
	defer emotes.Close()

	emote := Emote{
		Name: util.RandomString(10),
		Link: "https://cdn.7tv.app/emote/60a95f109d598ea72fad13bd/4x.webp",
	}

	emotes.Upsert(emote)
}
