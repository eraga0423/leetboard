package posts_handler

import (
	"log"
	"log/slog"
	"net/http"
	"text/template"
	"time"

	"1337b0rd/internal/constants"
)

func (h *PostsHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	slog.Info("get posts")
	tmpl := template.Must(template.ParseFiles(constants.Catalog))
	var data AllPost
	ctx := r.Context()
	log.Print("method:", " getPosts", " dir:", " rest")
	resp, err := h.ctrl.ListPosts(ctx)
	if err != nil {
		h.HandleError(w, 400)
		log.Print("error:", err)
		return
	}
	for _, v := range resp.GetList() {
		data.Posts = append(data.Posts, PostResp{
			PostID:      v.GetPostID(),
			PostTitle:   v.GetTitle(),
			PostContent: v.GetPostContent(),
			PostImage:   v.GetPostImageURL(),
			PostTime:    v.GetPostTime(),
		})
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		h.HandleError(w, 400)
		log.Print("method", "getPosts", "error:", err)
		return
	}
}

type AllPost struct {
	Posts []PostResp
}
type PostResp struct {
	PostID      int
	PostTitle   string
	PostContent string
	PostImage   string
	PostTime    time.Time
}
