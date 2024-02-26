package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/WelcomeSilverCity/do/model"
)

func main() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/shop?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "t_", // 表前缀
			SingularTable: true, // 表名单数
			//NoLowerCase:   true, //跳过蛇形命名
		},
	})
	if err != nil {
		panic(err)
	}
	// 自动迁移 (这是GORM自动创建表的一种方式--译者注)
	db.AutoMigrate(&model.Category{}, &model.Brands{}, &model.Banner{}, &model.GoodsCategoryBrand{}, &model.Goods{})

	//彩虹表

	//options := &password.Options{10, 10000, 50, md5.New}
	//salt, encodedPwd := password.Encode("123456", options)
	//newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd) // pbkdf2-sha512 $符号拼接

	//check := password.Verify("generic password", salt, encodedPwd, options)
	//fmt.Println(salt)       // true
	//fmt.Println(encodedPwd) // true
	//fmt.Println(newPassword) // true
}
