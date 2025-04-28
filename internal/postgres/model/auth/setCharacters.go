package auth

import (
	"fmt"

	"1337b0rd/internal/types/database"
)

func (a *Auth) UpdateCharacters(d database.RequestCharacters) error {
	charS := d.SetCharacters()
	var s string
	for i, char := range charS {
		if i == 0 {
			s = fmt.Sprintf(`(%d,%v)`, char.SetCharacterId(), char.SetCharacterStatus())
		} else {
			s += ","
			s1 := fmt.Sprintf(`(%d,%v)`, char.SetCharacterId(), char.SetCharacterStatus())
			s += s1
		}
	}
	query := fmt.Sprintf(`
	UPDATE avatars
	SET status = data.status
	FROM (VALUES %s) AS data(id, status)
	WHERE avatars.character_id = data.id
	`, s)
	_, err := a.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
