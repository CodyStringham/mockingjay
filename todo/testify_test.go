package todo

import (
	"errors"
	"testing"

	"github.com/codystringham/mockingjay/db"
	"github.com/codystringham/mockingjay/todo/mocks"
)

func TestIsFinished_MockTrue(t *testing.T) {
	mock := new(mocks.UseCaseRepo)
	uc := UseCase{
		Repo: mock,
	}

	mock.On("GetTodo", 2).Return(db.Todo{
		Finished: true,
	}, nil)

	IsFinished, _ := uc.IsFinished(1)
	if !IsFinished {
		t.Errorf("expected IsFinished to return true")
	}
	mock.AssertExpectations(t)
}

func TestIsFinished_MockFalse(t *testing.T) {
	mock := new(mocks.UseCaseRepo)
	uc := UseCase{
		Repo: mock,
	}

	mock.On("GetTodo", 1).Return(db.Todo{
		Finished: false,
	}, nil)

	IsFinished, _ := uc.IsFinished(1)
	if IsFinished {
		t.Errorf("expected IsFinished to return false")
	}
	mock.AssertExpectations(t)
}

func TestIsFinished_MockGetErr(t *testing.T) {
	mock := new(mocks.UseCaseRepo)
	uc := UseCase{
		Repo: mock,
	}

	mock.On("GetTodo", 1).Return(db.Todo{}, errors.New("boom"))

	_, err := uc.IsFinished(1)

	if err.Error() != "boom" {
		t.Errorf("expected IsFinished to return an error, err: %v", err)
	}
	mock.AssertExpectations(t)
}
