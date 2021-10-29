package my_log

import "log"

// 日志队列大小
const queueSize = 512

//
// Init 初始化
//
func Init(fileName string, prefix string) {
	w := &dailyFileWriter{
		fileName: fileName,
		lastYearDay: -1,
		logQ: make(chan []byte, queueSize),
	}

	log.SetOutput(w)
	log.SetPrefix(prefix)
	log.SetFlags(log.Ltime | log.Lmicroseconds | log.Lshortfile)

	// 写手开始工作
	go w.startWork()
}
