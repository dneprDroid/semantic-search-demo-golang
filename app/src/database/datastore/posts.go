package datastore 

import (
	"database/sql"

	"github.com/lib/pq"
)

type _Posts struct {}

var Posts _Posts

func (_Posts) InsertNew(db *sql.DB, postText string) (int, error) {
	var postId int
	err := db.QueryRow( 
		`INSERT INTO posts (content) VALUES ($1) RETURNING id`, 
		postText,
	).Scan(&postId)
	if err != nil {
		return 0, err 
	}
	return postId, nil 
}

func (_Posts) Find(db *sql.DB, postIds ...int) ([]Post, error) {
	rows, err := db.Query(
		`SELECT * FROM posts WHERE id = ANY($1) ORDER BY id`,
		pq.Array(postIds),
	)
	if err != nil {
		return nil, err 
	}
	defer rows.Close()

	posts := make([]Post, 0)
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.Id, &post.Content)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil 
}