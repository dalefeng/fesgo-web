package service

import (
	"fmt"
	"github.com/dalefeng/fesgo/orm"
)

type User struct {
	ID       int64
	Username string
	Password string
	Age      int
}

func SaveUser() {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", "dev", "13140236", "fungs.cn", 3306, "fesgo")
	db := orm.Open("mysql", dataSource)
	user := User{
		Username: "user1",
		Password: "123456",
		Age:      18,
	}
	db.NewSession().Table("fes_user").Insert(&user)
}
