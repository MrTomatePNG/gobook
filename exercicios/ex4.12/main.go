package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type WebComicXKCD struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

type ComicID int

var database = make(map[ComicID]WebComicXKCD)

func init() {
	if err := load(); err != nil {
		fmt.Println("Erro load database: ", err.Error())
	}
}

func main() {
	var comic WebComicXKCD
	resp, err := http.Get("https://xkcd.com/571/info.0.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error request : %s", err)
	}

	err = json.NewDecoder(resp.Body).Decode(&comic)

	if err != nil {
		fmt.Println("Deu erro menro", err)
	}
	fmt.Printf("%#v", comic)

	database[ComicID(comic.Num)] = comic

}

func load() error {
	file, err := os.Open("db.json")

	if err != nil {
		return err
	}

	defer file.Close()

	decode := json.NewDecoder(file)
	return decode.Decode(&database)
}

func save() error {
	file, err := os.Create("db.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(&database)
}
