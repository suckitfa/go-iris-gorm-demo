package main

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris/context"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
)

type User struct {
	Id    int64  `gorm:"PRIMARY_KEY;AUTO_INCREMENT"` // 编号，主键、自增长
	Name  string `gorm:"size:32;not null"`           // 用户名，长度32，不可空
	Email string `gorm:"size:128;not null"`          // 邮箱，长度128，不可空
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/test_gorm?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database!")
	}
	defer db.Close()
	// 根据上面定义好的模型自动生成表
	db.AutoMigrate(&User{})
	// 往表中加入一条数据
	// db.Create(&User{
	// 	Name:  "John",
	// 	Email: "testemailjhon@gmail.com"})
	// db.Create(&User{
	// 	Name:  "Bob",
	// 	Email: "testemailbob@gmail.com"})

	// 根据字段查找用户
	// var user User
	// db.Where("name = ?", "John").First(&user)
	// fmt.Println("小明的邮箱1 = ", user.Email)
	// 重新复制后，再次保存
	// user.Email = "helloworldJohn@gmail.com"
	// db.Save(user)
	// 从数据库中取出
	// db.Where("name = ?", "John").First(&user)
	// fmt.Println("小明的更新后的邮箱 = ", user.Email)

	// 查询列表数据, gorm帮我们封装好了
	// var users []User
	// db.Find(&users)
	// fmt.Println("总用户数", len(users))

	// 删除数据
	db.Where("name = ?", "John").Delete(&User{})
	var users []User
	db.Find(&users)
	fmt.Println("删除一个用户之后的人数 = ", len(users))

}

func test() {
	app := iris.New()

	// 迅速回复重启，保证程序不会崩溃
	app.Logger().SetLevel(("debug"))
	app.Use(recover.New())
	app.Use(logger.New())

	count := 0

	// 中间件使用
	mvc.Configure(app.Party("/root"), func(mvcApp *mvc.Application) {
		mvcApp.Router.Use(func(context context.Context) {
			if context.Path() == "/root/test" {
				count++
				fmt.Println("/root/test请求数目 = ", strconv.Itoa(count))
			}
			// 加上他让Web进程继续往下执行，否则不会执行controller方法
			context.Next()
		})
		// 使用自己定义的控制器
		mvcApp.Handle(new(MyController))
	})

	// Method get
	// resource : http:localhost:8080/html
	app.Handle("GET", "/html", func(ctx iris.Context) {
		ctx.HTML("<h1> Hello world! </h1>")
	})

	// Method get
	// resource : http:localhost:8080/json
	app.Get("/json", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"Name":    "Iris",
			"Message": "Hello World!",
		})
	})

	// 启动app
	app.Run(iris.Addr(":1919"), iris.WithoutServerError((iris.ErrServerClosed)))
}
