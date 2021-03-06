package routers

import "github.com/gin-gonic/gin"

// RegisterAPI 注册路由API
func (r *Router) RegisterAPI(app *gin.Engine) {
	g := app.Group("api")

	v1 := g.Group("/v1")
	{
		gTodo := v1.Group("/todo")
		{
			gTodo.POST("", r.TodoAPI.AddUserTodo)
		}
	}
}
