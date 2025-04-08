package posts_handler

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"1337b0rd/internal/constants"
)

func (h *PostsHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	var data allPost
	ctx := r.Context()
	fmt.Println("This GET /posts")
	resp, err := h.ctrl.ListPosts(ctx)
	if err != nil {
		fmt.Println("rest : method List posts")
		http.Error(w, "dont list post", 499) /////////////////////////////////////
		return
	}
	if resp==nil{
///////////////////////
	}

	

	for _, v := range resp.GetList() {
		data.posts = append(data.posts, postResp{
			PostID:      v.GetPostID(),
			PostTitle:   v.GetTitle(),
			PostContent: v.GetPostContent(),
			PostImage:   v.GetPostImageURL(),
			PostTime:    v.GetPostTime(),
		})
	}

	
	tmpl := template.Must(template.ParseFiles(constants.Catalog))

	tmpl.Execute(w, data)
}

type allPost struct {
	posts []postResp
}
type postResp struct {
	
	PostID      int
	PostTitle   string
	PostContent string
	PostImage   string
	PostTime    time.Time
}
