package main

import (
	lib "gin_scaffold/common/utils/utilLib"
	"gin_scaffold/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	_ = lib.InitModule("./conf/local/", []string{"base", "mysql", "redis"})
	defer lib.Destroy()

	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
