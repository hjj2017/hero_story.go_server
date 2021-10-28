package my_log

import "log"

func Init() {
	var w dailyFileWriter
	w.fileName = "./log/test.log"
	w.lastYearDay = -1
	w.logItemQ = make(chan []byte, 2048)

	log.SetOutput(&w)

	go w.startWork()
}
