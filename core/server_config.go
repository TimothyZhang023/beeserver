package server

import (
	"github.com/astaxie/beego/config"
	log "github.com/astaxie/beego/logs"
	"runtime"
)

type ServerConfig struct {
	Port int // listen port

	MaxProcs int

	FileName string
	Config   config.Configer
}

func NewServerConfig(filename string) *ServerConfig {
	c, err := config.NewConfig("ini", filename)
	if err != nil {
		if filename == "" {
			c, _ = config.NewConfigData("ini", []byte(""))
		} else {
			panic("read config file failed " + err.Error())
		}
	}

 	log.SetLevel(log.LevelInfo)



	sc := &ServerConfig{
		Port:     c.DefaultInt("server::port", 6379),
		MaxProcs: c.DefaultInt("server::cpus", runtime.NumCPU()),
		FileName: filename,
		Config:   c,
	}

	runtime.GOMAXPROCS(sc.MaxProcs)
	log.Warn("Set runtime GOMAXPROCS to %d, total cpu num %d", sc.MaxProcs, runtime.NumCPU())

	return sc
}
