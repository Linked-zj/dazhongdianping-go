package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func main() {
	//
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
	//	logger.Config{
	//		SlowThreshold:             time.Second, // 慢 SQL 阈值
	//		LogLevel:                  logger.Info, // 日志级别
	//		IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
	//		Colorful:                  true,        // 禁用彩色打印
	//	},
	//)
	//dsn := "root:12345678@tcp(127.0.0.1:3306)/dazhongdianping?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	Logger: newLogger,
	//	NamingStrategy: schema.NamingStrategy{
	//		SingularTable: true,
	//	},
	//})
	//if err != nil {
	//	fmt.Println("不存在")
	//}
	//db.AutoMigrate(&model.TbProduct{})
	//
	//var p model.TbProduct
	//db.Find(&p)
	//fmt.Println(p.ID)

	img := "\\image\\product\\haitai.png"
	fmt.Println(img)

	replace := strings.Replace(img, "\\", "/", -1)
	fmt.Println(replace)

}
