package leetboard

import "1337b0rd/internal/types/database"

func (l *Leetboard) CreatePost(req database.NewPostReq) (database.NewPostResp, error) {
	req.GetImage()
}
