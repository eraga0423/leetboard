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
	title     string
	content   string
	nick      string
	postImage string
	authorID  string
}

func (p PostsGovernor) NewPost(_ context.Context, request controller.NewPostReq) (controller.NewPostResp, error) {
	postImage := request.GetImage()
	authID := request.GetAuthorID()
	typeJPFPNG, err := p.checkImageType(postImage)
	if err != nil {
		log.Print("dir: ", "governor", "method: ", "checkImageType", err.Error())
		return nil, err
	}
	postImageURL, err := p.miniostor.UploadImage(authID, authID, typeJPFPNG, postImage)
	if err != nil {
		log.Print("dir: ", "governor", "method: ", "minioUploadImage", err.Error())
		return nil, err
	}
	newResp := new(resp)
	newResp = &resp{
		title:     request.GetTitle(),
		content:   request.GetPostContent(),
		nick:      request.GetName(),
		authorID:  authID,
		postImage: postImageURL,
	}

	_, err = p.db.CreatePost(newResp)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *req) GetTitle() string {
	return r.title
}

func (r *req) GetPostContent() string {
	return r.content
}

func (r *req) GetImage() []byte {
	return r.postImage
}

func (r *req) GetName() string {
	return r.nick
}

func (r *req) GetAuthor() string { return r.authorID }
