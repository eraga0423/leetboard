package rick_morty

type RestRickAndMorty interface {
	ParseDataJson() (RespDataJson, error)
}

type RespDataJson interface {
	RespParseDataJson() []DataJson
}

type DataJson interface {
	GetId() int
	GetName() string
	GetImage() string
}
