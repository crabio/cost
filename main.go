package main

import (
	"encoding/json"

	"github.com/iakrevetkho/cost/domain"
	"github.com/iakrevetkho/cost/internal/helpers"
	"github.com/jessevdk/go-flags"
	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
)

var opts domain.Opts
var cfg domain.Config

func init() {
	if _, err := flags.Parse(&opts); err != nil {
		logrus.WithError(err).Fatal("Can't parse arguments")
	}

	if err := configor.Load(&cfg, "config.yaml"); err != nil {
		logrus.WithError(err).Fatal("Can't parse conf")
	}

	if err := helpers.InitLogger(&cfg); err != nil {
		logrus.WithError(err).Fatal("Couldn't init logger")
	}

	if cfgJson, err := json.Marshal(cfg); err != nil {
		logrus.WithError(err).Fatal("Couldn't serialize config to JSON")
	} else {
		// Use Infof to prevent \" symbols if using WithField
		logrus.Infof("Loaded config: %s", cfgJson)
	}
}

func main() {
	logrus.WithField("filePath", opts.FilePath).Debug("opts")

}
