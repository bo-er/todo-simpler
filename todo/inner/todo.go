package inner

import (
	"time"

	"gorm.io/gorm"

	"github.com/bo-er/todo-simpler/todo"
	"github.com/bo-er/todo-simpler/utils/uuid"
)

// TimeTemplate 是时间模板
var TimeTemplate = "2006-01-02 15:04:05"

// Paginate 分页函数
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// AddUserTodo 是增加用户todo的数据层实现
func AddUserTodo(
	db *gorm.DB,
	userID,
	userTodoTitle,
	userTodoDescription,
	userTodoDueTime,
	userTodoRemindTime string,
	status int,
) (string, error) {
	userTodoID := uuid.MustString()
	convertedTodoTime, _ := time.ParseInLocation(TimeTemplate, userTodoDueTime, time.Local)
	convertedRemindTime, _ := time.ParseInLocation(TimeTemplate, userTodoRemindTime, time.Local)
	userTodo := &todo.UserTodo{
		UserID:              userID,
		UserTodoID:          userTodoID,
		UserTodoTitle:       userTodoTitle,
		UserTodoDescription: userTodoDescription,
		UserTodoDueTime:     convertedTodoTime,
		UserTodoRemindTime:  convertedRemindTime,
	}

	err := db.Create(userTodo)
	return userTodoID, err.Error
}
