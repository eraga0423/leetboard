package posts_governor

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"

	"1337b0rd/internal/types/controller"
)

type req struct {
	title    string
	content  string
	nick     string
	authorID string
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

type reqStorage struct {
	bucketName  string
	objectName  string
	objectSize  int64
	contentType string
	metaData    multipart.File
}

func (r *resp) GetTitle() string {
	return r.title
}

func (r *resp) GetPostContent() string {
	return r.content
}

func (r *resp) GetPostImageURL() string {
	return r.postImage
}

func (r *resp) GetAuthorAvatarURL() string {
	return r.avatarImage
}

func (r *resp) GetAuthorName() string {
	return r.nick
}

func (r *resp) GetAuthorSession() (idSessionUser string) {
	return r.authorSessionID
}

func (s *reqStorage) GetBucketName() string {
	return s.bucketName
}

func (s *reqStorage) GetObjectName() string {
	return s.objectName
}

func (s *reqStorage) GetObjectSize() int64 {
	return s.objectSize
}

func (s *reqStorage) GetContentType() string {
	return s.contentType
}

func (s *reqStorage) GetMetaData() multipart.File {
	return s.metaData
}

func (p *PostsGovernor) NewPost(ctx context.Context, request controller.NewPostReq) (controller.NewPostResp, error) {
	imageURL := ""
	log.Println("post: new post", "dir: governor")
	name := request.GetFormName()
	if name == "" {
		name = request.GetDefaultName()
	}
	size := request.GetImage().GetObjectSize()
	idSession := request.GetAuthorIDSession()
	if size != 0 {
		newReqStorage := reqStorage{
			bucketName:  idSession,
			objectName:  idSession,
			objectSize:  size,
			contentType: request.GetImage().GetContentType(),
			metaData:    request.GetImage().GetFileIO(),
		}
		if p.miniostor == nil {
			return resp{}, fmt.Errorf("minio storage is nil")
		}
		postImageURL, err := p.miniostor.UploadImage(ctx, &newReqStorage)
		if err != nil {
			log.Print("dir: ", "governor", "method: ", "minioUploadImage", err.Error())
			return nil, err
		}
		imageURL = postImageURL.GetImageURL()
	}

	newResp := &resp{
		title:   request.GetTitle(),
		content: request.GetPostContent(),
		nick:    name,

		authorSessionID: idSession,
		avatarImage:     request.GetAvatarImageURL(),
		postImage:       imageURL,
	}

	log.Println(newResp) //////////////////////////////////////////

	_, err := p.db.CreatePost(newResp)
	if err != nil {
		log.Print("dir: ", "governor", "method: ", "db.CreatePost", "  ERROR:  ", err.Error())
		return nil, err
	}
	return nil, nil
}
