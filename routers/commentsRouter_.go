package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["meranda/controllers:CategoriesContller"] = append(beego.GlobalControllerRouter["meranda/controllers:CategoriesContller"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/categories`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["meranda/controllers:IndexController"] = append(beego.GlobalControllerRouter["meranda/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
