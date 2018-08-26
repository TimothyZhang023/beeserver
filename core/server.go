package core

import (
	"github.com/TimothyZhang023/beeserver/core/mipush"
	"github.com/TimothyZhang023/beeserver/core/util"
	"github.com/astaxie/beego"
	log "github.com/astaxie/beego/logs"

	"time"
)

type Server struct {
	Quit    chan bool //quit
	Wg      util.WaitGroupWrapper
	Startup time.Time

	Conf *ServerConfig
	//config
}

func (server Server) Init() {
	//beego.SetLevel(beego.LevelInformational)
}
func (server Server) WebServer() {
	go func() {
		beego.Run()
		log.Info("WebServer Quit ....")
	}()
}

func (server Server) MiPushModule() {
	mipush.UpdateFeed()
	timer := time.NewTicker(1 * time.Minute)
	for {
		select {
		case <-timer.C:
			mipush.UpdateFeed()

		case <-server.Quit:
			goto quit
		}
	}

quit:
	log.Info("MiPushModule Timer Job Service Quit ....")
	server.Quit <- true
	timer.Stop()
}

func (server Server) WaitingForQuit() {
	for {
		select {
		case quit := <-server.Quit:
			log.Info("Receive quit signal", quit)
			if quit {
				goto quit
			}
		}
	}

quit:
	server.Quit <- true
}
func (server Server) Close() {
	beego.BeeApp.Server.Close()
	server.Wg.Wait()
}

func (server *Server) MonitorQuit() {

	util.RegisterSignalAndWait()
	server.Quit <- true
	log.Info("MonitorQuit Quit ....")

}

func NewServer(c *ServerConfig) *Server {

	s := &Server{
		Conf:    c,
		Quit:    make(chan bool, 2),
		Startup: time.Now(),
	}

	return s
}
