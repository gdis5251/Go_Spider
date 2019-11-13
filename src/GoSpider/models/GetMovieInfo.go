/*
@Time : 2019-11-12 上午 9:14
@Author : Gerald
@File : GetMovieInfo.go
@Software: GoLand
*/
package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
)

var (
	db orm.Ormer
)

type MovieInfo struct {
	Id                 int64
	MovieId            int64
	MovieName          string
	MoviePic           string
	MovieDirector      string
	MovieWriter        string
	MovieCountry       string
	MovieLanguage      string
	MovieMainCharacter string
	MovieType          string
	MovieOnTime        string
	MovieSpan          string
	MovieGrade         string
}

func init() {
	orm.RegisterModel(new(MovieInfo))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:peiban493.@tcp(127.0.0.1:3306)/spider_movie?charset=utf8", 30)
	db = orm.NewOrm()
}

func AddMovieInfo(movieInfo *MovieInfo) (int64, error) {
	id, insertErr := db.Insert(movieInfo)
	if insertErr != nil {
		fmt.Println(insertErr)
	}

	return id, insertErr
}

func GetMovieDirector(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}

	// 使用正则表达式来匹配导演标签
	// <a href="/celebrity/1040524/" rel="v:directedBy">彼得·杰克逊</a>
	reg := regexp.MustCompile(`<a .*? rel="v:directedBy">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	return result[0][1]
}

func GetMovieName(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}

	// 使用正则表达式匹配电影名
	// <span property="v:itemreviewed">他们已不再变老 They Shall Not Grow Old</span>
	reg := regexp.MustCompile(`<span.*?property="v:itemreviewed">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	return result[0][1]
}

func GetMovieType(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}

	// 使用正则表达式匹配电影类型
	// <span property="v:genre">纪录片</span>
	reg := regexp.MustCompile(`<span.*?property="v:genre">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	// 因为电影种类可能会有多种，使用循环把所有种类拼接在一起
	var movieType string
	for index, value := range result {
		movieType += value[1]
		if index != len(result)-1 {
			movieType += `/`
		}
	}

	return movieType
}

func GetMovieCountry(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}

	// 使用正则表达式匹配电影国家
	// <span class="pl">制片国家/地区:</span>
	//  英国 / 新西兰
	// <br/>
	reg := regexp.MustCompile(`<span class="pl">制片国家/地区:</span>\s(.*?)<br/>\s`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0 {
		return "Can't Match Any Country Or Regexp Error!!!!"
	}

	return result[0][1]
}

func GetMovieMainCharacter(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}

	// 使用正则表达式匹配电影主演
	// <a href="/celebrity/1324043/" rel="v:starring">大鹏</a>
	reg := regexp.MustCompile(`<a .*? rel="v:starring">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	// 因为电影种类可能会有多种，使用循环把所有种类拼接在一起
	var movieMainCharacter string
	for index, value := range result {
		movieMainCharacter += value[1]
		if index != len(result)-1 {
			movieMainCharacter += `/`
		}
	}

	return movieMainCharacter
}

func GetMovieWriter(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}

	// 使用正则表达式匹配编剧
	// <a href="/celebrity/1404918/">申奥</a>
	reg := regexp.MustCompile(`<a href="/celebrity/\d{7}/">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0 {
		return "Can't Match Any Writer Or Regexp Error!!!!"
	}

	var movieWriter string
	for index, value := range result {
		movieWriter += value[1]
		if index != len(result)-1 {
			movieWriter += `/`
		}
	}

	return movieWriter
}

func GetMovieLanguage(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}

	// 使用正则表达式匹配电影语种
	// <span class="pl">语言:</span>
	//  汉语普通话
	// <br/>
	reg := regexp.MustCompile(`<span class="pl">语言:</span>\s(.*?)<br/>\s`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0 {
		return "Can't Match Any Language Or Regexp Error!!!!"
	}

	var movieLanguage string
	for index, value := range result {
		movieLanguage += value[1]
		if index != len(result)-1 {
			movieLanguage += `/`
		}
	}

	return movieLanguage
}

func GetMovieOnTime(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}

	// 使用正则表达式匹配电影上映时间
	// <span property="v:initialReleaseDate" content="2019-11-08(中国大陆)">2019-11-08(中国大陆)</span>
	reg := regexp.MustCompile(`<span property="v:initialReleaseDate" content="[0-9-]{10}.*?">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0 {
		return "Can't Match Any Language Or Regexp Error!!!!"
	}

	var movieOnTime string
	for index, value := range result {
		movieOnTime += value[1]
		if index != len(result)-1 {
			movieOnTime += `/`
		}
	}

	return movieOnTime
}

func GetMovieSpan(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}

	// 使用正则表达式匹配电影时长
	// <span property="v:runtime" content="112">112分钟</span>
	reg := regexp.MustCompile(`<span property="v:runtime" content="[0-9]+">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0 {
		return "Can't Match Any Span Or Regexp Error!!!!"
	}

	return result[0][1]
}

func GetMovieGrade(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}

	// 使用正则表达式匹配评分
	// <strong class="ll rating_num" property="v:average">6.7</strong>
	reg := regexp.MustCompile(`<strong .*?property="v:average">(.*?)</strong>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0 {
		return "Can't Match Any Grade Or Regexp Error!!!!"
	}

	return result[0][1]
}
