package api

import (
	"github.com/bo-er/todo-simpler/admin/response"
	"github.com/bo-er/todo-simpler/todo"
	"github.com/gin-gonic/gin"
)

// Todo 定义Todo APi
type Todo struct {
	TodoService todo.TodoService
}

type paginationInput struct {
	Page     int `form:"page"`      //当前页
	PageSize int `form:"page_size"` // 每页大小
}

type addUserTodoInput struct {
	UserTodoTitle       string `json:"user_todo_title" binding:"required"`                  //用户Todo的标题
	UserTodoDescription string `json:"user_todo_description"`                               //用户Todo的描述
	UserTodoDueTime     string `json:"user_todo_due_time" example:"2021-01-22 16:09:00"`    //用户Todo截止时间
	UserTodoRemindTime  string `json:"user_todo_remind_time" example:"2021-01-21 16:09:00"` //用户Todo提前多久通知
	Status              bool   //用户Todo状态
}

// AddUserTodo 是添加新的UserTodo的路由处理函数
func (t *Todo) AddUserTodo(c *gin.Context) {
	var json addUserTodoInput
	if err := c.ShouldBindJSON(&json); err != nil {
		response.ResError(c, err)
		return
	}

	userID := "test"
	userTodoID, err := t.TodoService.AddUserTodo(userID, json.UserTodoTitle, json.UserTodoDescription, json.UserTodoDueTime, json.UserTodoRemindTime, 2)
	if err != nil {
		response.ResError(c, err)
	}
	response.ResSuccess(c, userTodoID)
}


