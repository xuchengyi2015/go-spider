package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/xuchengyi2015/go-spider/models"
	"github.com/xuchengyi2015/go-spider/tools"
	"time"
)

// Operations about object
type MovieInfoController struct {
	beego.Controller
}

// @router / [get]
func (m *MovieInfoController) GetMovieInfo() {
	//url := `https://movie.douban.com/subject/30211551/`
	//rsp := httplib.Get(url)
	//html, err := rsp.String()
	//tools.CheckErr(err)
	//
	//movie := models.GetMovieInfo(html)
	//m.Data["json"] = movie
	m.ServeJSON()
}

// @router /start_crawl [get]
func (m *MovieInfoController) CrawlMovie() {
	rootUrl := `https://movie.douban.com/subject/6786002/`
	models.PushQueue(rootUrl) //将根站点加入队列

	var length int
	for {
		length = models.GetQueueLength()
		if length == 0 {
			break //如果Url队列为空了，则退出程序
		}

		url := models.PopQueue()
		if models.IsVisitedUrl(url) {
			continue
		}
		response := httplib.Get(url)
		html, err := response.String()
		tools.CheckErr(err)

		if models.GetMovieName(html) != "" {
			_, err := models.GetMovieInfo(html)
			if err != nil {
				m.Data["json"] = tools.GetResult(
					tools.WithCode(1),
					tools.WithMessage("添加数据失败"),
					tools.WithData(err.Error()),
				)
				m.ServeJSON()
				return
			}
			models.SetQueue(url)
		}

		urls := models.GetMoviesUrls(html)
		for _, v := range urls {
			models.PushQueue(v)
		}

		time.Sleep(time.Second)
	}

	m.Data["json"] = tools.GetResult()
	m.ServeJSON()
}
