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
	"strconv"
)

type MovieController struct {
	beego.Controller
}

func (this *MovieController) Spider() {
	url := "https://movie.douban.com/subject/30242710/?from=showing"
	rsp := httplib.Get(url)

	// 获得网页源码
	movieHtml, error := rsp.String()
	if error != nil {
		panic(error)
	}

	var movieInfo models.MovieInfo

	// 获得电影名
	movieInfo.MovieName = models.GetMovieName(movieHtml)
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
	id, insertErr := models.AddMovieInfo(&movieInfo)
	if insertErr != nil {
		panic(insertErr)
	}
	this.Ctx.WriteString(strconv.Itoa(int(id)))
}
