package database 

import (
	"fmt"
	"database/sql"

	"common/log"

	"golang.org/x/exp/maps"
	_ "github.com/lib/pq"
)

type _Embeddings struct {}

var Embeddings _Embeddings 

func (_Embeddings) InsertPostData(db *sql.DB, postId int, embeddings []int) error {
	sqlStr := "INSERT INTO embeddings(postId, word_offset, word) VALUES "
	vals := make([]interface{}, len(embeddings) * 3)

	for offset, wordId := range embeddings {
		argNum := offset * 3 + 1
		sqlStr += fmt.Sprintf("($%v, $%v, $%v),", argNum, argNum+1, argNum+2)
		vals[offset * 3] = postId
		vals[offset * 3 + 1] = offset 
		vals[offset * 3 + 2] = wordId 
	}
	sqlStr = sqlStr[0:len(sqlStr)-1]

	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		log.Printf("sqlStr error: `%v`: %v", sqlStr, err)
		return err
	}
	_, err = stmt.Exec(vals...)
	if err != nil {
		return err
	}
	return nil 
}