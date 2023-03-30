package main

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Name         string  //允许读和创建
	Email        *string //允许读和更新
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	gorm.Model

	//Name string `gorm:"<-:create"` // 允许读和创建
	//Name string `gorm:"<-:update"` // 允许读和更新
	//Name string `gorm:"<-"`        // 允许读和写（创建和更新）
	//Name string `gorm:"<-:false"`  // 允许读，禁止写
	//Name string `gorm:"->"`        // 只读（除非有自定义配置，否则禁止写）
	//Name string `gorm:"->;<-:create"` // 允许读和写
	//Name string `gorm:"->:false;<-:create"` // 仅创建（禁止从 db 读）
	//Name string `gorm:"-"`  // 读写操作均会忽略该字段
	//Updated   int64 `gorm:"autoUpdateTime:nano"` // 使用时间戳填纳秒数充更新时间
	//Updated   int64 `gorm:"autoUpdateTime:milli"` // 使用时间戳毫秒数填充更新时间
	//Created   int64 `gorm:"autoCreateTime"`      // 使用时间戳秒数填充创建时间
}
type Author struct {
	Name  string
	Email string
}

type Blog struct {
	ID     int
	Author Author `gorm:"embedded"` //正常的结构体字段，你也可以通过标签 embedded 将其嵌入
	//Author  Author  `gorm:"embedded;embeddedPrefix:author_"` //可以使用标签 embeddedPrefix 来为 db 中的字段名添加前缀
	Upvotes int32
}

// Blog等效于Blog2
type Blog2 struct {
	ID      int64
	Name    string
	Email   string
	Upvotes int32
}
