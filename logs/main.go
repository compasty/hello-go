package main

import (
	"log"
	"os"

	"github.com/compasty/hello-go/logs/logger"
)

func main() {
	log.Println("Hello, World!")
	logger1 := log.New(os.Stdout, "HAHA:", log.LstdFlags|log.Lshortfile)
	logger1.Println("Hello, Logger1!")
	// 设定颜色
	errorLog := log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog := log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	errorLog.Println("Wow, Error!")
	infoLog.Println("Great Info!")

	logger.Errorf("Hello, Logger!")
}
