package db

import (
	"os"
	"time"

	l "log"

	"github.com/pangu-2/pangu-config/configs"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbEngine *gorm.DB

// demo:123456@(localhost)/demo?charset=utf8&parseTime=True&loc=Local
func dsn() string {
	dbC := configs.GetDatasource()
	dsn := ""
	if len(dbC.Url) > 0 {
		dsn = dbC.Username + ":" + dbC.Password + "@" + dbC.Url
	} else {
		dsn = dbC.Username + ":" + dbC.Password + "@(" + dbC.Db + ")/" + dbC.Db + "?charset=utf8mb4&parseTime=True&loc=Local"
	}
	log.Info("dsn = " + dsn)
	return dsn
}

func Db() *gorm.DB {
	if dbEngine == nil {
		log.Debugf("Model NewDB")

		newDb, err := newDb()
		if err != nil {
			panic(err)
		}
		sqlDB, err := newDb.DB()
		if err != nil {
			panic(err)
		}
		if err = sqlDB.Ping(); err != nil {
			log.Fatal("数据库 ping:", err.Error())
		}
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		dbEngine = newDb
	}

	return dbEngine
}

func newDb() (*gorm.DB, error) {
	//  默认的级别，会打印find找不到模型时的sql语句。
	// Silent 就不会。
	newLogger := logger.New(
		l.New(os.Stdout, "\r\n", l.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // 日志级别 可选 Silent，Error，Warn，Info
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,          // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(dsn()), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Error("failed to connect database")
		return nil, err
	}
	return db, nil
}

// 初始化
func Init() {
	Db()
	log.Info("init db")
}

func Close() {
}
