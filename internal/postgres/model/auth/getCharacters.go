package auth

import (
	"errors"

	"1337b0rd/internal/types/database"
)

type characterResp struct {
	characterID     int
	characterName   string
	characterImage  string
	characterStatus bool
}

type characterSResp struct {
	characters []characterResp
}

func (a *Auth) ListCharacters() (database.ResponseCharacters, error) {
	sql, err := a.db.Query(`
	SELECT 
	character_id,
	character_name,
	character_image,
	status
	FROM avatars`)
	var charS []characterResp
	if err != nil {
		return characterSResp{}, err
	}
	defer sql.Close()
	for sql.Next() {
		var char characterResp
		err := sql.Scan(
			&char.characterID,
			&char.characterName,
			&char.characterImage,
			&char.characterStatus,
		)
		if err != nil {
			return characterSResp{}, err
		}
		charS = append(charS, char)
	}
	if len(charS) == 0 {
		return characterSResp{}, errors.New("characters are emppty")
	}
	return characterSResp{characters: charS}, nil
}

func (c characterSResp) GetCharacters() []database.GetCharacter {
	b := make([]database.GetCharacter, len(c.characters))
	for i, v := range c.characters {
		b[i] = v
	}
	return b
}

func (c characterResp) GetCharacterId() int       { return c.characterID }
func (c characterResp) GetCharacterName() string  { return c.characterName }
func (c characterResp) GetCharacterImage() string { return c.characterImage }
func (c characterResp) GetCharacterStatus() bool  { return c.characterStatus }
