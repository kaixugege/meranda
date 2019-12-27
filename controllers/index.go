package controllers

type IndexController struct {
	BaseController
}


// @router / [get]
func (c *IndexController) Get() {
	c.TplName = "index.html"
}
