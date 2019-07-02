package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Todo struct {
	gorm.Model
	Name     string
	Finished bool
}

type Repo struct {
	DB *gorm.DB
}

func NewRepo() Repo {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Todo{})

	return Repo{
		DB: db,
	}
}

func (r Repo) CreateTodo(name string) {
	r.DB.Create(&Todo{Name: name})
}

func (r Repo) GetTodo(id int) (Todo, error) {
	var todo Todo
	r.DB.First(&todo, id)
	return todo, nil
}
