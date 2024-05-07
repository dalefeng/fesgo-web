package service

import (
	"fmt"
	"github.com/dalefeng/fesgo/orm"
)

type User struct {
	Id       int64  `fesgo:"id,auto_increment"`
	UserName string `fesgo:"username" form:"username" json:"username"`
	Password string `form:"password"`
	Age      int    `form:"age"`
}

func SaveUser() {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", "root", "yY13140236", "fungs.cn", 13306, "fesgo")
	db := orm.Open("mysql", dataSource)
	db.Prefix = "fes_"
	user := &User{
		UserName: "user1",
		Password: "123456",
		Age:      18,
	}
	user1 := &User{
		UserName: "user2",
		Password: "123456",
		Age:      21,
	}
	users := []any{user, user1}
	insert, _, err := db.NewSession(&User{}).InsertBatch(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(insert)
	db.Close()
}

func UpdateUser() {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", "root", "yY13140236", "fungs.cn", 13306, "fesgo")
	db := orm.Open("mysql", dataSource)
	db.Prefix = "fes_"

	var user User

	insert, _, err := db.NewSession(&user).Where("id", 7).Update("age", 21)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(insert)
	db.Close()
}

func SelectOne() {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", "root", "yY13140236", "fungs.cn", 13306, "fesgo")
	db := orm.Open("mysql", dataSource)
	db.Prefix = "fes_"
	var user User
	err := db.NewSession(&user).Where("id", 7).SelectOne(&user)
	if err != nil {
		fmt.Println(err)
	}
	db.Close()
	fmt.Printf("%+v \n", user)
}

func Select() {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", "root", "yY13140236", "fungs.cn", 13306, "fesgo")
	db := orm.Open("mysql", dataSource)
	db.Prefix = "fes_"
	var user User
	list, err := db.NewSession(&user).Where("age", 21).Select(&user)
	if err != nil {
		fmt.Println(err)
	}
	db.Close()
	fmt.Printf("%+v \n", list)
}
