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

	log.Init(path.Dir(ex) + "/log/gateway_server.log")

	fmt.Println("start gatewayServer")
	log.Info("Hello")

	time.Sleep(2 * time.Second)
}
