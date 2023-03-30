package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	"time"
)

// MySQL, PostgreSQL, SQlite, SQL Server
func ConfigMysql() *gorm.DB {
	//配置MySQL连接参数
	username := "root"  //账号
	password := "root"  //密码
	host := "localhost" //数据库地址，可以是Ip或者域名
	port := 3306        //数据库端口
	Dbname := "test"    //数据库名
	timeout := "10s"    //连接超时，10秒

	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	return db
}

func ConfigPostgreSql() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	FailOnErr(err, " fail connect to PostgreSQL")
	return db
}

func ConfigSQLite() *gorm.DB {
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	FailOnErr(err, "fail connect to SQLite")
	return db
}

func ConfigSQLServer() *gorm.DB {
	// github.com/denisenkom/go-mssqldb
	dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	FailOnErr(err, "fail connect to SQLServer")
	return db
}
func FailOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

func ConfigPool(db *gorm.DB) {
	sqlDB, err := db.DB()
	FailOnErr(err, "数据库连接出错")
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

}
