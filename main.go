package main

import (
	"flag"
	"fmt"
	"github.com/TimothyZhang023/beeserver/core"
	_ "github.com/TimothyZhang023/beeserver/routers"
	"os"
	"runtime"
)

var (
	Cfg     = flag.String("config", "setting.ini", "path to config file")
	Version = flag.Bool("version", false, "show current version")
)

const logo string = `
-------------------------------------------------
  Server Build Time: %s.
  Git Commit Hash: %s.
  Go Version: %s.
  Go OS/Arch: %s/%s.
-------------------------------------------------
`

var buildstamp = "No Version Provided"
var githash = "No GitStash Provided"

func main() {
	fmt.Printf(logo, buildstamp, githash, runtime.Version(), runtime.GOOS, runtime.GOARCH)
	flag.Parse()

	if *Version {
		os.Exit(0)
	}

	conf := core.NewServerConfig(*Cfg)

	s := core.NewServer(conf)
	s.Init()
	s.Wg.Wrap(s.MonitorQuit)

	s.Wg.Wrap(s.WebServer)
	s.Wg.Wrap(s.MiPushModule)

	s.WaitingForQuit()
	s.Close()
	os.Exit(0)
}
