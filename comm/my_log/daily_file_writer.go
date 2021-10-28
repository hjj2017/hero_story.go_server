package my_log

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"path"
	"sync"
	"time"
)

//
// dailyFileWriter 日志文件写手
//
type dailyFileWriter struct {
	// 文件名称
	fileName string
	// 上一次写入日期
	lastYearDay int
	// 读写锁
	mutex sync.RWMutex
	// 日志项队列
	logItemQ chan []byte
	// 输出文件
	outputFile *os.File
}

// @override
func (w *dailyFileWriter) Write(byteArray []byte) (n int, err error) {
	if nil == byteArray ||
		len(byteArray) <= 0 {
		return 0, nil
	}

	if nil == w.logItemQ {
		return 0, errors.Errorf("日志项队列为空")
	}

	// 令字节数组入队
	w.logItemQ <- byteArray

	return len(byteArray), nil
}

func (w *dailyFileWriter) startWork() {
	for {
		// 获取日志项
		logItem := <-w.logItemQ

		// 获取文件句柄
		fileHandler, err := w.getFileHandler()

		if err != nil {
			fmt.Printf("%v", err)
			continue
		}

		if fileHandler == nil {
			fmt.Printf("文件句柄为空")
			continue
		}

		_, _ = os.Stdout.Write(logItem)
		_, _ = fileHandler.Write(logItem)
	}
}

// 获取输出文件句柄
func (w *dailyFileWriter) getFileHandler() (io.Writer, error) {
	yearDay := time.Now().YearDay()

	if w.lastYearDay == yearDay &&
		nil != w.outputFile {
		// 如果当前日期和上一次日期一样,
		// 且输出文件也不为空,
		return w.outputFile, nil
	}

	w.lastYearDay = yearDay

	if err := os.MkdirAll(path.Dir(w.fileName), os.ModePerm); nil != err {
		return nil, errors.Wrap(err, "创建目录失败")
	}

	newDailyFile := w.fileName + "." + time.Now().Format("20060102")

	fileHandler, err := os.OpenFile(
		newDailyFile,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)

	if err != nil {
		return nil, errors.Errorf("打开文件失败! %s, %s", newDailyFile, err)
	}

	if nil != w.outputFile {
		_ = w.outputFile.Close()
	}

	w.outputFile = fileHandler

	return fileHandler, nil
}
