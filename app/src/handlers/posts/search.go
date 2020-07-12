package posts 

import (
	"fmt"
	"net/http"
	"database/sql"

	"common/rpc"

	"appserver/util"
	"appserver/env"
	"appserver/database"
	"appserver/entity"
)

func FindPost(env env.Env, w http.ResponseWriter, req *http.Request) {
	query, ok := req.URL.Query()["q"]
	if !ok || len(query) == 0 || len(query[0]) == 0 {
		util.RespErrorStr(w, fmt.Sprintf("Bad query: '%s'", query))
		return 
	}
	embeddings, err := requestEmbeddings(env, query[0])
	if err != nil {
		util.RespError(w, err)
		return 
	}
	var dbPosts []database.Post
	err = env.Db.InContextSync(func(dbConn *sql.DB) error {
		postIds, err := database.Embeddings.FindPostIds(dbConn, embeddings)
		if err != nil {
			return err
		}
		dbPosts, err = database.Posts.Find(dbConn, postIds...)
		return err  
	})
	if err != nil {
		util.RespError(w, err)
		return
	}
}