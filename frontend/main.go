package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"text/template"
)

var comments []Comment

type newTempl struct {
	TitlePost string
	Posts     []OnePost
}
type OnePost struct {
	Title    string
	PostID   string
	Comments []Comment
}
type AllP struct {
	P []OnePost
}

var Posts AllP

func AllPosts() newTempl {
	Posts.P = append(Posts.P, OnePost{
		Title:  "1post",
		PostID: "1",
	}, OnePost{
		Title:  "2post",
		PostID: "2",
	}, OnePost{
		Title:  "3post",
		PostID: "3",
	})

	data := newTempl{
		TitlePost: "catalog",
		Posts:     Posts.P,
	}
	return data
}

func catalog(w http.ResponseWriter, r *http.Request) {
	data := AllPosts()
	tmpl := template.Must(template.ParseFiles("temp/catalog.html"))

	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/post/{id}", post)
	http.HandleFunc("/catalog", catalog)
	http.HandleFunc("/archive", archive)
	http.HandleFunc("/create", create)
	http.HandleFunc("/post/{id}/comment", NewComment)
	fmt.Println("start server")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("aaa")
		return
	}
}

func archive(w http.ResponseWriter, r *http.Request) {
	// tmpl := template.Must(template.ParseFiles("temp/archive.html"))
	// data := newTempl{
	// 	Title: "archive",
	// 	Posts: []string{"post4", "post5", "post6"},
	// }
	// tmpl.Execute(w, data)
}

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var NewReq Req
		name := r.FormValue("name")
		title := r.FormValue("title")
		postText := r.FormValue("post")
		file, _, err := r.FormFile("image")
		if err != nil && err != http.ErrMissingFile {
			http.Error(w, "Не удалось добавить файл", http.StatusBadRequest)
			return
		}
		var newFile []byte
		if file != nil {
			newFile, err = io.ReadAll(file)
			if err != nil {
				return
			}
			defer file.Close()
		}

		NewReq = Req{
			Title:    title,
			FileByte: newFile,
			Post:     postText,
			Nick:     name,
		}

		fmt.Println("asdasd", NewReq)

	}
	tmpl := template.Must(template.ParseFiles("temp/create_post.html"))
	data := newTempl{
		TitlePost: "new post",
	}
	tmpl.Execute(w, data)
}

type Req struct {
	Title    string `json:"title"`
	Post     string `json:"post"`
	Nick     string `json:"name"`
	FileByte []byte
}
type Resp struct {
	IDPost    string
	TitlePost string
	Title     string `json:"title"`
	Post      string `json:"post"`
	Nick      string `json:"name"`
	ImageURL  string
	Comments  []Comment
}
type Comment struct {
	Nick      string
	UserID    string
	Content   string
	CommentID string
}

func post(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("temp/post.html"))
	data := AllPosts()
	id := r.PathValue("id")
	NewPost := OnePost{}
	var numberId int
	for i, v := range data.Posts {
		if v.PostID == id {
			NewPost = v
			numberId = i
		}
	}

	RespPost := Resp{
		IDPost:    NewPost.PostID,
		TitlePost: "post",
		Title:     NewPost.Title,
		Post:      "post body",
		Nick:      "eraga",
		ImageURL:  "https://docs.aws.amazon.com/images/AWSEC2/latest/UserGuide/images/launch-from-ami.png",
		Comments:  Posts.P[numberId].Comments,
	}

	err := tmpl.Execute(w, RespPost)
	if err != nil {
		fmt.Println("asdasd", err)
		return
	}
}

func NewComment(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	CommentContent := r.FormValue("new_comment")
	a := Comment{
		Nick:    "zhantas",
		UserID:  "1",
		Content: CommentContent,
	}
	AddNewComment(a, id)
	fmt.Println("signal newcomment")
	http.Redirect(w, r, fmt.Sprintf("/post/%s", id), http.StatusSeeOther)
}

func AddNewComment(newComment Comment, id string) {
	idInt, _ := strconv.Atoi(id)
	Posts.P[idInt].Comments = append(Posts.P[idInt].Comments, Comment{
		Nick:    newComment.Nick,
		UserID:  newComment.UserID,
		Content: newComment.Content,
	})
}
