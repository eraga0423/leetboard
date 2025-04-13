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
			s = fmt.Sprintf(`VALUES('%d','%v')`, char.SetCharacterId(), char.SetCharacterStatus())
		} else {
			s += ","
			s1 := fmt.Sprintf(`('%d','%v')`, char.SetCharacterId(), char.SetCharacterStatus())
			s += s1
		}
	}

	_, err := a.db.Exec(`
	UPDATE avatars
	SET status = data.status
	FROM ($1) AS data(id, status)
	WHERE character_id = data.id`, s)
	if err != nil {
		return err
	}

	return nil
}
