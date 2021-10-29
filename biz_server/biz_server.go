package main

import (
	"fmt"
	"hero_story.go_server/comm/my_log"
	"log"
	"os"
	"path"
	"time"
)

func main() {
	ex, err := os.Executable()

	if nil != err {
		panic(err)
	}

	my_log.Init(path.Dir(ex) + "/log/biz_server.log", "[ biz_server ] ")

	fmt.Println("start bizServer")
	log.Println("Hello")

	time.Sleep(2 * time.Second)
}
