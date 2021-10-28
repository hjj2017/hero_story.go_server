package my_log

import "log"

func Init() {
	var w DailyFileWriter
	w.fileName = "./log/test.log"
	w.lastYearDay = -1

	log.SetOutput(&w)
}
