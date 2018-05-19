package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Log struct {
	Id          int
	AppName        string
	Content     string
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func (this *Log) Save() {
	o := orm.NewOrm()
	println(" insert")
	_ ,err := o.Insert(this)
	if err != nil {
		println("数据库插入失败")
	}

}
