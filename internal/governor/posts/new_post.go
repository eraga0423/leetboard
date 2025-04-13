package posts_governor

import (
	"1337b0rd/internal/types/controller"
	"context"
	"log"
)

type req struct {
	title     string
	content   string
	nick      string
	postImage []byte
	authorID  string
}

type resp struct {
	title           string
	content         string
	nick            string
	postImage       string
	avatarImage     string
	authorSessionID string
	status          string
}

func (r *resp) GetTitle() string {
	return r.title
}

func (r *resp) GetPostContent() string {
	return r.content
}

func (r *resp) GetImage() string {
	return r.postImage
}

func (r *resp) GetAuthorSession() (idSessionUser string) {
	return r.authorSessionID
}

func (p *PostsGovernor) NewPost(_ context.Context, request controller.NewPostReq) (controller.NewPostResp, error) {
	postImage := request.GetImage()
	authID := request.GetAuthorID()
	typeJPGPNG, err := p.checkImageType(postImage)
	if err != nil {
		log.Print("dir: ", "governor", "method: ", "checkImageType", err.Error())
		return nil, err
	}
	newID, err := p.interceptor.GenerateSessionID()
	if err != nil {
		log.Print("dir: ", "governor", "method", "GenerateSessionID", err.Error())
		return nil, err
	}
	postImageURL, err := p.miniostor.UploadImage(newID, authID, typeJPGPNG, postImage)
	if err != nil {
		log.Print("dir: ", "governor", "method: ", "minioUploadImage", err.Error())
		return nil, err
	}
	newResp := new(resp)
	newResp = &resp{
		title:           request.GetTitle(),
		content:         request.GetPostContent(),
		nick:            request.GetName(),
		authorSessionID: authID,
		postImage:       postImageURL,
	}

	_, err = p.db.CreatePost(newResp)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
