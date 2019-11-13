/*
@Time : 2019-11-13 上午 9:34
@Author : Gerald
@File : RedisQueue.go
@Software: GoLand
*/
package models

import (
	"github.com/astaxie/goredis"
)

var (
	client    goredis.Client
	URL_QUEUE = "url_queue"
	URL_SET   = "url_set"
)

func ConnectRedis(addr string) {
	client.Addr = addr
}

// 使用 Lpush 方法，实现队列的入队
func PushQueue(url string) {
	pushErr := client.Lpush(URL_QUEUE, []byte(url))
	if pushErr != nil {
		panic(pushErr)
	}
}

// 使用 Rpop 方法，实现队列的出队
func PopQueue() string {
	url, popErr := client.Rpop(URL_QUEUE)
	if popErr != nil {
		panic(popErr)
	}

	return string(url)
}

// 使用 Sadd 方法，实现将 url 插入到集合中
func AddToSet(url string) {
	_, addErr := client.Sadd(URL_SET, []byte(url))
	if addErr != nil {
		panic(addErr)
	}
}

// 使用 Sismenber 方法，实现判断当前 url 是否已经存在
func IsVisited(url string) bool {
	ok, err := client.Sismember(URL_SET, []byte(url))
	if err != nil {
		panic(err)
	}

	return ok
}
