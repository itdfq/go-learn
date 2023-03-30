package entry

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	Id           uint `gorm:"primarykey"`
	Name         string
	Email        *string
	Age          uint8
	Birthday     JsonTime
	MemberNumber sql.NullString
	CreatedAt    JsonTime
	UpdatedAt    JsonTime
	DeletedAt    JsonTime `gorm:"index"`
}

// 指定表名
func (User) TableName() string {
	return "user"
}

// 钩子函数
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	log.Printf("执行了BeforeCreate")
	return nil
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	log.Printf("BeforeSave执行")
	return nil
}
func (u *User) AfterSave(tx *gorm.DB) (err error) {
	log.Printf("AfterSave执行")
	return nil
}
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	log.Printf("AfterCreate 执行")
	return nil
}

type JsonTime time.Time

// 实现它的json序列化方法

func (this JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
