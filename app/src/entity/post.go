package entity

type AddPostResponse struct {
	Id int `json:"id"`
}

type Post struct {
	Id int 			`json:"id"`
	Content string  `json:"content"`
}