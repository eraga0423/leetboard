package posts_handler

import (
	"log"
	"net/http"
	"text/template"
	"time"

	"1337b0rd/internal/constants"
)

type listArchivePostsResp struct {
	archivePosts []onePostArchive
}
type onePostArchive struct {
	postID       int
	title        string
	postContent  string
	postImageURL string
	postTime     time.Time
}

func (h *PostsHandler) GetArchive(w http.ResponseWriter, r *http.Request) {
	resp, err := h.ctrl.ListArchivePosts(r.Context())
	if err != nil {
		h.HandleError(w, http.StatusBadRequest)
		return
	}
	data := listArchivePostsResp{}
	for _, postResp := range resp.GetList() {
		data.archivePosts = append(data.archivePosts, onePostArchive{
			postID:       postResp.GetPostID(),
			title:        postResp.GetTitle(),
			postContent:  postResp.GetPostContent(),
			postImageURL: postResp.GetPostImageURL(),
			postTime:     postResp.GetPostTime(),
		})
	}
	log.Print("This GET /archive")
	tmpl := template.Must(template.ParseFiles(constants.Archive))
	err = tmpl.Execute(w, data)
	if err != nil {
		h.HandleError(w, http.StatusInternalServerError)
		return
	}

}
