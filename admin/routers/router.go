package routers

import (
	"github.com/bo-er/todo-simpler/admin/api"
	"github.com/gin-gonic/gin"
)

// MyRouter 注册路由
type MyRouter interface {
	Register(app *gin.Engine) error //注册路由函数
	Prefixes() []string             //路由相关的特定请求前缀
}

// Router 是路由管理器
type Router struct {
	TodoAPI *api.Todo
}

// Register 注册路由
func (r *Router) Register(app *gin.Engine) error {
	r.RegisterAPI(app)
	return nil
}

// Prefixes 返回路由前缀列表
func (r *Router) Prefixes() []string {
	return []string{
		"/api/",
	}
}
