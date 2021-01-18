package app

import (
	"time"

	"github.com/bo-er/todo-simpler/todo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/bo-er/todo-simpler/admin/config"
)

// Config 是数据库配置项
type Config struct {
	Debug        bool   //是否开启debug模式
	DBType       string //数据库类型
	DSN          string //数据源名称
	MaxLifetime  int    //最长连接时间
	MaxOpenConns int    //最大连接数量
	MaxIdleConns int    //最大空闲的连接数
}

// InitGormDB 初始化gorm数据库连接
func InitGormDB() (*gorm.DB, error) {
	cfg := config.C
	var dsn string
	dsn = cfg.Postgres.DSN()
	db, err := newDB(&Config{
		Debug:        cfg.Gorm.Debug,
		DBType:       cfg.Gorm.DBType,
		DSN:          dsn,
		MaxIdleConns: cfg.Gorm.MaxIdleConns,
		MaxLifetime:  cfg.Gorm.MaxLifetime,
		MaxOpenConns: cfg.Gorm.MaxOpenConns,
	})
	if err != nil {
		return nil, err
	}
	if cfg.Gorm.EnableAutoMigrate {
		err = autoMigrate(db)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

func newDB(c *Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(c.DSN), &gorm.Config{
		DisableAutomaticPing: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	if c.Debug {
		db = db.Debug()
	}
	sqlDB, err := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Duration(c.MaxLifetime) * time.Second)

	return db, nil
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(new(todo.UserTodo))
}
