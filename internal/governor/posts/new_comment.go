package posts_governor

import (
	"context"
	"log/slog"
	"mime/multipart"
	"strconv"

	"1337b0rd/internal/types/controller"
)

type respDB struct {
	avatarImageURL  string
	avatarName      string
	parentCommentID int
	sessionID       string
	content         string
	imageCommentURL string
	postID          int
}
type metaDataComment struct {
	bucketName  string
	objectName  string
	fileIO      multipart.File
	objectSize  int64
	contentType string
}

func (g *PostsGovernor) NewComment(req controller.NewCommentReq, ctx context.Context) (controller.NewCommentResp, error) {
	fileSize := req.GetImageComment().GetObjectSize()
	sessionID := req.GetSessionID()
	commentImageURl := ""
	newObjectName := ""
	newStorage := metaDataComment{
		fileIO:      req.GetImageComment().GetFileIO(),
		objectSize:  fileSize,
		contentType: req.GetImageComment().GetContentType(),
		bucketName:  sessionID,
	}
	if fileSize != 0 {
		resp, err := g.miniostor.ParseURL(ctx, &newStorage)
		if err != nil {
			return nil, err
		}
		newObjectName = resp.GetNewObjectName()
		commentImageURl = resp.GetImageURL()
	}
	parentCommentID := req.GetParentCommentID()
	postIDInt, err := strconv.Atoi(req.GetPostID())
	if err != nil {
		return nil, err
	}
	parentCommentInt := 0
	if parentCommentID != "" {
		parentCommentInt, err = strconv.Atoi(req.GetParentCommentID())
		if err != nil {
			return nil, err
		}
	}

	newRespDB := &respDB{
		avatarName:      req.GetAvatarName(),
		avatarImageURL:  req.GetAvatarImageURL(),
		parentCommentID: parentCommentInt,
		sessionID:       sessionID,
		content:         req.GetContent(),
		imageCommentURL: commentImageURl,
		postID:          postIDInt,
	}
	resp, err := g.db.CreateComment(ctx, newRespDB)
	if err != nil {
		return nil, err
	}
	if fileSize != 0 {
		newStorage := metaDataComment{
			bucketName:  sessionID,
			fileIO:      req.GetImageComment().GetFileIO(),
			objectSize:  fileSize,
			contentType: req.GetImageComment().GetContentType(),
			objectName:  newObjectName,
		}
		err := g.miniostor.UploadImage(ctx, &newStorage)
		if err != nil {
			err = resp.TxRollback(true)
			if err != nil {
				return nil, err
			}
			slog.Error("error upload image")
			return nil, err
		}
	}
	err = resp.TxRollback(false)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *metaDataComment) GetBucketName() string       { return c.bucketName }
func (c *metaDataComment) GetObjectName() string       { return c.objectName }
func (c *metaDataComment) GetContentType() string      { return c.contentType }
func (c *metaDataComment) GetObjectSize() int64        { return c.objectSize }
func (c *metaDataComment) GetMetaData() multipart.File { return c.fileIO }

func (d *respDB) GetAuthorName() string      { return d.avatarName }
func (d *respDB) GetAuthorAvatarURL() string { return d.avatarImageURL }
func (d *respDB) GetPostID() int             { return d.postID }
func (d *respDB) GetParentCommentID() int    { return d.parentCommentID }
func (d *respDB) GetCommentContent() string  { return d.content }
func (d *respDB) GetCommentImage() string    { return d.imageCommentURL }
func (d *respDB) GetAuthorSession() string   { return d.sessionID }
