package rpc

type Vector []int 

type TextRequest struct {
	Content string
}

type EmbeddingResponse struct {
	Vec Vector
}