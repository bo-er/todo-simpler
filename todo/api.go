package todo

// TodoService 是Todo服务的接口
type TodoService interface {
	// @title    GetDeviceIDRawValue
	// @description   添加用户todo
	// @param   userID               		string  "用户ID"
	// @param   userTodoTitle               string  "用户Todo标题"
	// @param   userTodoDescription         string  "用户Todo描述"
	// @param   userTodoDueTime             string  "用户Todo截止时间"
	// @param   userTodoRemindTime          string  "用户Todo开始提醒时间"
	// @param   status             			string  "用户Todo状态"
	// @return  usertodoId              	string  "用户TodoId"
	// @error	ErrSystem
	AddUserTodo(userID, userTodoTitle, userTodoDescription, userTodoDueTime, userTodoRemindTime string, status int) (string, error)

	GetUserTodo(userID, userTodoID string) (string, error)

	GetUserAllTodos(page, pageSize int) ([]*ResultUserTodo, error)

	UpdateUserTodo(userID, userTodoTitle, userTodoDescription, userTodoDueTime, userTodoRemindTime string, status int) (string, error)
	// @title   DeleteDeviceIDRawValue
	// @description   删除用户Todo
	// @param   userID                       string "用户ID"
	// @return  userTodoID                   string "用户TodoID"
	// error    ErrSystem
	DeleteUserTodo(userID, userTodoID string) (string, error)
}
