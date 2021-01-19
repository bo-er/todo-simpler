package todo

// TodoService 是Todo服务的接口
type TodoService interface {
	AddUserTodo(userID, userTodoTitle, userTodoDescription, userTodoDueTime, userTodoRemindTime string, status int) (string, error)

	GetUserTodo(userID, userTodoID string) (string, error)

	GetUserAllTodos(page, pageSize int) ([]*ResultUserTodo, error)

	UpdateUserTodo(userID, userTodoTitle, userTodoDescription, userTodoDueTime, userTodoRemindTime string, status int) (string, error)

	DeleteUserTodo(userID, userTodoID string) (string, error)
}
