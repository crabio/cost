package main

import (
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
		logrus.WithError(err).Fatal(domain.ErrParseArguments)
	}

	if err := configor.Load(&cfg, "config.yaml"); err != nil {
		logrus.WithError(err).Fatal(domain.ErrParseConfig)
	}

	if err := helpers.InitLogger(&cfg); err != nil {
		logrus.WithError(err).Fatal(domain.ErrInitLogger)
	}

	if cfgJson, err := cfg.ToJson(); err != nil {
		logrus.WithError(err).Fatal(domain.ErrConvertConfigToJson)
	} else {
		// Use Infof to prevent \" symbols if using WithField
		logrus.Infof("Loaded config: %s", cfgJson)
	}
}

func main() {
	logrus.WithField("filePath", opts.FilePath).Debug("opts")

	sc, err := domain.NewSchemeConfigFromYaml(opts.FilePath)
	if err != nil {
		logrus.WithError(err).Fatal(domain.ErrParseSchemeConfig)
	}

	logrus.WithField("cfg", sc).Debug("parsed scheme config")
}
