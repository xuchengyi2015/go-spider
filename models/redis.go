package models

import (
	"github.com/astaxie/goredis"
	"github.com/xuchengyi2015/go-spider/tools"
	"strings"
)

const (
	URLQUEUE = "urlqueue"
	URLVISITED = "urlvisited"
)

var client goredis.Client

func init() {
	client.Addr = "127.0.0.1:6379"
}

func PushQueue(url string) {
	if !IsVisitedUrl(url) && !strings.Contains(url,"celebrity"){
		client.Lpush(URLQUEUE, []byte(url))
	}
}

func PopQueue() string {
	res, err := client.Rpop(URLQUEUE)
	tools.CheckErr(err)

	return string(res)
}

func GetQueueLength() int {
	length, err := client.Llen(URLQUEUE)
	if err != nil {
		return 0
	}
	return length
}

// 记录已经访问过的url
func SetQueue(url string) {
	client.Sadd(URLVISITED, []byte(url))
}

// 此url是否已经访问过了
func IsVisitedUrl(url string) bool {
	isVisited, err := client.Sismember(URLVISITED, []byte(url))
	if err != nil {
		return false
	}
	return isVisited
}
