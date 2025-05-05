package rickmortyrest

import (
	"encoding/json"
	"log"
	"net/http"

	"1337b0rd/internal/types/rick_morty"
)

type parseJson struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	ID    int    `json:"id"`
}
type parseJsons struct {
	Posts []parseJson `json:"results"`
}

func (p *RickAndMorty) ParseDataJson() (rick_morty.RespDataJson, error) {
	resp, err := http.Get("https://rickandmortyapi.com/api/character")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	var data parseJsons
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &data, err
}

func (p *parseJsons) RespParseDataJson() []rick_morty.DataJson {
	newJsons := make([]rick_morty.DataJson, len(p.Posts))
	for i, post := range p.Posts {
		newJsons[i] = &post
	}
	return newJsons
}

func (p *parseJson) GetName() string {
	return p.Name
}

func (p *parseJson) GetImage() string {
	return p.Image
}

func (p *parseJson) GetId() int {
	return p.ID
}
