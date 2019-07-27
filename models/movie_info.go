package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
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

func AddMovie(movie *MovieInfo) error {
	o := orm.NewOrm()
	id, err := o.Insert(movie)
	//tools.CheckErr(err)
	if err != nil {
		return err
	}

	fmt.Printf("插入数据ID:%v\n", id)
	return nil
}

func GetMovieInfo(MovieHtml string) (MovieInfo, error) {
	movie := MovieInfo{
		MovieName:      GetMovieName(MovieHtml),
		MovieCharacter: getMovieCharacter(MovieHtml),
		MovieRate:      getMovieRate(MovieHtml),
		//MovieID:      getMovieID(MovieHtml),
		CreationTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	err := AddMovie(&movie)

	return movie, err
}

//func getMovieID(movieHtml string) string {
//	regstr := `<a.*?href="https://movie.douban.com/subject/(.*?)/photos?type=R" title="点击看更多海报">.*?</a>`
//	return getMovieInfoByRegex(movieHtml, regstr)
//}

func getMovieInfoByRegex(movieHtml string, regstr string) string {
	reg := regexp.MustCompile(regstr)
	res := reg.FindAllStringSubmatch(movieHtml, -1)

	character := ""
	for _, v := range res {
		character += v[1] + " | "
	}
	return strings.Trim(character, " | ")
}

func GetMovieName(MovieHtml string) string {
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

func GetMoviesUrls(html string) []string {
	reg := regexp.MustCompile(`<a.*?href="(https://movie.douban.com/.*?)"`)
	result := reg.FindAllStringSubmatch(html, -1)

	var movieSets []string
	for _, v := range result {
		movieSets = append(movieSets, v[1])
	}

	return movieSets
}
