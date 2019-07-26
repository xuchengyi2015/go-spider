package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/xuchengyi2015/go-spider/models"
	"github.com/xuchengyi2015/go-spider/tools"
)

// Operations about object
type MovieInfoController struct {
	beego.Controller
}

// @router / [get]
func (m *MovieInfoController) GetMovieInfo() {
	url := `https://movie.douban.com/subject/30211551/`
	rsp := httplib.Get(url)
	html, err := rsp.String()
	tools.CheckErr(err)

	movie := models.GetMovieInfo(html)
	m.Data["json"] = movie
	m.ServeJSON()
}

// @router /getlist [get]
func (m *MovieInfoController) GetMovieUrlOther() {
	url := `https://movie.douban.com/explore#!type=movie&tag=%E8%B1%86%E7%93%A3%E9%AB%98%E5%88%86&sort=time&page_limit=20&page_start=0`

	rsp := httplib.Get(url)
	html, err := rsp.String()
	tools.CheckErr(err)

	info := models.GetMovieUrlOther(html)
	m.Data["json"] = info
	m.ServeJSON()
}
