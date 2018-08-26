package util

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/astaxie/beego/logs"
)

func RegisterSignalAndWait() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	quit := <-sc
	log.Warning("Receive signal", quit.String())
}
