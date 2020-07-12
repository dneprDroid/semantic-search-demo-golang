package posts 

import (
	"io/ioutil"
	"net/http"
	"database/sql"

	"common/log"

	"appserver/util"
	"appserver/env"
	"appserver/database"
	"appserver/entity"
)

func AddPost(env env.Env, w http.ResponseWriter, req *http.Request) {
	bodyBytes, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		util.RespError(w, err)
		return
	}
	postText := string(bodyBytes)
	log.Printf("body: `%v`", postText)

	var postId int
	err = env.Db.InContextSync(func(db *sql.DB) error {
		postId, err = database.Posts.InsertNew(db, postText)
		if err != nil {
			return err 
		}
		return nil 
	})
	if err != nil {
		util.RespError(w, err)
		return
	}
	go func () {
		if err := processEmbeddings(env, postId, postText); err != nil {
			log.Printf("Embedding processing error: `%v`", err)
		}
	}()
	resp := entity.AddPostResponse{ Id: postId }
	util.RespWriteJson(w, resp)
}

func processEmbeddings(env env.Env, postId int, postText string) error {
	embeddings, err := requestEmbeddings(env, postText)
	if err != nil {
		return err 
	}
	return env.Db.InContextSync(func(db *sql.DB) error {
		return database.Embeddings.InsertPostData(db, postId, embeddings)
	})
}