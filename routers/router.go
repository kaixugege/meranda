package routers

import (
	"github.com/astaxie/beego"
	"meranda/controllers"
)

func init() {

	beego.Include(
		&controllers.IndexController{},
		&controllers.CategoriesContller{})
	//beego.Router("/", &controllers.MainController{})
}
