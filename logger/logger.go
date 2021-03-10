package logger

import (
	"log"
	"os"
)

const (
	// 控制输出日志信息的细节，不能控制输出的顺序和格式。
	// 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	Ldate         = 1 << iota     // 日期：2009/01/23
	Ltime                         // 时间：01:23:23
	Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
	Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
	Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
	LUTC                          // 使用UTC时间
	LstdFlags     = Ldate | Ltime // 标准logger的初始值
	LogFile = "./golang.log"
)

func Logs(LogInfo string)  {
	/*
		O_RDWR      读写模式打开文件
		O_APPEND    写操作时将数据附加到文件尾部
		O_CREATE    如果不存在将创建一个新文件
	*/

	logFile, err := os.OpenFile(LogFile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0766)
	if err != nil {
		log.Panic(err.Error())
	} else {
		log.Println(LogInfo)

		log.SetOutput(logFile)

		log.SetPrefix("[log]")
		log.SetFlags(log.LstdFlags | log.Llongfile | log.LUTC)
		log.Println([]string{LogInfo})

		//logger := log.New(logFile, "[logger]", log.LstdFlags | log.Lshortfile | log.LUTC)
		//logger.Println([]string{"你好", "golang日志 - logger"})
	}

	defer logFile.Close()
}
