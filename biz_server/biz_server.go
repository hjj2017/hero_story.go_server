package main

import (
	"fmt"
	"hero_story.go_server/comm/my_log"
	"log"
)

func main() {
	my_log.Init()

	fmt.Println("start bizServer")
	log.Println("Hello")
}
