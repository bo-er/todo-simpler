package api

import (
	"strconv"

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
	UserID              int    `json:"user_id"`                                             //用户id(新增)
	UserTodoTitle       string `json:"user_todo_title" binding:"required"`                  //用户Todo的标题
	UserTodoDescription string `json:"user_todo_description"`                               //用户Todo的描述
	UserTodoDueTime     string `json:"user_todo_due_time" example:"2021-01-22 16:09:00"`    //用户Todo截止时间
	UserTodoRemindTime  string `json:"user_todo_remind_time" example:"2021-01-21 16:09:00"` //用户Todo提前多久通知
	Status              bool   `json:"status"`                                              //用户Todo状态
}

type findUserTodo struct {
	UserID     string `json:"user_id"`      //用户id
	UserTodoID string `json:"user_todo_id"` //用户Todo id
}

// AddUserTodo 是添加新的UserTodo的路由处理函数
// @Tags UserTodo
// @Summary 添加新的UserTodo
// @Param body body addUserTodoInput true "创建新的todo"
// @Success 200 {object} response.Result "{Code: 200, Msg: "请求成功", Data: "0"}"
// @Failure 500 {object} response.Result "{Code: 500, Msg: "服务器内部错误", Data: "服务器内部错误"}"
// @Router /api/v1/todo [post]
func (t *Todo) AddUserTodo(c *gin.Context) {
	var json addUserTodoInput
	if err := c.ShouldBindJSON(&json); err != nil {
		response.ResError(c, err)
		return
	}
	var status int
	if json.Status {
		status = 1
	} else {
		status = 0
	}
	userID := strconv.Itoa(json.UserID)
	userTodoID, err := t.TodoService.AddUserTodo(userID, json.UserTodoTitle, json.UserTodoDescription,
		json.UserTodoDueTime, json.UserTodoRemindTime, status)
	if err != nil {
		response.ResError(c, err)
	}
	response.ResSuccess(c, userTodoID)
}

// GetUserTodo 获得一条todo的路由处理函数
func (t *Todo) GetUserTodo(c *gin.Context) {
	var json findUserTodo
	if err := c.ShouldBindJSON(&json); err != nil {
		response.ResError(c, err)
		return
	} //这里的if函数的目的是判断json里面是否有值？
	//userid,err:=strconv.Atoi(json.UserID)
	userTodoID, err := t.TodoService.GetUserTodo(json.UserID, json.UserTodoID)
	if err != nil {
		response.ResError(c, err)
	}
	response.ResSuccess(c, userTodoID)
}

// GetUserAllTodos 获得所有todo的路由处理函数
func (t *Todo) GetUserAllTodos(c *gin.Context) {
	// var page paginationInput
	// if err := c.ShouldBind(&page); err != nil {
	// 	response.ResError(c, err)
	// 	return
	// }
	page, _ := strconv.Atoi(c.PostForm("page"))
	pagesize, _ := strconv.Atoi(c.PostForm("page_size"))
	resultTodo, err := t.TodoService.GetUserAllTodos(page, pagesize)
	if err != nil {
		response.ResError(c, err)
	}
	response.ResSuccess(c, resultTodo)
}

//UpdateUserTodo 更新单个todo的路由处理函数
func (t *Todo) UpdateUserTodo(c *gin.Context) {
	var json addUserTodoInput
	if err := c.ShouldBindJSON(&json); err != nil {
		response.ResError(c, err)
		return
	}
	var status int
	if json.Status {
		status = 1
	} else {
		status = 0
	}
	userID := strconv.Itoa(json.UserID)
	userTodoID, err := t.TodoService.UpdateUserTodo(userID, json.UserTodoTitle, json.UserTodoDescription,
		json.UserTodoDueTime, json.UserTodoRemindTime, status) //这里的t相当于参数，直接可以运用？
	if err != nil {
		response.ResError(c, err)
	}
	response.ResSuccess(c, userTodoID)
}

//DeleteUserTodo 是删除单个UserTodo的路由处理函数
func (t *Todo) DeleteUserTodo(c *gin.Context) {
	var json findUserTodo
	if err := c.ShouldBindJSON(&json); err != nil {
		response.ResError(c, err)
		return
	}
	userTodoID, err := t.TodoService.DeleteUserTodo(json.UserID, json.UserTodoID)
	if err != nil {
		response.ResError(c, err)
	}
	response.ResSuccess(c, userTodoID)
}
