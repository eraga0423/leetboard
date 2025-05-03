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

type newPostResp struct {
	newName string
}

func (p *PostsGovernor) NewPost(ctx context.Context, request controller.NewPostReq) (controller.NewPostResp, error) {
	imageURL := ""
	log.Println("post: new post", "dir: governor")
	name := request.GetFormName()
	newObjectName := ""
	newName := newPostResp{}
	if name == "" {
		name = request.GetDefaultName()
	} else if name != "" {
		newName.newName = name
	}
	size := request.GetImage().GetObjectSize()
	idSession := request.GetAuthorIDSession()
	newReqStorage := reqStorage{
		bucketName:  idSession,
		objectSize:  size,
		contentType: request.GetImage().GetContentType(),
		metaData:    request.GetImage().GetFileIO(),
	}
	if size != 0 {

		temp, err := p.miniostor.ParseURL(ctx, &newReqStorage)
		if err != nil {
			return nil, err
		}
		imageURL = temp.GetImageURL()
		newObjectName = temp.GetNewObjectName()
	}
	newResp := &resp{
		title:           request.GetTitle(),
		content:         request.GetPostContent(),
		nick:            name,
		authorSessionID: idSession,
		avatarImage:     request.GetAvatarImageURL(),
		postImage:       imageURL,
	}

	resp, err := p.db.CreatePost(ctx, newResp)
	if err != nil {
		log.Print("dir: ", "governor", "method: ", "db.CreatePost", "  ERROR:  ", err.Error())
		return nil, err
	}
	if size != 0 {
		newReqStorage := reqStorage{
			bucketName:  idSession,
			objectSize:  size,
			objectName:  newObjectName,
			contentType: request.GetImage().GetContentType(),
			metaData:    request.GetImage().GetFileIO(),
		}
		if p.miniostor == nil {
			return nil, fmt.Errorf("minio storage is nil")
		}

		err = p.miniostor.UploadImage(ctx, &newReqStorage)
		if err != nil {
			log.Print("dir: ", "governor", "method: ", "minioUploadImage", "error", err.Error())
			err = resp.TxRollback(true)
			if err != nil {
				return nil, err
			}
			return nil, err
		}
	}
	err = resp.TxRollback(false)
	if err != nil {
		return nil, err
	}

	return &newName, nil
}

func (r *newPostResp) GetNewName() string {
	return r.newName
}
