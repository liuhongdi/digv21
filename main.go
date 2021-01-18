package main

import (
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/liuhongdi/digv21/router"
)

func init() {
}

func main() {
	//引入路由
	r := router.Router()
	//run
	r.Run(":8080")
}




