package database

type ResponseCharacters interface {
	GetCharacters() []GetCharacter
}

type GetCharacter interface {
	GetCharacterId() int
	GetCharacterName() string
	GetCharacterImage() string
	GetCharacterStatus() bool
}
type RequestCharacters interface {
	SetCharacters() []SetCharacter
}

type SetCharacter interface {
	SetCharacterId() int
	SetCharacterStatus() bool
}

type InsertCharacters interface {
	Insert() []InsertCharacter
}

type InsertCharacter interface {
	InsCharacterId() int
	InsCharacterName() string
	InsCharacterImage() string
	InsCharacterStatus() bool
}
