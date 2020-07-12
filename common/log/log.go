package log

import (
	"fmt"
	golog "log"
)

func Printf(s string, args ...interface{}) {
	golog.Printf(s, args...)
}

func Println(args ...interface{}) string {
	msg := fmt.Sprintln(args...)
	golog.Println(msg)
	return msg
}

func Fatalln(args ...interface{}) {
	panic(Println(args...))
}