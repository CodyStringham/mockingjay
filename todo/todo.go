package todo

import (
	"github.com/codystringham/mockingjay/db"
)

type UseCaseRepo interface {
	GetTodo(int) (db.Todo, error)
}

type UseCase struct {
	Repo UseCaseRepo
}

func (uc UseCase) IsFinished(id int) (bool, error) {
	todo, err := uc.Repo.GetTodo(id)
	if err != nil {
		return false, err
	}

	return todo.Finished, nil
}
