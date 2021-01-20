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
	var _status bool
	if status == 1 {
		_status = true
	} else {
		_status = false
	}
	userTodo := &todo.UserTodo{
		UserID:              userID,
		UserTodoID:          userTodoID,
		UserTodoTitle:       userTodoTitle,
		UserTodoDescription: userTodoDescription,
		UserTodoDueTime:     convertedTodoTime,
		UserTodoRemindTime:  convertedRemindTime,
		Status:              _status,
	}

	err := db.Create(userTodo)
	return userTodoID, err.Error
}

//GetUserTodo  获得一条todo数据
func GetUserTodo(db *gorm.DB,
	userID,
	userTodoID string,
) (string, error) {

	var resulttodo todo.UserTodo
	//fmt.Print(userID, userTodoID)
	err := db.Where("user_id = ? and user_todo_id=?", userID, userTodoID).First(&resulttodo)
	//return "resulttodo.UserTodoID", err.Error
	return resulttodo.UserTodoID, err.Error
}

//GetUserAllTodos 获取所有todo数据
func GetUserAllTodos(db *gorm.DB,
	page,
	pageSize int,
) ([]*todo.ResultUserTodo, error) {
	var resulttodos []*todo.ResultUserTodo
	//var resulttodo	todo.ResultUserTodo
	var usertodo []todo.UserTodo
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
	err := db.Offset(offset).Limit(pageSize).Find(&usertodo)
	for _, item := range usertodo {
		resulttodos = append(resulttodos, &todo.ResultUserTodo{UserTodoID: item.UserTodoID, UserTodoTitle: item.UserTodoTitle,
			UserTodoDescription: item.UserTodoDescription,
		})
	}

	return resulttodos, err.Error
}

//UpdateUserTodo 更新某一条todo数据
func UpdateUserTodo(db *gorm.DB,
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
	var _status bool
	if status == 1 {
		_status = true
	} else {
		_status = false
	}
	var userTodo todo.UserTodo

	err := db.Model(&userTodo).Where("user_id = ?", userID).Updates(map[string]interface{}{"UserTodoID": userTodoID, "UserTodoTitle": userTodoTitle,
		"UserTodoDescription": userTodoDescription, "UserTodoDueTime": convertedTodoTime, "UserTodoRemindTime": convertedRemindTime,
		"Status": _status})
	return userTodoID, err.Error
}

//DeleteUserTodo 删除某一条UserTodo数据
func DeleteUserTodo(db *gorm.DB, userID, userTodoID string) (string, error) {
	var resulttodo todo.UserTodo
	err := db.Where("user_todo_id = ?", userTodoID).Delete(&resulttodo)
	return resulttodo.UserTodoID, err.Error
}
