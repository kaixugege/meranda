package controllers

type CategoriesContller struct {
	BaseController
}

//主页新闻
// @router /categories [get]
func (c *CategoriesContller) Get() {
	c.TplName = "categories.html"
}
