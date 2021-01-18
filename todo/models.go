package todo

import (
	"time"

	"github.com/bo-er/todo-simpler/models"
)

// UserTodo 定义UserTodo表结构
type UserTodo struct {
	models.Base
	UserID              string    `gorm:"size:50;not null;"`  //用户ID
	UserTodoID          string    `gorm:"size:50;not null;"`  //用户Todo ID
	UserTodoTitle       string    `gorm:"size:50;not null;"`  //用户Todo的标题
	UserTodoDescription string    `gorm:"size:200;not null;"` //用户Todo的描述
	UserTodoDueTime     time.Time //用户Todo截止时间
	UserTodoRemindTime  time.Time //用户Todo提前多久通知
	Status              bool      //用户Todo状态
}

// ResultUserTodo 是UserTodo的数据库查询结构体
type ResultUserTodo struct {
	models.Base
	UserTodoID          string `json:"user_todo_id"`          //用户todo ID
	UserTodoTitle       string `json:"user_todo_title"`       //用户todo 标题
	UserTodoDescription string `json:"user_todo_description"` //用户todo 描述
	UserTodoDueTime     int    `json:"user_todo_due_time"`    //用户todo 截止时间
	UserTodoRemindTime  int    `json:"user_todo_remind_time"` //用户todo 开始提醒时间
	Status              string `json:"status"`                //用户todo 状态 `1`表示有效,`2表示无效`

}
