package main

import (
	"fmt"
	log "hero_story.go_server/comm/my_log"
	"os"
	"path"
	"time"
)

func main() {
	ex, err := os.Executable()

	if nil != err {
		panic(err)
	}

	log.Init(path.Dir(ex) + "/log/biz_server.log")

	fmt.Println("start bizServer")
	log.Error("这里有一个错误")
	log.Info("Hello")
	log.Warning("我警告你")

	time.Sleep(2 * time.Second)
}
