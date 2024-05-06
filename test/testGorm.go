package main

import (
	"fmt"
	"github.com/fatih/structs"
	"go-chat/utils"
)

type Person struct {
	Name string `structs:"name"`
	Age  int    `structs:"age"`
}

func main() {
	//core.InitConf()
	//dsn := global.CONFIG.Mysql.Dsn()
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//
	//// Migrate the schema
	//db.AutoMigrate(&model.UserBasic{})
	//
	//// Create
	//db.Create(&model.UserBasic{Name: "liming", Password: "123"})
	//
	//// Read
	//var user model.UserBasic
	//db.First(&user, 1)                    // find product with integer primary key
	//db.First(&user, "name = ?", "liming") // find product with code D42
	//
	//// Update - update product's price to 200
	//db.Model(&user).Update("Name", "liming3")
	//// Update - update multiple fields
	//db.Model(&user).Updates(model.UserBasic{Name: "liming2", Password: "123456"}) // non-zero fields
	//db.Model(&user).Updates(map[string]interface{}{"Name": "liming10086", "Password": "123456"})
	//
	//// Delete - delete product
	//db.Delete(&user, 1)
	//
	//res := model.UserBasic{}
	//db.Find(&res)
	//fmt.Println("res:", res)

	p := Person{
		Name: "xiaoming",
	}

	pMap := structs.Map(p)

	for k, v := range pMap {
		if !utils.IsEmpty(v) {
			fmt.Println(k, v)
		}
	}
}
