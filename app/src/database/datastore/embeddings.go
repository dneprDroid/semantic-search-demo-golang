package datastore 

import (
	"fmt"
	"database/sql"

	"common/log"

	_ "github.com/lib/pq"
)

type _Embeddings struct {}

var Embeddings _Embeddings 

func (_Embeddings) FindPostIds(db *sql.DB, embeddings []int) ([]int, error) {
	const (
		wordSim = 250
		contextDistance = 4
	)
	postIds := make(map[int]int, 0)
	
	queryArgs := make([]interface{}, 0)
	queryArgs = append(queryArgs, wordSim)
	if len(embeddings) > 1 {
		queryArgs = append(queryArgs, contextDistance)
	}
	for _, emb := range embeddings {
		queryArgs = append(queryArgs, emb)
	}
	rows, err := db.Query(
		buildBertQuery(
			len(embeddings), 
			len(queryArgs) - len(embeddings),
		),
		queryArgs...,
	)
	if err != nil {
		return nil, err 
	}
	defer rows.Close()
	for rows.Next() {
		var (
			postId int
		)
		err := rows.Scan(&postId)
		if err != nil {
			return nil, err
		}
		postIds[postId] += 1
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	result := make([]int, 0)
	for id, _ := range postIds {
		result = append(result, id)
	}
	return result, nil 
}

func (_Embeddings) InsertPostData(db *sql.DB, postId int, embeddings []int) error {
	sqlStr := "INSERT INTO embeddings(postId, word_offset, word) VALUES "
	vals := make([]interface{}, len(embeddings) * 3)

	for offset, wordId := range embeddings {
		argNum := offset * 3
		sqlStr += fmt.Sprintf("($%v, $%v, $%v),", argNum + 1, argNum + 2, argNum + 3)
		vals[argNum] = postId
		vals[argNum + 1] = offset 
		vals[argNum + 2] = wordId 
	}
	sqlStr = sqlStr[:len(sqlStr)-1]

	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		log.Printf("sqlStr error: `%v`: %v", sqlStr, err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(vals...)
	if err != nil {
		return err
	}
	return nil 
}