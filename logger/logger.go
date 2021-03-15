package logger

import (
	"log"
	"os"
)

const (
	// Control the details of the output log information, and cannot control the order and format of the output.
	// The output log will be separated by a colon after each item：such as 2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	Ldate         = 1 << iota     // Date：2009/01/23
	Ltime                         // Time：01:23:23
	Lmicroseconds                 // Microsecond time：01:23:23.123123（Used to enhance the Ltime bit）
	Llongfile                     // File full path name + line number： /a/b/c/d.go:23
	Lshortfile                    // File full path name + line number：d.go:23（Will overwrite Llongfile）
	LUTC                          // Use UTC time
	LstdFlags     = Ldate | Ltime // The initial value of the standard logger
	LogFile = "./golang.log"
)

func Logs(LogInfo string)  {
	/*
		O_RDWR      Open file in read-write mode
		O_APPEND    Append data to the end of the file when writing
		O_CREATE    If it does not exist, a new file will be created
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
		//logger.Println([]string{"hello", "golang logs - logger"})
	}

	defer logFile.Close()
}
