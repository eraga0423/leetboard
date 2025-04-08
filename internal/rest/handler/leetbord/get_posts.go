package posts_handler

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"1337b0rd/internal/constants"
)

func (h *PostsHandler) GetPosts(w http.ResponseWriter, r *http.Request) { //
	var data AllPost
	ctx := r.Context()
	fmt.Println("This GET /posts")
	resp, err := h.ctrl.ListPosts(ctx)
	if err != nil {
		fmt.Println("rest : method List posts")
		http.Error(w, err.Error(), 499) /////////////////////////////////////
		return
	}
	if resp == nil {
		///////////////////////
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

	tmpl := template.Must(template.ParseFiles(constants.Catalog))

	err = tmpl.Execute(w, data)
	if err != nil {

		fmt.Println(err, "rest : method List posts!!!!!!!!!!!")
		http.Error(w, "dont list post", 499) /////////////////////////
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
