package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"github.com/xuchengyi2015/go-spider/tools"
	"regexp"
	"strings"
	"time"
)

type MovieInfo struct {
	Id             int    `orm:"column(id)"`
	MovieName      string `orm:"column(movie_name)"`
	MovieCharacter string `orm:"column(movie_character)"`
	MovieRate      string `orm:"column(movie_rate)"`
	CreationTime   string `orm:"column(creation_time)"`
}

const dbString = "host=localhost port=5432 user=postgres password=postgres dbname=go-spider sslmode=disable"

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", dbString)
	orm.RegisterModel(new(MovieInfo))
}

func AddMovie(movie *MovieInfo) {
	o := orm.NewOrm()
	id, err := o.Insert(movie)
	tools.CheckErr(err)

	fmt.Printf("插入数据ID:%v\n", id)
}

func GetMovieInfo(MovieHtml string) MovieInfo {
	movie := MovieInfo{
		MovieName:      getMovieName(MovieHtml),
		MovieCharacter: getMovieCharacter(MovieHtml),
		MovieRate:      getMovieRate(MovieHtml),
		CreationTime:   time.Now().Format("2006-01-02 15:04:05"),
	}

	AddMovie(&movie)

	return movie
}

func getMovieInfoByRegex(movieHtml string, regstr string) string {
	reg := regexp.MustCompile(regstr)
	res := reg.FindAllStringSubmatch(movieHtml, -1)

	character := ""
	for _, v := range res {
		character += v[1] + " | "
	}
	return strings.Trim(character, " | ")
}

func getMovieName(MovieHtml string) string {
	// <span property="v:itemreviewed">哪吒之魔童降世</span>

	regstr := `<span\s*property="v:itemreviewed">(.*?)</span>`
	return getMovieInfoByRegex(MovieHtml, regstr)
}

func getMovieCharacter(MovieHtml string) string {
	// <a href="/celebrity/1419996/" rel="v:starring">吕艳婷</a>

	regstr := `<a.*?rel="v:starring">(.*?)</a>`
	return getMovieInfoByRegex(MovieHtml, regstr)
}

func getMovieRate(MovieHtml string) string {
	//<strong class="ll rating_num" property="v:average">8.8</strong>

	regstr := `<strong.*?property="v:average">(.*)</strong>`
	return getMovieInfoByRegex(MovieHtml, regstr)
}

func GetMovieUrlOther(html string) interface{} {
	span := `<a/s*class="item".*?href="(.*?)">.*?<img/s*src="(.*?)".*?</a>`

	reg := regexp.MustCompile(span)
	res := reg.FindAllStringSubmatch(html, -1)
	return res
}
