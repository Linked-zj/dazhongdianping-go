package initialize

import (
	"dazhongdianping-go/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"strconv"
	"time"
)

func GormInit() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	host := global.ServerConfig.MysqlConfig.Host
	port := global.ServerConfig.MysqlConfig.Port
	username := global.ServerConfig.MysqlConfig.Username
	password := global.ServerConfig.MysqlConfig.Password

	dsn := username + ":" + password + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/dazhongdianping?charset=utf8mb4&parseTime=True&loc=Local"

	fmt.Println(dsn)
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
}
