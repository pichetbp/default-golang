package main

import (
	"default-repo/helpers"
	dbUsecases "default-repo/libs/db/usecases"
	logUsecases "default-repo/libs/log/usecases"
	"fmt"

	"github.com/sirupsen/logrus"
)

func init() {
	helpers.InitConfig()
	logUsecases.NewLogEntry(logrus.DebugLevel, logUsecases.GenDefaultModels())
	dbUsecases.NewGormCon()
}

func main() {
	fmt.Println("main")
}
