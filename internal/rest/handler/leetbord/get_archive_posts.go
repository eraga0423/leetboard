package posts_handler

import (
	"log"
	"net/http"
	"text/template"
	"time"

	"1337b0rd/internal/constants"
)

type listArchivePostsResp struct {
	TitlePost    string
	ArchivePosts []OnePostArchive
}
type OnePostArchive struct {
	PostID       int
	Title        string
	PostContent  string
	PostImageURL string
	PostTime     time.Time
}

func (h *PostsHandler) GetArchive(w http.ResponseWriter, r *http.Request) {
	resp, err := h.ctrl.ListArchivePosts(r.Context())
	if err != nil {
		h.HandleError(w, http.StatusBadRequest)
		return
	}
	data := listArchivePostsResp{
		TitlePost: "ARCHIVE POSTS",
	}
	for _, postResp := range resp.GetList() {
		data.ArchivePosts = append(data.ArchivePosts, OnePostArchive{
			PostID:       postResp.GetPostID(),
			Title:        postResp.GetTitle(),
			PostContent:  postResp.GetPostContent(),
			PostImageURL: postResp.GetPostImageURL(),
			PostTime:     postResp.GetPostTime(),
		})
	}
	log.Print("This GET /archive")
	tmpl := template.Must(template.ParseFiles(constants.Archive))
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err)
		h.HandleError(w, http.StatusInternalServerError)
		return
	}

}
