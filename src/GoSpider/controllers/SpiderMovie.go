/*
@Time : 2019-11-12 上午 9:19
@Author : Gerald
@File : SpiderMovie.go
@Software: GoLand
*/
package controllers

import (
	"GoSpider/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"time"
)

type MovieController struct {
	beego.Controller
}

func (this *MovieController) Spider() {
	// 连接 redis
	models.ConnectRedis("127.0.0.1:6379")

	// 爬虫入口 url
	url := "https://movie.douban.com/subject/30242710/?from=showing"

	// 先将当前的 url 入队
	models.PushQueue(url)

	var movieInfo models.MovieInfo

	for true {
		// 先判断队列是否为空
		length := models.GetQueueLength()

		if length == 0 { // 如果队列为空，说明爬虫结束，退出
			this.Ctx.WriteString("Queue is empty!\n")
			break
		}

		// 先获得当前队首 url
		url := models.PopQueue()

		// 判断当前 url 是否已经存在过
		isVisited := models.IsVisited(url)
		if isVisited {
			continue
		}

		// 获得当前 url 的 HTML 源码
		rsp := httplib.Get(url)
		movieHtml, _ := rsp.String()

		// 如果不存在，就先获取电影名
		// 获得电影名
		movieInfo.MovieName = models.GetMovieName(movieHtml)
		this.Ctx.WriteString(movieInfo.MovieName + "\n")
		if movieInfo.MovieName != "" { // 如果可以获取到电影名，说明是电影信息网页，那么就提取网页信息，插入到数据库
			// 获得导演名
			movieInfo.MovieDirector = models.GetMovieDirector(movieHtml)
			// 获得电影种类
			movieInfo.MovieType = models.GetMovieType(movieHtml)
			// 获得电影主演
			movieInfo.MovieMainCharacter = models.GetMovieMainCharacter(movieHtml)
			// 获得电影编剧
			movieInfo.MovieWriter = models.GetMovieWriter(movieHtml)
			// 获得电影制作国家
			movieInfo.MovieCountry = models.GetMovieCountry(movieHtml)
			// 获得电影语种
			movieInfo.MovieLanguage = models.GetMovieLanguage(movieHtml)
			// 获得电影的上映时间
			movieInfo.MovieOnTime = models.GetMovieOnTime(movieHtml)
			// 获得电影时长
			movieInfo.MovieSpan = models.GetMovieSpan(movieHtml)
			// 获得电影评分
			movieInfo.MovieGrade = models.GetMovieGrade(movieHtml)

			// 将抓取的电影信息插入到数据库中
			models.AddMovieInfo(&movieInfo)

			// 在这里打印成功插入数据库的电影名
			this.Ctx.WriteString(movieInfo.MovieName + "\n")
		}

		// 获取网页的所有 url
		urls := models.GetHtmlUrls(movieHtml)
		for _, value := range urls {
			models.PushQueue(value)
			models.AddToSet(url)
		}

		time.Sleep(time.Second)
	}

	this.Ctx.WriteString("end~bye~\n")
}
