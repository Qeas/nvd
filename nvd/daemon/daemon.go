package daemon

import (
	log "github.com/Sirupsen/logrus"
	"github.com/docker/go-plugins-helpers/volume"
	"path/filepath"
)

const socketAddress = "/run/docker/plugins/nvd.sock"

var (
	defaultDir = filepath.Join(volume.DefaultDockerRootDirectory, "nvd")
)

func Start(cfgFile string, debug bool) {
	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	log.Info("Default docker root nvd: ", defaultDir)
	d := DriverAlloc(cfgFile)
	h := volume.NewHandler(d)
	log.Info("Driver Created, Handler Initialized")
	log.Info(h.ServeUnix(socketAddress, 0))
}
