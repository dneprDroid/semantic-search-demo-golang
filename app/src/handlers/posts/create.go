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
	resp := entity.AddPostResponse{ Id: postId }
	util.RespWriteJson(w, resp)
}