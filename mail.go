package main

import (
	"film/config"
	"film/processor"
	"fmt"
	"os"
)


func main() {
	baseConfig := initConfig()
	processor.Run(*baseConfig)
}

func initConfig() *config.Config  {
	initConfig ,e := config.Read(".env")
	if e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Cannot read config file %v \n", e)
		os.Exit(1)
	}
	return initConfig
}