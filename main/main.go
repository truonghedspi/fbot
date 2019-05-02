package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/fbot/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type serverConf struct {
	port string
}

func main() {
	api.Start()

	cfgPath := flag.String("c", "../conf/config.yml", "location of config file")
	flag.Parse()
	cfg := loadConfig(*cfgPath)

	log.Infof("run service on port: %s", cfg.port)
	http.ListenAndServe(fmt.Sprintf(":%s", cfg.port), nil)
}

func loadConfig(confFilePath string) serverConf {
	viper.SetConfigFile(confFilePath)
	if err := viper.ReadInConfig(); err != nil {
		log.Errorf("fail when read config file: %s", confFilePath)
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	var servConf serverConf
	servConf.port = viper.GetString("server.port")
	return servConf
}
