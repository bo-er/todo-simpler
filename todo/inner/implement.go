package inner

import (
	"gorm.io/gorm"

	"github.com/bo-er/todo-simpler/todo"
	
	"github.com/bo-er/todo-simpler/todo/errors"
)

// DefaultTodoService todo service的默认实现
type DefaultTodoService struct {
	db *gorm.DB
}

// NewTodoService 产生一个新的TodoService
func NewTodoService(db *gorm.DB) (service todo.TodoService, err error) {
	if db == nil {
		err = errors.ErrDbIsNil
		return
	}
	service = &DefaultTodoService{
		db: db,
	}
	return service, nil
}

// AddUserTodo 默认todo service 实现增加用户todo接口
func (dts *DefaultTodoService) AddUserTodo(userID, userTodoTitle, userTodoDescription, userTodoDueTime, userTodoRemindTime string, status int) (string, error) {
	return AddUserTodo(dts.db, userID, userTodoTitle, userTodoDescription, userTodoDueTime, userTodoRemindTime, status)
}
