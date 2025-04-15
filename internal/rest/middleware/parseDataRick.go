package middleware

type avatar struct {
	name     string
	imageURL string
}

var mapAvatar = make(map[avatar]bool)

func (m *Middleware) parseJson() (avatar, error) {
	listAvatars := make([]avatar, 0)
	resp, err := m.mortyRick.ParseDataJson()
	if err != nil {
		return avatar{}, err
	}
	ListAvatars := resp.RespParseDataJson()
	for _, avatarOne := range ListAvatars {
		listAvatars = append(listAvatars, avatar{
			name:     avatarOne.GetName(),
			imageURL: avatarOne.GetImage(),
		})
		////////////////////////////////////
	}

	return parsOneAvatar(listAvatars), nil

}

func parsOneAvatar(list []avatar) avatar {
	for _, a := range list {
		if mapAvatar[a] {
			mapAvatar[a] = false
			return a
		}
	}
	//////////////////////////////////////////

	return avatar{}
}
