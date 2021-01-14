package config

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/koding/multiconfig"
	"github.com/bo-er/todo-simpler/utils/json"
)

var (
	// C 全局配置(需要先执行MustLoad，否则拿不到配置)
	C    = new(Config)
	once sync.Once
)

// MustLoad 加载配置
func MustLoad(fpaths ...string) {
	once.Do(func() {
		loaders := []multiconfig.Loader{
			&multiconfig.TagLoader{},
			&multiconfig.EnvironmentLoader{},
		}

		for _, fpath := range fpaths {
			if strings.HasSuffix(fpath, "toml") {
				loaders = append(loaders, &multiconfig.TOMLLoader{Path: fpath})
			}
			if strings.HasSuffix(fpath, "json") {
				loaders = append(loaders, &multiconfig.JSONLoader{Path: fpath})
			}
			if strings.HasSuffix(fpath, "yaml") {
				loaders = append(loaders, &multiconfig.YAMLLoader{Path: fpath})
			}
		}

		m := multiconfig.DefaultLoader{
			Loader:    multiconfig.MultiLoader(loaders...),
			Validator: multiconfig.MultiValidator(&multiconfig.RequiredValidator{}),
		}
		m.MustLoad(C)
	})
}

// PrintWithJSON 基于JSON格式输出配置
func PrintWithJSON() {
	if C.PrintConfig {
		b, err := json.MarshalIndent(C, "", " ")
		if err != nil {
			os.Stdout.WriteString("[CONFIG] JSON marshal error: " + err.Error())
			return
		}
		os.Stdout.WriteString(string(b) + "\n")
	}
}

// Config 配置参数
type Config struct {
	RunMode      string       //运行模式
	WWW          string       //前端网页存放路径
	Swagger      bool         //是否开启swagger
	PrintConfig  bool         //是否打印设置
	HTTP         HTTP         //配置HTTP
	Menu         Menu         //配置菜单
	Casbin       Casbin       //casbin配置
	Log          Log          //日志配置
	LogGormHook  LogGormHook  //gorm日志存储钩子
	LogMongoHook LogMongoHook //Mongo日志存储钩子
	Root         Root         //Root用户配置
	JWTAuth      JWTAuth      //JWT认证配置
	Monitor      Monitor      //监控配置
	Captcha      Captcha      //图形验证码
	RateLimiter  RateLimiter  //请求频率限制参数
	CORS         CORS         //跨域请求配置
	GZIP         GZIP         //GZIP配置
	Redis        Redis        //Redis配置
	Gorm         Gorm         //Gorm配置
	Postgres     Postgres     //Postgres配置
	Sqlite3      Sqlite3      //Sqlite3配置
}

// IsDebugMode 是否是debug模式
func (c *Config) IsDebugMode() bool {
	return c.RunMode == "debug"
}

// Menu 菜单配置参数
type Menu struct {
	Enable bool   //是否启用
	Data   string //菜单数据
}

// Casbin casbin配置参数
type Casbin struct {
	Enable           bool   //是否启用casbin
	Debug            bool   //是否开debug模式
	Model            string //casbin模型
	AutoLoad         bool   //是否自动加载
	AutoLoadInternal int    //自动加载时间
}

// LogHook 日志钩子
type LogHook string

// IsGorm 是否是gorm钩子
func (h LogHook) IsGorm() bool {
	return h == "gorm"
}

// IsMongo 是否是mongo钩子
func (h LogHook) IsMongo() bool {
	return h == "mongo"
}

// Log 日志配置参数
type Log struct {
	Level         int      //日志打印界别
	Format        string   //日志格式
	Output        string   //输出
	OutputFile    string   //输出文件
	EnableHook    bool     //是否启用钩子
	HookLevels    []string //钩子的级别
	Hook          LogHook  //日志钩子
	HookMaxThread int      //钩子最大线程数量
	HookMaxBuffer int      //钩子最大缓冲大小
}

// LogGormHook 日志gorm钩子配置
type LogGormHook struct {
	DBType       string //数据库类型
	MaxLifetime  int    //最长连接时间
	MaxOpenConns int    //最大开放连接数量
	MaxIdleConns int    //最大空闲连接数量
	Table        string //表
}

// LogMongoHook 日志mongo钩子配置
type LogMongoHook struct {
	Collection string //mongodb数据库表名
}

// Root root用户
type Root struct {
	UserName string //用户名
	Password string //密码
	RealName string //真实姓名
}

// JWTAuth 用户认证
type JWTAuth struct {
	Enable        bool   //是否启用
	SigningMethod string //签名方法
	SigningKey    string //签名
	Expired       int    //过期时间
	Store         string //存储令牌的数据库名称
	FilePath      string //文件路径
	RedisDB       int    //Redis数据库地址
	RedisPrefix   string //Redis前缀名称
}

// HTTP http配置参数
type HTTP struct {
	Host             string //服务器地址
	Port             int    //端口
	CertFile         string //证书文件
	KeyFile          string //密钥文件
	ShutdownTimeout  int    //超时关闭时间
	MaxContentLength int64  //最大内容长度
	MaxLoggerLength  int    `default:"4096"` //最大日志长度
}

// Monitor 监控配置参数
type Monitor struct {
	Enable    bool   //是否启用监控
	Addr      string //监控服务地址
	ConfigDir string //监控配置目录
}

// Captcha 图形验证码配置参数
type Captcha struct {
	Store       string //图形验证吗存储方式
	Length      int    //验证码文本长度
	Width       int    //图形验证码宽度
	Height      int    //图形验证码高度
	RedisDB     int    //Redis数据库地址
	RedisPrefix string //Redis前缀
}

// RateLimiter 请求频率限制配置参数
type RateLimiter struct {
	Enable  bool  //是否启用请求频率限制
	Count   int64 //请求次数统计
	RedisDB int   //Redis数据库地址
}

// CORS 跨域请求配置参数
type CORS struct {
	Enable           bool     //是否启用跨域请求
	AllowOrigins     []string //允许访问的来源IP地址
	AllowMethods     []string //允许的HTTP请求方法
	AllowHeaders     []string //允许的请求头
	AllowCredentials bool     //是否允许
	MaxAge           int      //预检查的最长有效期
}

// GZIP gzip压缩
type GZIP struct {
	Enable             bool     //是否开启Gzip压缩
	ExcludedExtentions []string //需要排除的扩展
	ExcludedPaths      []string //需要排除的路径
}

// Redis redis配置参数
type Redis struct {
	Addr     string //Redis数据库服务地址
	Password string //Redis数据库密码
}

// Gorm gorm配置参数
type Gorm struct {
	Debug             bool   //是否开启debug模式
	DBType            string //数据库类型
	MaxLifetime       int    //最长连接时间
	MaxOpenConns      int    //最大开放连接数量
	MaxIdleConns      int    //最大空闲的连接数量
	TablePrefix       string //表前缀
	EnableAutoMigrate bool   //是否自动开启迁移
}

// Postgres postgres配置参数
type Postgres struct {
	Host     string //Postgres数据库地址
	Port     int    //Postgres端口号
	User     string //Postgres用户
	Password string //Postgres密码
	DBName   string //Postgres数据库名称
	SSLMode  string //Postgres SSL模式
}

// DSN 数据库连接串
func (a Postgres) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		a.Host, a.Port, a.User, a.DBName, a.Password, a.SSLMode)
}

// Sqlite3 sqlite3配置参数
type Sqlite3 struct {
	Path string //数据库地址
}

// DSN 数据库连接串
func (a Sqlite3) DSN() string {
	return a.Path
}
