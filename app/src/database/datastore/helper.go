package datastore 

import (
	"fmt"
)

func buildBertQuery(embCount int, constCount int) string {
	sqlQuery := `select  emb1.postId from `
	for i := 0; i < embCount; i++ {
		sqlQuery += fmt.Sprintf(" embeddings emb%v,", i+1)
	}
	sqlQuery = sqlQuery[:len(sqlQuery)-1]
	sqlQuery += " where\n "

	sqlQuery += "("
	for i := 0; i < embCount; i++ {
		sqlQuery += fmt.Sprintf(" abs(emb%v.word - $%v) <= $1 ", i+1, i+1+constCount)
		if i < embCount - 1 {
			sqlQuery += "and"
		}
	}
	sqlQuery += ")"

	if embCount > 1 {
		sqlQuery += " and \n("
		for i := 0; i < embCount - 1; i++ {
			sqlQuery += fmt.Sprintf(" emb%v.postId = emb%v.postId  ", i+1, i+2)
			if i < embCount - 2 {
				sqlQuery += "and"
			}
		}
		sqlQuery += ")"
	}

	if embCount > 1 {
		sqlQuery += "and \n(("
		for i := 0; i < embCount - 1; i++ {
			sqlQuery += fmt.Sprintf(" abs(emb%v.word_offset - emb%v.word_offset)  ", i+2, i+1)
			if i < embCount - 2 {
				sqlQuery += "+"
			}
		}
		sqlQuery += fmt.Sprintf(")/%v) <= $2", embCount - 1)
	}
	return sqlQuery
}