package rickmortyrest

import (
	"1337b0rd/internal/types/rick_morty"
	"encoding/json"
	"log"
	"net/http"
)

type parseJson struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	ID    string `json:"id"`
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

	return nil
}
func (p *parseJson) GetName() string {
	return p.Name
}
func (p *parseJson) GetImage() string {
	return p.Image
}
