package app

import (
	"github.com/LyricTian/gzip"
	"github.com/bo-er/todo-simpler/admin/config"
	"github.com/bo-er/todo-simpler/admin/middlewares"
	"github.com/bo-er/todo-simpler/admin/routers"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitGinEngine 用于初始化Gin引擎
func InitGinEngine(r routers.MyRouter) *gin.Engine {

	app := gin.New()
	// GZIP
	if config.C.CORS.Enable {
		app.Use(middlewares.CORSMiddleware())
	}

	if config.C.GZIP.Enable {
		app.Use(gzip.Gzip(gzip.BestCompression,
			gzip.WithExcludedExtensions(config.C.GZIP.ExcludedExtentions),
			gzip.WithExcludedPaths(config.C.GZIP.ExcludedPaths),
		))
	}

	// Router register
	_ = r.Register(app)
	if config.C.Swagger {
		app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	return app
}
