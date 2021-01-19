package api

import (
	//"fmt"
	"fmt"
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
	Page     int `form:"page" json:"page"`           //当前页
	PageSize int `form:"page_size" json:"page_size"` // 每页大小
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
	UserID     string `form:"user_id" json:"user_id" binding:"required"`           //用户id
	UserTodoID string `form:"user_todo_id" json:"user_todo_id" binding:"required"` //用户Todo id
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

// GetUserTodo 根据UserID获得一条todo的路由处理函数
// @Tags UserTodo
// @Summary 查询一条数据
// @Param   user_id     query    string     true        "用户ID"
// @Param   user_todo_id     query    string     true        "用户Todo ID"
// @Success 200 {object} response.Result "{Code: 200, Msg: "请求成功", Data: "0"}"
// @Failure 500 {object} response.Result "{Code: 500, Msg: "服务器内部错误", Data: "服务器内部错误"}"
// @Router /api/v1/todo/user_id [get]
func (t *Todo) GetUserTodo(c *gin.Context) {
	var j findUserTodo
	if err := c.ShouldBind(&j); err != nil {

		response.ResError(c, err)
		return
	} //这里的if函数的目的是判断json里面是否有值？
	//userid,err:=strconv.Atoi(json.UserID)
	fmt.Print(j)
	userTodoID, err := t.TodoService.GetUserTodo(j.UserID, j.UserTodoID)
	if err != nil {
		response.ResError(c, err)
	}
	response.ResSuccess(c, userTodoID)
}

// GetUserAllTodos 分页获取所有数据的路由处理函数
// @Tags UserTodo
// @Summary 查询所有数据
// @Param   page     query    int     true        "当前页码"
// @Param   page_size     query    int     true        "每一页的显示数量"
// @Success 200 {object} response.Result "{Code: 200, Msg: "请求成功", Data: "0"}"
// @Failure 500 {object} response.Result "{Code: 500, Msg: "服务器内部错误", Data: "服务器内部错误"}"
// @Router /api/v1/todo [get]
func (t *Todo) GetUserAllTodos(c *gin.Context) {
	var formdata paginationInput
	// page, _ := strconv.Atoi(c.PostForm("page"))
	// pagesize, _ := strconv.Atoi(c.PostForm("page_size"))
	if err := c.ShouldBind(&formdata); err != nil {
		response.ResError(c, err)
		return
	}
	resultTodo, err := t.TodoService.GetUserAllTodos(formdata.Page, formdata.PageSize)
	if err != nil {
		response.ResError(c, err)
	}
	response.ResSuccess(c, resultTodo)
}

// UpdateUserTodo 是更新单条数据的路由处理函数
// @Tags UserTodo
// @Summary 更新一条UserTodo
// @Param addUserTodoInput body addUserTodoInput true "更新todo"
// @Success 200 {object} response.Result "{Code: 200, Msg: "请求成功", Data: "0"}"
// @Failure 500 {object} response.Result "{Code: 500, Msg: "服务器内部错误", Data: "服务器内部错误"}"
// @Router /api/v1/todo [put]
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
		json.UserTodoDueTime, json.UserTodoRemindTime, status)
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
