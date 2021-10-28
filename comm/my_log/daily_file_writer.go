package my_log

import (
	"github.com/pkg/errors"
	"io"
	"os"
	"path"
	"sync"
	"time"
)

//
// DailyFileWriter 日志文件写手
//
type DailyFileWriter struct {
	// 文件名称
	fileName string
	// 上一次写入日期
	lastYearDay int
	// 读写锁
	mutex sync.RWMutex
	// 输出文件
	outputFile *os.File
}

// @override
func (w *DailyFileWriter) Write(p []byte) (n int, err error) {
	if nil == p {
		return 0, nil
	}

	w.mutex.Lock()
	defer w.mutex.Unlock()

	fileHandler, err := w.getFileHandler()

	if err != nil {
		return 0, err
	}

	if fileHandler == nil {
		return 0, errors.Wrap(err, "文件句柄为空")
	}

	os.Stdout.Write(p)
	return fileHandler.Write(p)
}

// 获取输出文件句柄
func (w *DailyFileWriter) getFileHandler() (io.Writer, error) {
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
		w.outputFile.Close()
	}

	w.outputFile = fileHandler

	return fileHandler, nil
}

func (w *DailyFileWriter) Close() error {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	var err error

	if nil != w.outputFile {
		err = w.outputFile.Close()
		w.outputFile = nil
	}

	return err
}
