package posts 

import (
	"io/ioutil"
	"net/http"
	"database/sql"

	"common/log"

	"appserver/response"
	"appserver/env"
	"appserver/database/datastore"
	"appserver/entity"
)

func AddPost(env env.Env, w http.ResponseWriter, req *http.Request) {
	bodyBytes, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		response.Error(w, err)
		return
	}
	postText := string(bodyBytes)
	log.Printf("body: `%v`", postText)

	var postId int
	err = env.Db.InContextSync(func(db *sql.DB) error {
		postId, err = datastore.Posts.InsertNew(db, postText)
		if err != nil {
			return err 
		}
		return nil 
	})
	if err != nil {
		response.Error(w, err)
		return
	}
	go func () {
		if err := processEmbeddings(env, postId, postText); err != nil {
			log.Printf("Embedding processing error: `%v`", err)
		}
	}()
	resp := entity.AddPostResponse{ Id: postId }
	response.WriteJson(w, resp)
}

func processEmbeddings(env env.Env, postId int, postText string) error {
	embeddings, err := requestEmbeddings(env, postText)
	if err != nil {
		return err 
	}
	return env.Db.InContextSync(func(db *sql.DB) error {
		return datastore.Embeddings.InsertPostData(db, postId, embeddings)
	})
}