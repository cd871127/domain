package common

import (
	"log"
	"os"
)

func InitLogger(file string) {
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	// 将文件设置为log输出的文件
	log.SetOutput(logFile)
	//log.SetPrefix("[qSkipTool]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}
