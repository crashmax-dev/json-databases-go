package main

import (
	"fmt"
	"jsondb-go/util"
	"os"
	"path/filepath"

	"github.com/FantasyGao/jsondb"
)

type Emote struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

func init() {
	dir := filepath.Join("db", "jsondb")

	if _, err := os.Stat(dir); err != nil {
		if !os.IsNotExist(err) {
			fmt.Printf("failed checking directory %s: %v", dir, err)
		}

		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Printf("failed creating directory %s: %v", dir, err)
		}
	}
}

func main() {
	db := jsondb.Create("db/jsondb/emotes.json")
	emoteKey := util.RandomString(10)
	db.Write(emoteKey, Emote{
		Name: "Kappa",
		Link: "https://static-cdn.jtvnw.net/emoticons/v2/25/default/dark/2.0",
	}).Save()

	emote := db.Read(emoteKey).(Emote)
	fmt.Println(emote.Name, emote.Link)
}
