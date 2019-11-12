package routers

import (
	"GoSpider/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/spider_movie", &controllers.MovieController{}, "*:Spider")
}
