package main

import (
	"fmt"
	"github.com/astaxie/goredis"
)

func main() {
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"

	//字符串操作
	client.Set("a", []byte("hello"))
	val, _ := client.Get("a")
	fmt.Printf("get value : %s\n", val)
	client.Del("a")

	//list操作
	vals := []string{"a", "b", "c", "d"}
	for _, v := range vals {
		client.Rpush("l", []byte(v))
	}

	dbvals, _ := client.Lrange("l", 0, 4)
	for i, v := range dbvals {
		fmt.Printf("%v : %v\n", i, string(v))
	}
}
