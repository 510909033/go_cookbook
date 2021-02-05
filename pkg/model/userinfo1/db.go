package userinfo1

import (
	"510909033/go_cookbook/pkg/mylog"
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"runtime"
	"time"
)

var db *gorm.DB
var Mydb = db
var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold: time.Second, // 慢 SQL 阈值
		LogLevel:      logger.Info, // Log level
		Colorful:      false,       // 禁用彩色打印
	},
)
var tmp int
func init() {
	tmp = 23

	log.SetFlags(log.Llongfile)
	log.Println(tmp,&tmp)
	fmt.Println("----------------------init---------------------------\n")

	ctx := context.Background()
	wbtLogger := mylog.NewLoggerWithTransactionId(ctx)
	//db, err := gorm.Open("mysql","root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	dsn := "mytest1:JkL^%3kLs@tcp(rm-bp140ee29co2u46j84o.mysql.rds.aliyuncs.com:3306)/cook1?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		wbtLogger.Error("gorm.Open返回err,err=", fmt.Sprintf("%+v", err))
		panic(err)
	}
	//db.Debug()

	//defer db.Close()

	sqlDB, _ := db.DB()
	_ = sqlDB

	//sqlDB.Close()
	sqlDB.SetMaxOpenConns(30)
	sqlDB.SetMaxIdleConns(30)
	sqlDB.SetConnMaxLifetime(time.Second)

	wbtLogger.Info("sqlDB设置相关参数， SetMaxOpenConns = 30, SetMaxIdleConns = 30")

	go func() {
		for {
			wbtLogger.Info(fmt.Sprintf("sqlDB.Stats = %+v\n", sqlDB.Stats()))
			time.Sleep(time.Second * 30)
		}
	}()

}

func Usetimes() func() {
	t := time.Now()
	pc, file, line, ok := runtime.Caller(1)
	_ = line
	_ = ok
	_ = file
	return func() {

		fmt.Printf("use time:%s, pc=%s, "+
			//"file=%s, " +
			//"line=%d, " +
			//"ok=%b" +
			"\n",
			time.Since(t).String(),
			runtime.FuncForPC(pc).Name(),
			//file,
			//line,
			//ok,
		)
	}
}

func GetDB() *gorm.DB {
	return db
}
