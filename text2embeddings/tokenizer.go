package main 

import (
	"common/rpc"

	"github.com/sugarme/tokenizer"
	"github.com/sugarme/tokenizer/pretrained"
)

type Tokenizer struct {
	tk *tokenizer.Tokenizer
}

func NewTokenizer() *Tokenizer {
	return &Tokenizer{
		tk: pretrained.BertBaseUncased(),
	}
}

func (self *Tokenizer) Encode(sentence string) (rpc.Vector, error) {
	en, err := self.tk.EncodeSingle(sentence)
	if err != nil {
		return nil, err 
	}
	return en.Ids, nil 
}