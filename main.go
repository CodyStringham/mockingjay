package main

import (
	"fmt"

	"github.com/codystringham/mockingjay/db"
	"github.com/codystringham/mockingjay/todo"
)

func main() {
	repo := db.NewRepo()
	uc := todo.UseCase{
		Repo: repo,
	}
	IsFinished, _ := uc.IsFinished(1)
	fmt.Printf("Todo Finished: %v", IsFinished)
}
