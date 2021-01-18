package app

import (
	"github.com/bo-er/todo-simpler/todo"
	"github.com/bo-er/todo-simpler/todo/inner"
)

var _TodoService todo.TodoService

// MustGetTodoService 获取todo 服务
func MustGetTodoService() todo.TodoService {
	if _TodoService != nil {
		return _TodoService
	}
	TodoService, err := inner.NewTodoService(DB)
	if err != nil {
		panic(err)
	}
	_TodoService = TodoService
	return _TodoService
}
