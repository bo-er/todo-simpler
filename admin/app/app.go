package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bo-er/todo-simpler/admin/api"
	"github.com/bo-er/todo-simpler/admin/config"
	"github.com/bo-er/todo-simpler/admin/routers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 声明全局变量
var (
	DB *gorm.DB
)

type options struct {
	ConfigFile string //配置文件
}

// Option 定义Option函数类型
type Option func(*options)

// SetConfigFile 设置配置文件,返回Option类型的函数
func SetConfigFile(s string) Option {
	return func(o *options) {
		o.ConfigFile = s
	}
}

// Init 应用初始化,opts是Option类型的函数，比如SetConfigFile
func Init(ctx context.Context, opts ...Option) *gin.Engine {
	var o options
	// 遍历执行Option类型的函数
	for _, optionFunc := range opts {
		optionFunc(&o)
	}

	config.MustLoad(o.ConfigFile)
	log.Printf("Todo应用启动\n")
	log.Printf("Todo应用运行模式:%s\n", config.C.RunMode)
	log.Printf("Todo应用进程号:%d\n", os.Getpid())
	db, err := InitGormDB()
	if err != nil {
		log.Fatalf("数据库初始化失败: %s", err.Error())
	}
	DB = db
	apiTodo := &api.Todo{
		TodoService: MustGetTodoService(),
	}
	routerRouter := &routers.Router{
		TodoAPI: apiTodo,
	}
	engine := InitGinEngine(routerRouter)
	return engine
}

// Run 运行服务
func Run(ctx context.Context, opts ...Option) error {
	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	engine := Init(ctx, opts...)
	go func() {
		engine.Run(":7088")
	}()

EXIT:
	for {
		fmt.Println("hello！")
		sig := <-sc
		log.Printf("接收到信号[%s]", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	log.Printf("服务退出")
	time.Sleep(time.Second)
	os.Exit(state)
	return nil
}
