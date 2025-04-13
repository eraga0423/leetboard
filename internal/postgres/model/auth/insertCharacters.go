package auth

import (
	"fmt"

	"1337b0rd/internal/types/database"
)

func (a *Auth) InserCartoonCharacters(d database.InsertCharacters) error {
	charS := d.Insert()
	var s string
	for i, char := range charS {
		id := char.InsCharacterId()
		name := char.InsCharacterName()
		image := char.InsCharacterImage()
		status := char.InsCharacterStatus()
		if i == 0 {
			s = fmt.Sprintf(`VALUES('%d', '%s', '%s' ,'%v')`, id, name, image, status)
		} else {
			s += ","
			s1 := fmt.Sprintf(`('%d', '%s', '%s' ,'%v')`, id, name, image, status)
			s += s1
		}
	}

	_, err := a.db.Exec(fmt.Sprintf(`
	INSERT INTO avatars
	(character_id, character_name, character_image, status)
	%s`, s))
	if err != nil {
		return err
	}

	return nil
}
