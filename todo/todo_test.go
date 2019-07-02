package todo

import (
	"errors"
	"testing"

	"github.com/codystringham/mockingjay/db"
)

type MockRepo struct {
	getResult db.Todo
	getErr    error
}

func (m MockRepo) GetTodo(id int) (db.Todo, error) {
	return m.getResult, m.getErr
}

func TestIsFinished_True(t *testing.T) {
	uc := UseCase{
		Repo: MockRepo{
			getResult: db.Todo{
				Finished: true,
			},
			getErr: nil,
		},
	}
	IsFinished, _ := uc.IsFinished(1)
	if !IsFinished {
		t.Errorf("expected IsFinished to return true")
	}
}

func TestIsFinished_False(t *testing.T) {
	uc := UseCase{
		Repo: MockRepo{
			getResult: db.Todo{
				Finished: false,
			},
			getErr: nil,
		},
	}
	IsFinished, _ := uc.IsFinished(1)
	if IsFinished {
		t.Errorf("expected IsFinished to return false")
	}
}

func TestIsFinished_GetErr(t *testing.T) {
	uc := UseCase{
		Repo: MockRepo{
			getResult: db.Todo{},
			getErr:    errors.New("boom"),
		},
	}
	_, err := uc.IsFinished(1)
	if err.Error() != "boom" {
		t.Errorf("expected IsFinished to return an error, err: %v", err)
	}
}
